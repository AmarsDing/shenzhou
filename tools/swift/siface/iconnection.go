/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-01 20:34:44
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

import "net"

type IConnection interface {
	// 启动链接
	Start()
	// 停止链接
	Stop()
	// 获取当前socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前链接的ID
	GetConnID() uint32
	// 获取远程客户端的 TCP状态  IP PORT
	RemoteAddr() net.Addr
	// 发送数据, 将数据发送给客户端
	SendMsg(uint32, []byte) error
	// 设置链接属性
	SetProperty(key string, value interface{})
	// 获取链接属性
	GetProperty(key string) (interface{}, error)
	// 移除链接属性
	RemoveProperty(key string)
}

// 定义个处理链接的业务的方法

type HandleFunc func(*net.TCPConn, []byte, int) error
