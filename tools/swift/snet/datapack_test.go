/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-03 16:49:02
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

// 测试 拆包 和 封包
func TestDataPack(t *testing.T) {
	listner, err := net.Listen("tcp", "127.0.0.1:9013")
	if err != nil {
		fmt.Println("====1", err)
		return
	}
	go func() {
		for {
			conn, err := listner.Accept()
			if err != nil {
				fmt.Println("====2", err)
				return
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("====3", err)
						break
					}
					headMsg, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("====4", err)
						return
					}
					if headMsg.GetMsgLen() > 0 {
						msg := headMsg.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("====5", err)
							return
						}
						// 读完一整帧数据
						fmt.Println("Recv msg MsgID:", msg.GetMsgId(), "msgLen:", msg.GetMsgLen(), "msgData:", string(msg.GetData()))
					}

				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:9013")
	if err != nil {
		fmt.Println("====6", err)
		return
	}
	dp := NewDataPack()

	// 模拟粘包  发送两帧
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'s', 'w', 'i', 'f', 't'},
	}
	senddata1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("====7", err)
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'h', 'e', 'l', 'l', 'o', '!', '!'},
	}
	senddata2, err := dp.Pack(msg2)

	if err != nil {
		fmt.Println("====8", err)
	}
	senddata1 = append(senddata1, senddata2...)
	_, err = conn.Write(senddata1)
	fmt.Println("====send")
	if err != nil {
		fmt.Println("====9", err)
	}

	time.Sleep(time.Second * 4)
}
