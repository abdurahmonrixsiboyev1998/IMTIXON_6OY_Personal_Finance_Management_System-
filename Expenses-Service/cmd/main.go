package main

import (
	"expenses/config"
	"expenses/db"
	"expenses/internal/handler"
	"expenses/internal/repository"
	"expenses/internal/service"
	"expenses/proto"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := db.Connect(cfg.DbURI)
	defer dbConn.Close()

	db.Migrate(dbConn)

	repo := repository.NewTransactionRepository(dbConn)
	svc := service.NewTransactionService(repo)

	restHandler := handler.NewTransactionHandler(svc)
	grpcHandler := handler.NewGRPCTransactionHandler(svc)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/transactions/income", restHandler.LogIncome).Methods("POST")
	r.HandleFunc("/api/v1/transactions/expense", restHandler.LogExpense).Methods("POST")
	r.HandleFunc("/api/v1/transactions", restHandler.GetTransactions).Methods("GET")

	go func() {
		log.Printf("Starting REST server on port %s", cfg.RESTPort)
		if err := http.ListenAndServe(":"+cfg.RESTPort, r); err != nil {
			log.Fatalf("Error starting REST server: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterTransactionServiceServer(s, grpcHandler)

	log.Printf("Starting gRPC server on port %s", cfg.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
