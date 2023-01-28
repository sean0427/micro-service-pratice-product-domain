package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	service "github.com/sean0427/micro-service-pratice-product-domain"
	config "github.com/sean0427/micro-service-pratice-product-domain/config"
	repository "github.com/sean0427/micro-service-pratice-product-domain/mongodb"
	handler "github.com/sean0427/micro-service-pratice-product-domain/net"
)

var (
	port = flag.Int("port", 8080, "port")
)

func newMongoDB() (*mongo.Client, error) {
	addr, err := config.GetMongoDBAddress()
	if err != nil {
		panic(err)
	}
	username, err := config.GetMongoDBName()
	if err != nil {
		panic(err)
	}
	password, err := config.GetMongoDBPassword()
	if err != nil {
		panic(err)
	}
	port := config.GetMongoDBPort()
	return mongo.Connect(context.Background(), options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s", addr, port)).
		SetAuth(options.Credential{
			Username: username,
			Password: password,
		}))
}

func startServer() {
	fmt.Println("Starting server...")

	mongoClient, err := newMongoDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	databaseName, err := config.GetMongoDBName()
	if err != nil {
		panic(err)
	}

	r := repository.New(mongoClient.Database(databaseName))
	s := service.New(r)
	h := handler.New(s)

	handler := h.Handler()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: handler,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("admission-webhook-server stopped: %v", err)
		}
	}()
	log.Printf("admission webhook server started and listening on %d", *port)

	// gracefully shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Print("admission webhook received kill signal")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	wg.Wait()
}

func main() {
	startServer()
}
