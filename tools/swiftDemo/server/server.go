/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-05-31 21:34:00
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package main

import "shenzhou/tools/swift/snet"

func main() {
	// 创建server
	s := snet.NewServer("snet v0.0.2")

	// 启动server
	s.Serve()
}
