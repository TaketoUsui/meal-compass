package model

import "time"

// ShoppingPlan は、買い物計画全体を表すモデルです。
type ShoppingPlan struct {
	BaseModel
	PeriodStartAt           time.Time                `gorm:"type:date;not null" json:"period_start_at"`
	PlanningMealItems       []PlanningMealItem       `gorm:"foreignKey:PlanID" json:"-"`
	ShoppingIngredientItems []ShoppingIngredientItem `gorm:"foreignKey:PlanID" json:"-"`
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (ShoppingPlan) TableName() string {
	return "shopping_plans"
}