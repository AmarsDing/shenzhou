/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-02 22:13:45
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

type IRouter interface {
	// 处理业务之前HOOK
	BeforeHandle(request IRequest)
	// 处理业务
	Handle(request IRequest)
	// 处理业务之后HOOK
	AfterHandle(request IRequest)
}
