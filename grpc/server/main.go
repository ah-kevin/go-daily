package main

import (
	context "context"
	"fmt"
	person "go-daily/grpc/pb/person"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{
		Name: "我收到了" + name,
	}
	return res, nil
}
func (*personServe) SearchIn(server person.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{
				Name: "完成了",
				Age:  18,
			})
			break
		}
	}
	return nil
}
func (*personServe) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.Name
	i := 0
	for {
		if i > 10 {
			break
		}
		i++
		time.Sleep(1 * time.Second)
		server.Send(&person.PersonRes{Name: "我拿到了" + name})
	}
	return nil
}
func (*personServe) SearchIO(server person.SearchService_SearchIOServer) error {
	str := make(chan string)
	i := 0
	go func() {
		for {
			req, _ := server.Recv()
			if i > 10 {
				str <- "结束"
				break
			}
			i++
			fmt.Println(i)
			str <- req.Name
		}
	}()
	for {
		s := <-str
		if s == "结束" {
			break
		}
		server.Send(&person.PersonRes{Name: <-str})
	}
	return nil
}
func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServe{})
	s.Serve(listen)
}
