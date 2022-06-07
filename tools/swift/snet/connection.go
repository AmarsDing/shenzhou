/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-01 20:34:57
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import (
	"fmt"
	"net"
	"shenzhou/tools/swift/siface"
)

// 当前链接的模块
type Connection struct {
	// 当前链接的conn
	Conn *net.TCPConn

	// 链接ID
	ConnID uint32

	// 链接状态
	isClosed bool

	// 当前链接所绑定的业务处理方法
	handleAPI siface.HandleFunc

	// 停止链接时，线程退出信号
	ExitChan chan bool
}

// 初始化链接模块的方法

func NewConnection(conn *net.TCPConn, connID uint32, callback_api siface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChan:  make(chan bool),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader groutine is running....")
	defer fmt.Println(c.ConnID, " Reader is exit, remote addr is", c.RemoteAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, 1024)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err :", err)
			break
		}

		// 调用绑定的业务处理函数
		err = c.handleAPI(c.Conn, buf[:n], n)
		if err != nil {
			fmt.Println("handleApi err :", err, c.ConnID, c.RemoteAddr())
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println(c.ConnID, " start...")

	// 启动read数据的业务
	go c.StartReader()
	// 启动write数据的业务
}

func (c *Connection) Stop() {
	fmt.Println(c.ConnID, " is closed")
	if c.isClosed {
		return
	}
	c.isClosed = false
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() {

}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send() {

}
