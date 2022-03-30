package main

import (
	"fmt"
	"project/zinx/ziface"
	"project/zinx/znet"
)

/*
基于Zinx框架来开发的服务器端应用程序
*/

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {

	fmt.Println("Call PingRouter Handle...")

	//先读取客户端的数据，再回写ping ping ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(), ", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("err")
	}

}

//hello zinx test 自定义路由
type HelloZinxRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle...")

	//先读取客户端的数据，再回写ping ping ping
	fmt.Println("recv from client: msgID = ", request.GetMsgID(), ", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome to Zinx\n"))
	if err != nil {
		fmt.Println("err")
	}

}

//创建连接之后执行钩子函数
func DoConnctionBegin(conn ziface.IConnection) {
	fmt.Println("====>DoConnctionBegin is called...")
	if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}

	//给当前的连接设置一些属性
	fmt.Println("Set conn property...")
	conn.SetProperty("Name", "www.baidu.com")

}

//连接断开之前需要执行的钩子函数
func DoConnctionLost(conn ziface.IConnection) {
	fmt.Println("====>DoConnctionLost is called...")
	fmt.Println("conn ID = ", conn.GetConnID(), "is Lost...")

	//获取连接属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}

}

func main() {
	//1 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.6]")

	//2 注册连接Hook钩子函数
	s.SetOnConnStart(DoConnctionBegin)
	s.SetOnConnStop(DoConnctionLost)

	//3 给当前zinx框架添加一个自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//4启动server
	s.Serve()
}
