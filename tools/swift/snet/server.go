/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-05-31 21:21:32
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import (
	"fmt"
	"net"
	"shenzhou/tools/swift/globalobj"
	"shenzhou/tools/swift/siface"
)

type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	// Router
	MsgHandler siface.IMsgHandler
	// server 链接管理器
	ConnMgr siface.IConnManager
	// 启动时HOOK函数
	OnConnStart func(conn siface.IConnection)
	// 结束时HOOK函数
	OnConnStop func(conn siface.IConnection)
}

// 开启服务
func (s *Server) Start() {
	fmt.Printf("[SWIFT SERVER START]\n ServerName:%s\n,Listenr:%s\n,Port:%d\n,MaxConn:%d\n,MaxPackageSzie:%d\n",
		globalobj.GlobalConfig.Name,
		globalobj.GlobalConfig.Host,
		globalobj.GlobalConfig.TcpPort,
		globalobj.GlobalConfig.MaxConn,
		globalobj.GlobalConfig.MaxPackageSize)
	// 获取地址
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建tcp服务
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 开启工作线程池 workerPool
	s.MsgHandler.StartWorkPool()
	var cid uint32
	cid = 0
	// 启动tcp服务,阻塞等待client链接
	go func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}
			// 设置最大链接个数判断, 超出最大，关闭此链接
			if s.ConnMgr.Len() >= globalobj.GlobalConfig.MaxConn {
				// TODO: 给客户端发送一个  server承载达到最大值
				fmt.Println(" too many connection")
				conn.Close()
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()

}

// 停止服务
func (s *Server) Stop() {
	s.ConnMgr.ClearConn()
	fmt.Println("---server stop---")
}

// 运行服务
func (s *Server) Serve() {
	// 启动server
	s.Start()
	// 阻塞
	select {}
}

// Add Router

func (s *Server) AddRouter(msgId uint32, router siface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}

// 获取server 的conncetion  manager
func (s *Server) GetConnMgr() siface.IConnManager {
	return s.ConnMgr
}

// 初始化server
func NewServer(name string) siface.IServer {
	s := &Server{
		Name:       globalobj.GlobalConfig.Name,
		IPVersion:  "tcp4",
		IP:         globalobj.GlobalConfig.Host,
		Port:       globalobj.GlobalConfig.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

// 注册OnConnStart
func (s *Server) SetOnConnStart(hookfunc func(conn siface.IConnection)) {
	s.OnConnStart = hookfunc
}

// 注册OnConnStop
func (s *Server) SetOnConnStop(hookfunc func(conn siface.IConnection)) {
	s.OnConnStop = hookfunc
}

// 调用OnConnStart
func (s *Server) CallOnConnStart(conn siface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("Call OnConnStart")
		s.OnConnStart(conn)
	}
}

// 调用OnConnStop
func (s *Server) CallOnConnStop(conn siface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("Call OnConnStop")
		s.OnConnStop(conn)
	}
}
