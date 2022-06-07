/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-01 20:14:33
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println("client dail err :", err)
		return
	}
	for {

		_, err := conn.Write([]byte("hello swift V0.0.2  !"))
		if err != nil {
			fmt.Println("client write err : ", err)
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client write err :", err)
			return
		}
		fmt.Printf("recv from server : %s, 字节数: %d \n", buf[:n], n)
		time.Sleep(time.Second)
	}
}
