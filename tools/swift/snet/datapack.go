/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-03 16:16:07
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */

// TLV序列化  解决TCP粘包问题
/*
   包的定义    datalen     +     dataid    +     data
   长度      uint32 4字节     uint32 4字节
*/
// 封包  拆包

package snet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"shenzhou/tools/swift/globalobj"
	"shenzhou/tools/swift/siface"
)

type DataPack struct {
}

// 初始化实例

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包的长度
func (dp *DataPack) GetHeadLen() uint32 {
	// datalen (4字节) + dataid(4字节)
	// 4 +4 = 8
	return 8
}

// 封包
func (dp *DataPack) Pack(msg siface.IMessage) ([]byte, error) {
	// 创建缓冲区
	buff := bytes.NewBuffer([]byte{})

	// 1将datalen写入buff
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// 2将dataid写入buff
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 3将data写入buff
	if err := binary.Write(buff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// 拆包  [[[datalen],[dataid] head] data]
func (dp *DataPack) UnPack(binaryData []byte) (siface.IMessage, error) {
	// 创建ioReader, 读取binaryData
	buff := bytes.NewReader(binaryData)

	// 解析Head  得到datalen和dataid

	// 创建信息存储结构  Message
	msg := &Message{}

	// 读datalen
	if err := binary.Read(buff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 读dataid
	if err := binary.Read(buff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断datalen是否超出了 MaxPackageSize
	if globalobj.GlobalConfig.MaxPackageSize > 0 && msg.DataLen > globalobj.GlobalConfig.MaxPackageSize {
		return nil, errors.New("package too large, config data len is %d, in face data len is %d")
	}
	return msg, nil
}
