package model

import "time"

// MealPeriod は、食事の時間帯を表す型です。
type MealPeriod string

const (
	Morning MealPeriod = "MORNING"
	Lunch   MealPeriod = "LUNCH"
	Dinner  MealPeriod = "DINNER"
)

// PlanningMealItem は、計画された個々の食事を表すモデルです。
type PlanningMealItem struct {
	BaseModel
	PlanID     string     `gorm:"type:char(36);not null" json:"plan_id"`
	MenuID     string     `gorm:"type:char(36);not null" json:"menu_id"`
	Date       time.Time  `gorm:"type:date;not null" json:"date"`
	MealPeriod MealPeriod `gorm:"type:enum('MORNING', 'LUNCH', 'DINNER');not null" json:"meal_period"`
	Menu       Menu       `gorm:"foreignKey:MenuID" json:"-"`
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (PlanningMealItem) TableName() string {
	return "planning_meal_items"
}