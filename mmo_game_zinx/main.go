package main

import (
	"fmt"
	"project/mmo_game_zinx/core"
	"project/zinx/ziface"
	"project/zinx/znet"
)

//当前客户端建立连接之后的HOOK函数
func OnConnnectionAdd(conn ziface.IConnection) {
	//创建一个Player对象
	player := core.NewPlayer(conn)

	//给客户端发送MsgID:1的消息:同步当前Player的ID给客户端
	player.SyncPid()

	//给客户端发送MsgID:200的消息：同步当前Player的初始位置给客户端
	player.BroadCastStartPosition()

	//将当前新上线的玩家添加到WorldManager中
	core.WorldMgrObj.AddPlayer(player)

	fmt.Println("=========> Player pid= ", player.Pid, " is arrived <=======")
}

func main() {
	//创建zinx server句柄
	s := znet.NewServer("MMO Game Zinx")

	//连接创建和销毁的HOOK钩子函数
	s.SetOnConnStart(OnConnnectionAdd)

	//注册一些路由业务

	//启动服务
	s.Serve()
}
