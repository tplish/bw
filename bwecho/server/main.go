package main

import (
	"context"
	"github.com/tplish/bw/bwecho/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 构建tcp
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("tcp listen failed:%v", err)
	}

	// 创建gRPC
	server := grpc.NewServer()
	log.Printf("grpc EchoService start success")

	// 注册rcp方法到grpc
	protos.RegisterBwEchoServer(server, &BwEcho{})

	// 监听tcp
	_ = server.Serve(listen)
}

type BwEcho struct {
	protos.UnimplementedBwEchoServer
}

func (bw *BwEcho) Echo(ctx context.Context, req *protos.EchoReq) (*protos.EchoResp, error) {
	return &protos.EchoResp{Msg: req.Msg}, nil
}
