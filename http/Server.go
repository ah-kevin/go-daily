package main

import "net/http"

type Server interface {
	Route(pattern string, handleFun http.HandlerFunc)
	Start(address string) error
}
