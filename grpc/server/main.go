package main

import (
	context "context"
	person "go-daily/grpc/pb/person"
	"google.golang.org/grpc"
	"log"
	"net"
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
func (*personServe) SearchIn(person.SearchService_SearchInServer) error {
	return nil
}
func (*personServe) SearchOut(*person.PersonReq, person.SearchService_SearchOutServer) error {
	return nil
}
func (*personServe) SearchIO(person.SearchService_SearchIOServer) error {
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
