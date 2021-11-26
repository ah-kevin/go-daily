package main

import (
	"context"
	"fmt"
	pb "go-daily/grpc/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewHelloGRPCClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)
	req, err := client.SayHi(context.Background(), &pb.Req{Message: "我从客户端来"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(req.GetMessage())

}
