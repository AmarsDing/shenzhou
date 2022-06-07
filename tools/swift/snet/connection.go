/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-01 20:34:57
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"shenzhou/tools/swift/globalobj"
	"shenzhou/tools/swift/siface"
	"sync"
)

// 当前链接的模块
type Connection struct {
	// connection 属于的server
	TcpServer siface.IServer
	// 当前链接的conn
	Conn *net.TCPConn

	// 链接ID
	ConnID uint32

	// 链接状态
	isClosed bool

	// 停止链接时，线程退出信号
	ExitChan chan bool

	// 无缓冲chan 用于读写groutine之间的消息通信
	msgChan chan []byte

	// 当前server的消息管理模块, 用来绑定msgid和业务处理关系
	MsgHandler siface.IMsgHandler

	// 链接属性集合
	property map[string]interface{}
	// 连接属性修改锁
	propertyLock sync.RWMutex
}

// 初始化链接模块的方法

func NewConnection(server siface.IServer, conn *net.TCPConn, connID uint32, handler siface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		MsgHandler: handler,
		ExitChan:   make(chan bool),
		msgChan:    make(chan []byte),
		property:   make(map[string]interface{}),
	}
	// 将conn加入的manager
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader groutine is running....")
	defer fmt.Println(c.ConnID, " Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()
	for {
		// 创建一个拆包对象
		dp := NewDataPack()
		// 将收到的数据按规则拆包
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head err:", err)
			break
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("Unpack headmsg err :", err)
		}
		if msg.GetMsgLen() > 0 {
			data := make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("Unpack read data err :", err)
				break
			}
			msg.SetData(data)
		}

		// 读取数据的head

		// 读取数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		// 判断是否开启了工作池机制
		if globalobj.GlobalConfig.WorkPoolSize > 0 {
			c.MsgHandler.SendMsgToQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

// 写消息goroutine  发送给客户端
func (c *Connection) StartWriter() {
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("server writer is err :", err)
				return
			}
		case <-c.ExitChan:
			// reader 退出  writer也要退出
			fmt.Println("server writer exit")
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println(c.ConnID, " start...")

	// 启动read数据的业务
	go c.StartReader()
	// 启动write数据的业务
	go c.StartWriter()

	// 服务connection正常启动, 此处调用创建链接后需要处理的hook业务
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	fmt.Println(c.ConnID, " is closed")
	// conn  close
	if c.isClosed {
		return
	}
	// 改变connection 当前状态
	c.isClosed = false
	// connction 销毁之前需要执行hook业务
	c.TcpServer.CallOnConnStop(c)
	// 关闭套接字
	c.Conn.Close()
	// 通知writer退出
	c.ExitChan <- true
	// 将conn 从manager中剔除
	c.TcpServer.GetConnMgr().Remove(c)
	// 关闭管道
	close(c.ExitChan)
	close(c.msgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// send msg
// 将发送给client的数据进行封包
func (c *Connection) SendMsg(msgid uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}

	// 创建封包对象
	dp := NewDataPack()

	binary, err := dp.Pack(NewMessage(msgid, data))
	if err != nil {
		fmt.Println("msg pack err:", err)
	}
	c.msgChan <- binary
	return nil
}

// 设置链接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

// 获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property")
	}
}

// 移除链接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}
