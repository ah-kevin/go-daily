package main

import (
	"context"
	"fmt"
	pb "go-daily/grpc/pb/person"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewSearchServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)
	res, err := client.Search(context.Background(), &pb.PersonReq{
		Name: "我是kk",
		Age:  18,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(res)

}
