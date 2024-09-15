package repository

import (
	"budget/internal/model"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BudgetRepository interface {
	CreateBudget(budget *model.Budget) error
	GetAll() ([]model.Budget, error)
	GetById(budgetId string) (*model.Budget, error)
	Update(budget *model.Budget) error
}

type budgetRepository struct {
	db    *mongo.Collection
	redis *redis.Client
}

func NewBudgetRepository(db *mongo.Collection, redis *redis.Client) BudgetRepository {
	return &budgetRepository{db: db, redis: redis}
}

func (r *budgetRepository) CreateBudget(budget *model.Budget) error {
	_, err := r.db.InsertOne(context.Background(), budget)
	if err != nil {
		return err
	}

	budgetJSON, _ := json.Marshal(budget)
	return r.redis.Set(context.Background(), "budget:"+budget.ID, budgetJSON, 24*time.Hour).Err()
}

func (r *budgetRepository) GetAll() ([]model.Budget, error) {
	var budgets []model.Budget
	cursor, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var budget model.Budget
		if err := cursor.Decode(&budget); err != nil {
			return nil, err
		}
		budgets = append(budgets, budget)
	}
	return budgets, nil
}

func (r *budgetRepository) GetById(budgetId string) (*model.Budget, error) {
	budgetJSON, err := r.redis.Get(context.Background(), "budget:"+budgetId).Result()
	if err == nil {
		var budget model.Budget
		if err := json.Unmarshal([]byte(budgetJSON), &budget); err == nil {
			return &budget, nil
		}
	}

	var budget model.Budget
	err = r.db.FindOne(context.Background(), bson.M{"_id": budgetId}).Decode(&budget)
	if err != nil {
		return nil, errors.New("budget not found")
	}

	budgetJSONbyte, _ := json.Marshal(budget)
	r.redis.Set(context.Background(), "budget:"+budgetId, string(budgetJSONbyte), 24*time.Hour)

	return &budget, nil
}

func (r *budgetRepository) Update(budget *model.Budget) error {
	_, err := r.db.UpdateOne(context.Background(), bson.M{"_id": budget.ID}, bson.M{"$set": budget})
	if err != nil {
		return err
	}

	budgetJSON, _ := json.Marshal(budget)
	return r.redis.Set(context.Background(), "budget:"+budget.ID, budgetJSON, 24*time.Hour).Err()
}
