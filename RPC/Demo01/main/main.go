package main

import (
	pb "RPC/Demo01/main/addressbook"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)
func main(){
	p:=pb.Person{
		Id: 1234,
		Name: "Isaac",
		Email: "1842501760@qq.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number:"555-4321",Type:pb.Person_HOME},
		},
	}
	addressbook:= &pb.AddressBook{
		People: []*pb.Person{&p},
	}

	//将addressBook对象序列化
	out,err:=proto.Marshal(addressbook)
	if err!=nil{
		panic(err)
	}
	//将序列化的数据写入文件
	if err = os.WriteFile("address",out,0644);err!=nil{
		panic(err)
	}
	// 从文件中读取数据
    in, err := os.ReadFile("address")
    if err != nil {
        panic(err)
    }
	//将数据反序列化为addressbook对象
	address:=&pb.AddressBook{}
	if err= proto.Unmarshal(in,address); err!=nil{
		panic(err)
	}
	fmt.Println(address)
}