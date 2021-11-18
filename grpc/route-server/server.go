package main

import (
	"context"
	"go-daily/grpc/route"
)

type routerGuideServer struct {
	route.UnimplementedRouteGuideServer
}

func (r routerGuideServer) GetFeature(ctx context.Context, point *route.Point) (*route.Feature, error) {
	panic("implement me")
}

func (r routerGuideServer) ListFeatures(rectangle *route.Rectangle, server route.RouteGuide_ListFeaturesServer) error {
	panic("implement me")
}

func (r routerGuideServer) RecordRoute(server route.RouteGuide_RecordRouteServer) error {
	panic("implement me")
}

func (r routerGuideServer) Recommend(server route.RouteGuide_RecommendServer) error {
	panic("implement me")
}

//
//func newServer() *route.RouteGuideServer {
//
//}
func main() {
	//listen, err := net.Listen("tcp", "localhost:5000")
	//if err != nil {
	//	log.Fatalln("cannot create a listener at the address")
	//}
	//grpcServer := grpc.NewServer()
	//route.RegisterRouteGuideServer(grpcServer,)
}
