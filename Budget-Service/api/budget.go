package api

import (
	"context"
	"encoding/json"
	"net/http"

	pb "budget/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type BudgetHandler struct {
	BudgetService pb.BudgetServiceClient
}

func NewBudgetHandler(router *mux.Router, grpcConn *grpc.ClientConn) {
	budgetService := pb.NewBudgetServiceClient(grpcConn)
	handler := &BudgetHandler{BudgetService: budgetService}

	router.HandleFunc("/budgets", handler.CreateBudget).Methods("POST")
	router.HandleFunc("/budgets", handler.GetBudgets).Methods("GET")
	router.HandleFunc("/budgets/{id}", handler.UpdateBudget).Methods("PUT")
}

func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *http.Request) {
	var request pb.CreateBudgetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.BudgetService.CreateBudget(context.Background(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *BudgetHandler) GetBudgets(w http.ResponseWriter, r *http.Request) {
	request := &pb.GetBudgetsRequest{}

	response, err := h.BudgetService.GetBudgets(context.Background(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *BudgetHandler) UpdateBudget(w http.ResponseWriter, r *http.Request) {
	budgetID := mux.Vars(r)["id"]

	var request pb.UpdateBudgetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.BudgetId = budgetID

	response, err := h.BudgetService.UpdateBudget(context.Background(), &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}