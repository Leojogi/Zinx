package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	//初始化AOIManager
	aoiMgr := NewAOIManager(100, 300, 5, 100, 300, 5)

	//打印AOIManager
	fmt.Println(aoiMgr)
}
