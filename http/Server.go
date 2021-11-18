package main

import "net/http"

type Server interface {
	Route(pattern string, handleFun http.HandlerFunc)
	Start(address string) error
}

// sdkServer httpserver 抽象
type sdkServer struct {
	Name string
}

// Route 做注册路由
func (s *sdkServer) Route(pattern string, handleFun http.HandlerFunc) {
	panic("implement me")
}

func (s *sdkServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkServer{
		Name: name,
	}
}
