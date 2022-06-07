/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-03 16:15:59
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */

package siface

type IDataPack interface {
	// 获取包的长度
	GetHeadLen() uint32
	// 封包
	Pack(msg IMessage) ([]byte, error)
	// 拆包
	UnPack([]byte) (IMessage, error)
}
