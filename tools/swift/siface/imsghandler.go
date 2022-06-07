/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-04 16:05:45
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package siface

type IMsgHandler interface {
	// 执行对应的Router消息方法
	DoMsgHandler(request IRequest)
	// 添加路由
	AddRouter(msgId uint32, router IRouter)
	// 启动工作池
	StartWorkPool()
	// 将消息发送到任务队列
	SendMsgToQueue(req IRequest)
}
