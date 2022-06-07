/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-03 16:01:54
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

type Message struct {
	Id      uint32 // 消息ID
	DataLen uint32 // 数据长度
	Data    []byte // 数据
}

// New对象

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 获取消息的id
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

// 获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen

}

// 获取数据内容
func (m *Message) GetData() []byte {
	return m.Data
}

// 设置消息ID
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

// 设置消息长度
func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

// 设置消息内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}
