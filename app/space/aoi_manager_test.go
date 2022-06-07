/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-06 21:53:55
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package space

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoiMgr := NewAOIManager(100, 300, 100, 300, 4, 4)
	fmt.Println(aoiMgr)
}

func TestAOIManageAroundGridsByGid(t *testing.T) {
	aoiMgr := NewAOIManager(0, 250, 0, 250, 5, 5)
	for gid, _ := range aoiMgr.grids {
		grids := aoiMgr.GetAroundGridsByGid(gid)
		fmt.Println(" GID = ", gid, "Len = ", len(grids))
		gids := make([]int, 0, len(grids))
		for _, grid := range grids {
			gids = append(gids, grid.GID)
		}
		fmt.Println(" ==== AROUND GIDS = ", gids)
	}
}
