package main

import (
	"fmt"
	"io"
	"net"
	"project/zinx/znet"
	"time"
)

func main() {
	fmt.Println("client0 start ...")

	time.Sleep(1 * time.Second)

	//1 直接链接远程服务区，得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client0 start err,exit!")
		return
	}

	for {
		//发送封包的message消息 MsgID:0
		dp := znet.NewDataPack()
		binnayMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx client0 Test Message")))
		if err != nil {
			fmt.Println("Pack error:", err)
			return
		}

		if _, err := conn.Write(binnayMsg); err != nil {
			fmt.Println("Write error:", err)
			return
		}

		//服务器就应该给我们回复一个messages数据，MsgID:1 pingpingping

		//1 先读取流中的head部分 得到ID 和 dataLen

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}

		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error ", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//2 再根据dataLen进行二次读取，将data读出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := is.ReadFull(conn, msg.Data); err != nil {
				return
			}

			fmt.Println("---> Recv Server Msg : ID = ", msg.Id, ", Len = ", msg.DataLen, ", data = ", string(msg.Data))
		}

		//阻塞
		time.Sleep(1 * time.Second)

	}

}
