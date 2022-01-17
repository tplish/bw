package main

import (
	"context"
	"github.com/tplish/bw/bwapi/protos"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 链接tcp端口
	connect, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("tcp connect failed:%v", err)
	}

	// 新建client
	client := protos.NewBwApiClient(connect)

	// 调用
	resp, err := client.ListItems(context.Background(), &protos.ListItemsReq{})
	if err != nil {
		log.Fatalf("client.Echo err: %v", err)
	}
	log.Printf("%v", resp)
}
