package main

import (
	"fmt"
	"net/http"

	service "github.com/sean0427/micro-service-pratice-product-domain"
	handler "github.com/sean0427/micro-service-pratice-product-domain/net"
	repository "github.com/sean0427/micro-service-pratice-product-domain/rdbrepositry"
)

func startServer() {
	fmt.Println("Starting server...")
	r := repository.New()
	s := service.New(r)
	h := handler.New(s)

	handler := h.InitHandler()
	http.ListenAndServe(":8080", handler)

	fmt.Println("Stoping server...")
}

func main() {
	startServer()
}
