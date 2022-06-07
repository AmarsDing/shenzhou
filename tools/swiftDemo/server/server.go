/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-05-31 21:34:00
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package main

import (
	"fmt"
	"shenzhou/tools/swift/siface"
	"shenzhou/tools/swift/snet"
)

type PingRouter struct {
	snet.BaseRouter
}

type HelloRouter struct {
	snet.BaseRouter
}

// 处理业务之前HOOK
// func (br *PingRouter) BeforeHandle(request siface.IRequest) {
// 	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping"))
// 	if err != nil {
// 		fmt.Println("before Handle  err :", err)
// 	}
// }

// 处理业务
func (br *PingRouter) Handle(request siface.IRequest) {
	fmt.Println("msg id :", request.GetMsgID(), "msg data: ", string(request.GetData()))
	// 先读取client数据
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("PING...."))
	if err != nil {
		fmt.Println(err)
	}
	// 回写ping
}

// 处理业务
func (br *HelloRouter) Handle(request siface.IRequest) {
	fmt.Println("msg id :", request.GetMsgID(), "msg data: ", string(request.GetData()))
	// 先读取client数据
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("Hello swift...."))
	if err != nil {
		fmt.Println(err)
	}
	// 回写ping
}

func BeginConnHook(conn siface.IConnection) {
	fmt.Println("=====> OnConnStart---")
	if err := conn.SendMsg(202, []byte("====OnConnStart====")); err != nil {
		fmt.Println(err)
	}
	// 测试链接属性
	// 给链接设置属性
	fmt.Println("set conn name ")
	conn.SetProperty("name", "amarsding")
	conn.SetProperty("home", "www.baidu.com")
	conn.SetProperty("github", "github.com/AmarsDing")
}

func BehindConnHook(conn siface.IConnection) {
	fmt.Println("======> OnConnStop ---", "Connid:", conn.GetConnID())

	// 获取链接属性
	if value, err := conn.GetProperty("name"); err == nil {
		fmt.Println("name:", value)
	}
	if value, err := conn.GetProperty("home"); err == nil {
		fmt.Println("home:", value)
	}
	if value, err := conn.GetProperty("github"); err == nil {
		fmt.Println("github:", value)
	}

}

// 处理业务之后HOOK
// func (br *PingRouter) AfterHandle(request siface.IRequest) {
// 	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping"))
// 	if err != nil {
// 		fmt.Println("after Handle  err :", err)
// 	}
// }

func main() {
	// 创建server
	s := snet.NewServer("snet v0.0.2")
	s.SetOnConnStart(BeginConnHook)
	s.SetOnConnStop(BehindConnHook)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	// 启动server
	s.Serve()
}
