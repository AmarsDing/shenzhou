/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-06 21:27:54
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package space

import (
	"fmt"
)

type AOIManager struct {
	// 区域的左坐标边界
	MinX int
	// 区域的右坐标边界
	MaxX int
	// X方向格子数量
	CountX int
	//区域的上边界坐标
	MinY int
	//区域的下边界坐标
	MaxY int
	// Y方向格子数量
	CountY int
	// 当前区域有哪些格子  map  key = gridid  value = gird{}
	grids map[int]*Grid
}

// 创建

func NewAOIManager(minX, maxX, minY, maxY, countX, countY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:   minX,
		MaxX:   maxX,
		MinY:   minY,
		MaxY:   maxY,
		CountX: countX,
		CountY: countY,
		grids:  map[int]*Grid{},
	}
	for y := 0; y < countY; y++ {
		for x := 0; x < countX; x++ {
			// 格子编号算法  id= idy*n + idx
			gid := y*countX + x
			aoiMgr.grids[gid] = NewGrid(
				gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridHight(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridHight(),
			)
		}
	}

	return aoiMgr
}

// 得到每个格子在X轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CountX
}

// 得到每个格子在Y轴方向的高度
func (m *AOIManager) gridHight() int {
	return (m.MaxY - m.MinY) / m.CountY
}

// 打印格子信息

func (m *AOIManager) String() string {

	// manager 信息
	s := fmt.Sprintf("AOIManager MinX=%d MaxX=%d \n MinY=%d MaxY = %d \n CountX = %d CountY = %d \n",
		m.MinX, m.MaxX, m.MinY, m.MaxY, m.CountX, m.CountY)
	// 各个grid信息
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid.String())
	}
	return s
}

// 根据格子的id得到周边九宫格格子的id
func (m *AOIManager) GetAroundGridsByGid(gid int) (grids []*Grid) {
	// 判断当前gid是否在aoimanager中
	if _, ok := m.grids[gid]; !ok {
		return
	}
	// 将当前格子放入九宫格
	grids = append(grids, m.grids[gid])
	// 判断gid左边是否有格子    判断右边是否有格子
	idx := gid % m.CountX
	// 根据gid得到x坐标
	if idx > 0 {
		grids = append(grids, m.grids[gid-1])
	}
	if idx < m.CountX-1 {
		grids = append(grids, m.grids[gid+1])
	}
	// 合并 x 轴 grids
	gridX := make([]int, 0, len(grids))
	for _, v := range grids {
		gridX = append(gridX, v.GID)
	}

	// 遍历 X轴 得到Y轴grid
	for _, v := range gridX {
		idy := v / m.CountY
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CountX])
		}
		if idy < m.CountY-1 {
			grids = append(grids, m.grids[v+m.CountX])
		}
	}
	return
}

// 通过坐标得到九宫格内的全部playerid
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	// 得到格子编号
	gid := m.GetGidByPos(x, y)
	// 根据格子的id获取周围格子id
	grids := m.GetAroundGridsByGid(gid)
	// 获取周围格子里的playerid
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayers()...)
	}
	return
}

// 通过x y 得到当前格子的gridid
func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridHight()
	return idy*m.CountX + idx
}

// 添加一个playerid 到一个格子中
func (m *AOIManager) AddPidToGrid(pid, gid int) {
	m.grids[gid].Add(pid)
}

// 移除一个格子中的playerid
func (m *AOIManager) RemovePidFromGrid(pid, gid int) {
	m.grids[gid].Del(pid)
}

// 通过gid 获取全部的playerid
func (m *AOIManager) GetPidsByGid(gid int) (playerids []int) {
	playerids = m.grids[gid].GetPlayers()
	return
}

// 通过坐标将Player添加到一个格子中
func (m *AOIManager) AddToGridByPos(pid int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	grid := m.grids[gid]
	grid.Add(pid)
}

// 通过坐标把一个player从一个格子中删除
func (m *AOIManager) RemoveFromGridByPos(pid int, x, y float32) {
	gid := m.GetGidByPos(x, y)
	grid := m.grids[gid]
	grid.Del(pid)
}
