/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-03 16:01:45
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

type IMessage interface {
	// 获取消息的id
	GetMsgId() uint32
	// 获取消息长度
	GetMsgLen() uint32
	// 获取数据内容
	GetData() []byte

	// 设置消息ID
	SetMsgId(uint32)
	// 设置消息长度
	SetMsgLen(uint32)
	// 设置消息内容
	SetData([]byte)
}
