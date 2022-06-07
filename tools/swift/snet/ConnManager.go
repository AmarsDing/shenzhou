/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-05 12:51:42
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import (
	"errors"
	"fmt"
	"shenzhou/tools/swift/siface"
	"sync"
)

type ConnManager struct {
	// 链接集合
	connections map[uint32]siface.IConnection
	// 读写锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]siface.IConnection),
	}
}

// 添加链接
func (cm *ConnManager) Add(conn siface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 将conn 加入集合
	cm.connections[conn.GetConnID()] = conn

	fmt.Println("Connection map add success, totals:", cm.Len(), "connction id = ", conn.GetConnID())
}

// 删除链接
func (cm *ConnManager) Remove(conn siface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 将conn 删除
	delete(cm.connections, conn.GetConnID())
	fmt.Println("Connection map del success, totals:", cm.Len(), "connction id = ", conn.GetConnID())
}

// 根据ID获取链接
func (cm *ConnManager) Get(connID uint32) (siface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("conncetion not found")
}

// 得到当前链接总数
func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

// 清除所有链接
func (cm *ConnManager) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	// 删除connect  并停止connection的工作
	for connid, conn := range cm.connections {
		// conn 停止工作
		conn.Stop()
		delete(cm.connections, connid)
	}
	fmt.Println("Clear All connection! conn num = ", cm.Len())
}
