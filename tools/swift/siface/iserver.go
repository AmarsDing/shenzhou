/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-05-31 21:20:59
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

type IServer interface {

	// 启动服务
	Start()

	// 停止服务
	Stop()

	// 运行服务
	Serve()
}
