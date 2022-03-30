package main

import (
	"fmt"
	"project/myDemo/protobufDemo/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	//定义一个Person结构对象
	person := &pb.Person{
		Name:   "Leo",
		Age:    18,
		Emails: []string{"1112@qq.com", "2222@gmail.com", "33333@gmail.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "122222222222",
				Type:   pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number: "144444444444",
				Type:   pb.PhoneType_WORK,
			},
			&pb.PhoneNumber{
				Number: "1555555555555",
				Type:   pb.PhoneType_HOME,
			},
		},
	}

	//编码
	//将person对象 就是将protobuf的message进行序列化，得到一个二进制文件
	data, err := proto.Marshal(person)
	//data就是我们要进行网络传输的数据，对端需要按照Message Person格式进行解析
	if err != nil {
		fmt.Println("marshal err:", err)
	}

	//解码
	newPerson := &pb.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		fmt.Println("unmaishal err:", err)
	}
	fmt.Println("源数据: ", person)
	fmt.Println("解码之后的数据: ", newPerson)

}
