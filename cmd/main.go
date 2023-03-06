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
	"time"

	"github.com/gin-gonic/gin"
	service "github.com/sean0427/micro-service-pratice-product-domain"
	config "github.com/sean0427/micro-service-pratice-product-domain/config"
	repository "github.com/sean0427/micro-service-pratice-product-domain/mongodb"
	handler "github.com/sean0427/micro-service-pratice-product-domain/net"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	port = flag.Int("port", 8080, "port")
)

func newMongoOptions() (*options.ClientOptions, string, error) {
	addr, err := config.GetMongoDBAddress()
	if err != nil {
		panic(err)
	}
	username, err := config.GetMongoDBUser()
	if err != nil {
		panic(err)
	}
	password, err := config.GetMongoDBPassword()
	if err != nil {
		panic(err)
	}

	port := config.GetMongoDBPort()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	databaseName, err := config.GetMongoDBName()
	if err != nil {
		panic(err)
	}

	return options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?directConnection=true&serverSelectionTimeoutMS=2000&retryWrites=true&w=majority", username, password, addr, port, databaseName)).
		SetServerAPIOptions(serverAPIOptions), databaseName, nil
}

func startServer() {
	fmt.Println("Starting server...")

	options, databaseName, err := newMongoOptions()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(err)
	}

	r := repository.New(mongoClient.Database(databaseName))
	s := service.New(r)
	h := handler.New(s)

	g := gin.Default()

	h.Register(g)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: g,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("erver stopped: %v", err)
		}
	}()
	log.Printf("server started and listening on %d", *port)

	// gracefully shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Print("received kill signal")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	wg.Wait()
}

func main() {
	startServer()
}
