package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"meal-compass/backend/internal/usecase"
)

// PlanHandler は、計画関連のHTTPリクエストを処理します。
type PlanHandler struct {
	planUsecase usecase.PlanUsecase // Usecaseへのインターフェースを保持
}

// NewPlanHandler は新しい PlanHandler のインスタンスを生成します。
func NewPlanHandler(planUsecase usecase.PlanUsecase) *PlanHandler {
	return &PlanHandler{planUsecase: planUsecase}
}

// CreateNewPlan は POST /api/create-new-plan のリクエストを処理します。
func (h *PlanHandler) CreateNewPlan(c *gin.Context) {
	// リクエストボディをバインドするための構造体
	var req struct {
		PlannedMeals []struct {
			DateOffset int    `json:"date_offset"`
			MealPeriod string `json:"meal_period"`
		} `json:"planned_meals" binding:"required"`
	}

	// JSONボディを構造体にバインド。形式が不正な場合は400エラー。
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Usecase層に渡すためのデータに変換 (DTO: Data Transfer Object)
	plannedMealsDTO := make([]usecase.PlannedMealInput, len(req.PlannedMeals))
	for i, meal := range req.PlannedMeals {
		// バリデーション: date_offsetが負数でないか、meal_periodが正しい値か
		if meal.DateOffset < 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "date_offset cannot be negative"})
			return
		}
		if !(meal.MealPeriod == "MORNING" || meal.MealPeriod == "LUNCH" || meal.MealPeriod == "DINNER") {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid meal_period"})
			return
		}
		plannedMealsDTO[i] = usecase.PlannedMealInput{
			DateOffset: meal.DateOffset,
			MealPeriod: meal.MealPeriod,
		}
	}

	// Usecaseを呼び出し
	output, err := h.planUsecase.CreatePlan(c.Request.Context(), plannedMealsDTO)
	if err != nil {
		// Usecaseから返されたエラーに応じてレスポンスを返す
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plan: " + err.Error()})
		return
	}

	// 成功レスポンスを返す
	c.JSON(http.StatusCreated, output)
}

// GetMenuList は GET /api/menu-list/:shopping_plan_id のリクエストを処理します。
func (h *PlanHandler) GetMenuList(c *gin.Context) {
	planID := c.Param("shopping_plan_id")

	output, err := h.planUsecase.GetMenuList(c.Request.Context(), planID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get menu list"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"meals": output})
}

// GetIngredientList は GET /api/ingredient-list/:shopping_plan_id のリクエストを処理します。
func (h *PlanHandler) GetIngredientList(c *gin.Context) {
	planID := c.Param("shopping_plan_id")

	output, err := h.planUsecase.GetIngredientList(c.Request.Context(), planID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ingredients": output})
}