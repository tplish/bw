package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/tplish/bw/bwapi/protos"
	"google.golang.org/grpc"
	"net"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	// 构建tcp
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		glog.Error(err)
	}

	// 创建gRPC
	server := grpc.NewServer()
	glog.Info("ok")
	glog.Info("ok")

	// 注册rcp方法到grpc
	protos.RegisterBwApiServer(server, &BwApi{})

	// 监听tcp
	_ = server.Serve(listen)
}

type BwApi struct {
	protos.UnimplementedBwApiServer
}

func (bw *BwApi) ListItems(ctx context.Context, req *protos.ListItemsReq) (*protos.ListItemsResp, error) {
	var items []*protos.Item
	for i := 0; i < 10; i++ {
		items = append(items, &protos.Item{ItemId: uint32(i)})
	}
	return &protos.ListItemsResp{Items: items}, nil
}
