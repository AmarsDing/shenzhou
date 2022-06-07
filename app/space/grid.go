/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-06 20:47:22
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */

package space

import (
	"fmt"
	"sync"
)

type Grid struct {
	// 格子的ID
	GID int
	// 格子的左边 坐标边界
	MinX int
	// 格子的右边 坐标边界
	MaxX int
	// 格子的上边界 坐标
	MinY int
	// 格子的下边界 坐标
	MaxY int
	// 当前格子内玩家或者物体成员的ID集合
	playerIDs map[int]bool
	// playerIDS  集合锁
	pIDLock sync.RWMutex
}

// 创建空间
func NewGrid(gid, minx, maxx, miny, maxy int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minx,
		MaxX:      maxx,
		MinY:      miny,
		MaxY:      maxy,
		playerIDs: make(map[int]bool),
	}
}

// 给格子  添加  一个玩家
func (g *Grid) Add(playerid int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerid] = true
}

// 给格子  删除 一个玩家
func (g *Grid) Del(playerid int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerid)
}

// 给格子  删除 一个玩家
func (g *Grid) GetPlayers() (playerids []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for k, _ := range g.playerIDs {
		playerids = append(playerids, k)
	}
	return
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id = %d \n minX = %d maxX = %d \n minY = %d maxY = %d\n playerids: %v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
