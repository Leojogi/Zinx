package core

import "fmt"

//定义一些AOI的边界值
const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

/*
	AOI区域管理模块
*/

type AOIManager struct {
	//区域的左边界坐标
	MinX int

	//区域的右边界坐标
	MaxX int

	//X方向格子的数量
	CntsX int

	//区域的上边界坐标
	MinY int

	//区域的下边界坐标
	MaxY int

	//Y方向格子的数量
	CntsY int

	//当前区域中有哪些格子map-key=格子的ID，value=格子对象
	grids map[int]*Grid
}

/*
	初始化一个AOI区域管理模块
*/
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}

	//给AOI初始化区域的所有格子进行编号和初始化
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//计算格子ID 根据x,y编号
			//格子编号： id = idy *cntsX + idx
			gid := y*cntsX + x

			//初始化gid格子
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinX+y*aoiMgr.gridLength(),
				aoiMgr.MinX+(y+1)*aoiMgr.gridLength(),
			)

		}
	}
	return aoiMgr
}

//得到每个格子在X轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

//得到每个格子在Y轴方向的长度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

//打印格子信息
func (m *AOIManager) String() string {
	//打印AOIManager信息
	s := fmt.Sprintf("AOIManager:\n MinX:%d, MaxX:%d, cntsX;%d, MinY;%d, MaxY;%d, cntsY;%d\n Grids in AOIManagher:\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)

	//打印全部格子信息
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}

//根据格子GID得到周边九宫格格子集合
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否在AOIManager中
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//将当前gid本身加入到九宫格切片中
	grids = append(grids, m.grids[gID])

	//需要判断gID左边是否有格子？右边是否有格子
	//需要通过gID得到当前格子x轴的编号    idx = id %nx
	idx := gID % m.CntsX

	//判断idx编号左边是否有格子？
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}

	//判断idx编号右边是否有格子？
	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	//将x轴当前的格子去取出，进行遍历，再分别得到每个格子上下是否还有格子
	//得到当前x轴格子的ID集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	//遍历gidsX编号中每个格子的gid
	for _, v := range gidsX {
		//得到当前格子id的y轴的编号   idy = id / ny
		idy := v / m.CntsY

		//gid上边是否还有格子
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntsX])
		}
		//gid下边是否还有格子
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[v+m.CntsX])
		}
	}

	return

}

//通过x,y横纵坐标得到GID格子标号
func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridLength()

	return idx + idy*m.CntsX
}

//通过横纵坐标得到周边九宫格内全部的PlayerIDs
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	//得到当前玩家的GID格子id
	gID := m.GetGidByPos(x, y)

	//通过GID得到周边九宫格信息
	grids := m.GetSurroundGridsByGid(gID)

	//将九宫格的信息里的全部Player的id 累加到playerIDs
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		fmt.Printf("====> grid ID : %d, pids: %v ====", grid.GID, grid.GetPlayerIDs())
	}

	return
}

//添加一个PlayerID一个格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

//移除一个格子中的PlayerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

//通过GID获取全部的PlayerID
func (m *AOIManager) GetPidsByGid(gID int) (playIDs []int) {
	playIDs = m.grids[gID].GetPlayerIDs()
	return
}

//通过坐标将Player添加到一个格子中
func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}

//通过坐标把一个Player从一个格子中删除
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(pID)
}
