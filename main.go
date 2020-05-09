package main

import (
	pb "cloudrun/gen/go/proto/v1"
	"cloudrun/service"
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/golang/glog"

	_ "github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	getCatEndpoint = flag.String("getCatEndpoint", "localhost:8090", "endpoint")
)

func runGW() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterCatHandlerFromEndpoint(ctx, mux, *getCatEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8080", mux)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	go func() {
		listenPort, err := net.Listen("tcp", ":8090")
		if err != nil {
			log.Fatalln(err)
		}
		server := grpc.NewServer()
		pb.RegisterCatServer(server, &service.MyCatService{})
		if err := server.Serve(listenPort); err != nil {
			log.Fatal("serve end")
		}
	}()
	if err := runGW(); err != nil {
		glog.Fatal(err)
	}
	//listenPort, err := net.Listen("tcp", ":8090")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//server := grpc.NewServer()
	//pb.RegisterCatServer(server, &service.MyCatService{})
	//if err := server.Serve(listenPort); err != nil {
	//	log.Fatal("serve end")
	//}
}
