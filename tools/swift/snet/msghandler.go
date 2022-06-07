/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-04 16:05:58
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import (
	"fmt"
	"shenzhou/tools/swift/globalobj"
	"shenzhou/tools/swift/siface"
)

type MsgHandle struct {
	// id 对应的处理方法
	Apis map[uint32]siface.IRouter
	// 消息队列
	TaskQueue []chan siface.IRequest
	// work工作池
	WorkPoolSize uint32
}

// 创建MsgHandle

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:         make(map[uint32]siface.IRouter),
		WorkPoolSize: globalobj.GlobalConfig.WorkPoolSize,
		TaskQueue:    make([]chan siface.IRequest, globalobj.GlobalConfig.WorkPoolSize),
	}
}

// 执行对应的Router消息方法
func (mh *MsgHandle) DoMsgHandler(request siface.IRequest) {
	// 从request 中找到 msgid
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("msgid dont exist, need register, msgid = ", request.GetMsgID())
		return
	}
	// 根据msgid 调度router
	handler.BeforeHandle(request)
	handler.Handle(request)
	handler.AfterHandle(request)
}

// 添加路由
func (mh *MsgHandle) AddRouter(msgId uint32, router siface.IRouter) {
	// 判断路由是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		fmt.Println("msgid already exist,msgid = ", msgId)
		return
	}
	mh.Apis[msgId] = router
	fmt.Println("add router success  routerid = ", msgId)
	// 添加msg与路由的绑定关系
}

// 启动一个worker工作池
func (mh *MsgHandle) StartWorkPool() {
	// 根据workpoolsize 创建worker goroutine
	for i := 0; i < int(mh.WorkPoolSize); i++ {
		// 创建消息队列 chan 空间
		mh.TaskQueue[i] = make(chan siface.IRequest, globalobj.GlobalConfig.MaxWorkerTaskLen)
		// 启动一个work， 阻塞等待消息
		go mh.startOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动worker工作流程
func (mh *MsgHandle) startOneWorker(workId int, taskQueue chan siface.IRequest) {
	for {
		select {
		// 收到消息，对消息进行处理
		case req := <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

// 将数据交给TaskQueue  由worker进行数据处理
func (mh *MsgHandle) SendMsgToQueue(req siface.IRequest) {
	// 将消息平均分配到WorkPool
	workid := req.GetConnection().GetConnID() % mh.WorkPoolSize
	fmt.Println("Conn ID :", req.GetConnection().GetConnID(), "Msg Id:", req.GetMsgID(), "--->Workid:", workid)
	// 将消息发送给work对应的TaskQueue
	mh.TaskQueue[workid] <- req
}
