package main

import (
	"context"
	"fmt"
	pb "go-daily/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *pb.Req) (res *pb.Res, err error) {
	fmt.Println(req.GetMessage())
	return &pb.Res{Message: "我是从服务端返回的grpc的内容"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloGRPCServer(s, &server{})
	s.Serve(listen)
}
