/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-05 12:51:29
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

type IConnManager interface {
	// 添加链接
	Add(conn IConnection)
	// 删除链接
	Remove(conn IConnection)
	// 根据ID获取链接
	Get(connID uint32) (IConnection, error)
	// 得到当前链接总数
	Len() int
	// 清除所有链接
	ClearConn()
}
