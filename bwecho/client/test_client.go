package main

import (
	"context"
	"github.com/tplish/bw/bwecho/protos"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 链接tcp端口
	connect, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("tcp connect failed:%v", err)
	}

	// 新建client
	client := protos.NewBwEchoClient(connect)

	// 调用
	resp, err := client.Echo(context.Background(), &protos.EchoReq{Msg: "jack"})
	if err != nil {
		log.Fatalf("client.Echo err: %v", err)
	}
	log.Printf("%v", resp)
}
