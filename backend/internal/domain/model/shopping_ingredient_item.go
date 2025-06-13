package model

// ShoppingIngredientItem は、買い物リストの個々のアイテムを表すモデルです。
type ShoppingIngredientItem struct {
	BaseModel
	PlanID       string     `gorm:"type:char(36);not null;uniqueIndex:uq_plan_ingredient" json:"plan_id"`
	IngredientID string     `gorm:"type:char(36);not null;uniqueIndex:uq_plan_ingredient" json:"ingredient_id"`
	Amount       float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	Bought       bool       `gorm:"not null;default:false" json:"bought"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID" json:"-"`
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (ShoppingIngredientItem) TableName() string {
	return "shopping_ingredient_items"
}