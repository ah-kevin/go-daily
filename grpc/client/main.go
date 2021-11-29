package main

import (
	"context"
	"fmt"
	pb "go-daily/grpc/pb/person"
	"google.golang.org/grpc"
	"log"
	"sync"
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
	//res, err := client.Search(context.Background(), &pb.PersonReq{
	//	Name: "我是kk",
	//	Age:  18,
	//})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//c, _ := client.SearchIn(context.Background())
	//i := 0
	//for {
	//	if i > 10 {
	//		res, _ := c.CloseAndRecv()
	//		fmt.Println(res)
	//		break
	//	}
	//	time.Sleep(1*time.Second)
	//	c.Send(&pb.PersonReq{
	//		Name: "我是进来的信息",
	//	})
	//	i++
	//}
	//out, _ := client.SearchOut(context.Background(), &pb.PersonReq{Name: "kkkk"})
	//for  {
	//	req, err := out.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(req)
	//}
	c, _ := client.SearchIO(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			//time.Sleep(1 * time.Second)
			err := c.Send(&pb.PersonReq{Name: "bjke"})
			if err != nil {
				wg.Done()
				break
			}
		}
	}()
	go func() {
		for {
			res, err := c.Recv()
			if err != nil {
				wg.Done()
				fmt.Println(err)
				break
			}
			fmt.Println(res)
		}
	}()
	wg.Wait()
}
