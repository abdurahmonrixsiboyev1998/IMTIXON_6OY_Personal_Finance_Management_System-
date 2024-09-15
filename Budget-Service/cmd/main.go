package main

import (
	"budget/api"
	"budget/config"
	"budget/internal/repository"
	"budget/internal/service"
	"budget/proto"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.LoadConfig()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	db := client.Database("personal_finance")
	budgetRepo := repository.NewBudgetRepository(db.Collection("budgets"), redisClient)
	budgetService := service.NewBudgetService(budgetRepo)

	grpcServer := grpc.NewServer()
	proto.RegisterBudgetServiceServer(grpcServer, budgetService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	router := mux.NewRouter()
	grpcConn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC: %v", err)
	}
	api.NewBudgetHandler(router, grpcConn)
	fmt.Println("Start...")

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
