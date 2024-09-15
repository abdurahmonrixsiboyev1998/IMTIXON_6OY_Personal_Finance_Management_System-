package handler

import (
	"context"
	"expenses/internal/service"
	"expenses/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCTransactionHandler struct {
	proto.UnimplementedTransactionServiceServer
	service *service.TransactionService
}

func NewGRPCTransactionHandler(service *service.TransactionService) *GRPCTransactionHandler {
	return &GRPCTransactionHandler{service: service}
}

func (h *GRPCTransactionHandler) LogIncome(ctx context.Context, req *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid date format")
	}

	tx, err := h.service.LogIncome(req.Amount, req.Currency, req.Category, date)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to log income")
	}

	return &proto.TransactionResponse{
		Message:       "Income logged successfully",
		TransactionId: tx.TransactionID,
	}, nil
}

func (h *GRPCTransactionHandler) LogExpense(ctx context.Context, req *proto.TransactionRequest) (*proto.TransactionResponse, error) {
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid date format")
	}

	tx, err := h.service.LogExpense(req.Amount, req.Currency, req.Category, date)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to log expense")
	}

	return &proto.TransactionResponse{
		Message:       "Expense logged successfully",
		TransactionId: tx.TransactionID,
	}, nil
}

func (h *GRPCTransactionHandler) GetTransactions(ctx context.Context, req *proto.GetTransactionsRequest) (*proto.GetTransactionsResponse, error) {
	transactions, err := h.service.GetTransactions()
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get transactions")
	}

	var protoTransactions []*proto.Transaction
	for _, tx := range transactions {
		protoTransactions = append(protoTransactions, &proto.Transaction{
			TransactionId: tx.TransactionID,
			Type:          tx.Type,
			Amount:        tx.Amount,
			Currency:      tx.Currency,
			Category:      tx.Category,
			Date:          tx.Date.Format("2006-01-02"),
		})
	}

	return &proto.GetTransactionsResponse{
		Transactions: protoTransactions,
	}, nil
}