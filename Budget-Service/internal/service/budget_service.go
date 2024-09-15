package service

import (
	"budget/internal/model"
	"budget/internal/repository"
	"budget/proto"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BudgetService interface {
	CreateBudget(ctx context.Context, req *proto.CreateBudgetRequest) (*proto.CreateBudgetResponse, error)
	GetBudgets(ctx context.Context, req *proto.GetBudgetsRequest) (*proto.GetBudgetsResponse, error)
	UpdateBudget(ctx context.Context, req *proto.UpdateBudgetRequest) (*proto.UpdateBudgetResponse, error)
}

type budgetService struct {
	repo repository.BudgetRepository
	proto.UnimplementedBudgetServiceServer
}

func NewBudgetService(repo repository.BudgetRepository) proto.BudgetServiceServer {
	return &budgetService{repo: repo}
}

func (s *budgetService) CreateBudget(ctx context.Context, req *proto.CreateBudgetRequest) (*proto.CreateBudgetResponse, error) {
	createdTime := time.Now().Add(time.Hour * 5)
	budget := &model.Budget{
		ID:        primitive.NewObjectID().Hex(),
		Category:  req.Category,
		Amount:    float64(req.Amount),
		Currency:  req.Currency,
		CreatedAt: createdTime,
	}

	err := s.repo.CreateBudget(budget)
	if err != nil {
		return nil, err
	}

	return &proto.CreateBudgetResponse{
		Message:  "Budget created successfully",
		BudgetId: budget.ID,
	}, nil
}

func (s *budgetService) GetBudgets(ctx context.Context, req *proto.GetBudgetsRequest) (*proto.GetBudgetsResponse, error) {
	budgets, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var pbBudgets []*proto.Budget
	for _, b := range budgets {
		pbBudgets = append(pbBudgets, &proto.Budget{
			Id:       b.ID,
			Category: b.Category,
			Amount:   float32(b.Amount),
			Spent:    float32(b.Spent),
			Currency: b.Currency,
		})
	}

	return &proto.GetBudgetsResponse{
		Budgets: pbBudgets,
	}, nil
}

func (s *budgetService) UpdateBudget(ctx context.Context, req *proto.UpdateBudgetRequest) (*proto.UpdateBudgetResponse, error) {
	budget, err := s.repo.GetById(req.BudgetId)
	if err != nil {
		return nil, err
	}

	budget.Amount = float64(req.Amount)
	err = s.repo.Update(budget)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateBudgetResponse{
		Message: "Budget updated successfully",
	}, nil
}
