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
}

func CallbackFunc(conn *net.TCPConn, data []byte, n int) error {
	fmt.Println("Conn handle CallbackFunc")
	if _, err := conn.Write(data[:n]); err != nil {
		return err
	}
	return nil
}

// 开启服务
func (s *Server) Start() {

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

			dealConn := NewConnection(conn, cid, CallbackFunc)
			cid++
			go dealConn.Start()
		}
	}()

}

// 停止服务
func (s *Server) Stop() {

}

// 运行服务
func (s *Server) Serve() {
	// 启动server
	s.Start()
	// 阻塞
	select {}
}

// 初始化server
func NewServer(name string) siface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9001,
	}
	return s
}
