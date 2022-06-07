/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-01 20:14:33
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package main

import (
	"fmt"
	"io"
	"net"
	"shenzhou/tools/swift/snet"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println("client dail err :", err)
		return
	}
	for {

		// 封包
		dp := snet.NewDataPack()
		binary, err := dp.Pack(snet.NewMessage(0, []byte("swift 0.0.9 client message ")))
		if err != nil {
			fmt.Println("datapack err : ", err)
			break
		}
		// 向服务器写入数据
		_, err = conn.Write(binary)
		if err != nil {
			fmt.Println("write err ", err)
		}
		// 从服务器读取数据
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("client readfull err ", err)
			break
		}
		msgHead, err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("client read data unpack err:", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*snet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err ", err)
				break
			}
			fmt.Println("----------> Recv Server Msg id=", msg.GetMsgId(), "Msg data = ", string(msg.GetData()))
		}
		time.Sleep(time.Second * 1)
	}
}
