package main

import (
	"log"
	"net/http"

	"github.com/emorydu/common"
	pb "github.com/emorydu/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr          = common.EnvString("HTTP_ADDR", ":3000")
	ordersServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(ordersServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at: ", ordersServiceAddr)

	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Println("Starting HTTP server at: ", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
