package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"meal-compass/backend/internal/usecase"
)

// IngredientHandler は、買い物リストのアイテム関連のHTTPリクエストを処理します。
type IngredientHandler struct {
	planUsecase usecase.PlanUsecase // 買い物リストも計画の一部なのでplanUsecaseを利用
}

// NewIngredientHandler は新しい IngredientHandler のインスタンスを生成します。
func NewIngredientHandler(planUsecase usecase.PlanUsecase) *IngredientHandler {
	return &IngredientHandler{planUsecase: planUsecase}
}

// UpdateShoppingIngredientItem は PATCH /api/shopping_ingredient_items/:item_id のリクエストを処理します。
func (h *IngredientHandler) UpdateShoppingIngredientItem(c *gin.Context) {
	itemID := c.Param("item_id")

	var req struct {
		Bought bool `json:"bought"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	input := usecase.UpdateShoppingIngredientItemInput{
		ItemID: itemID,
		Bought: req.Bought,
	}

	output, err := h.planUsecase.UpdateShoppingIngredientItem(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get ingredient list"})
		}
		return
	}

	c.JSON(http.StatusOK, output)
}