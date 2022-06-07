/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-02 22:08:30
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte

	GetMsgID() uint32
}
