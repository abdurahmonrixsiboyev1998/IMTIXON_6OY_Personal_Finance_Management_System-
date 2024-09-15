package handler

import (
	"budget/internal/model"
	"budget/internal/service"
	"budget/proto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BudgetHandler struct {
	Service service.BudgetService
}

func NewBudgetHandler(service service.BudgetService) *BudgetHandler {
	return &BudgetHandler{Service: service}
}

func (h *BudgetHandler) CreateBudget(c *gin.Context) {
	var request model.Budget
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.ID = uuid.NewString()
	request.CreatedAt = time.Now()

	if _, err := h.Service.CreateBudget(c, &proto.CreateBudgetRequest{
		Category: request.Category,
		Amount:   float32(request.Amount),
		Currency: request.Currency,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget created successfully", "budget": request.ID})
}

func (h *BudgetHandler) GetBudgets(c *gin.Context) {
	budgets, err := h.Service.GetBudgets(c, &proto.GetBudgetsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budgets)
}

func (h *BudgetHandler) UpdateBudget(c *gin.Context) {
	var request model.Budget
	budgetId := c.Param("budgetId")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.Service.UpdateBudget(c, &proto.UpdateBudgetRequest{
		BudgetId: budgetId,
		Amount:   float32(request.Amount),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget updated successfully"})
}
