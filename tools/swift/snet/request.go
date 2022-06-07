/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-02 22:11:00
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import "shenzhou/tools/swift/siface"

type Request struct {
	conn siface.IConnection
	msg  siface.IMessage
}

func (r *Request) GetConnection() siface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
