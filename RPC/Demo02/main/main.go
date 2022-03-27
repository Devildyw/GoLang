package main

import (
	pb "RPC/Demo02/proto/SearchRequest"
	"fmt"
)

func main(){
	add:=&pb.SearchRequest{
		Query: "1",
		PageNumber: 2,
		ResultPerPage: 2,
		Search: []*pb.SearchRequest_Search{
			{Id: "1"},
		},
	}
	
	fmt.Println(add)
}