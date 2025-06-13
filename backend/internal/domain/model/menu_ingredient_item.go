package model

// MenuIngredientItem は、メニューと食材の関連（レシピ）を表すモデルです。
type MenuIngredientItem struct {
	BaseModel
	MenuID       string     `gorm:"type:char(36);not null;uniqueIndex:uq_menu_ingredient" json:"menu_id"`
	IngredientID string     `gorm:"type:char(36);not null;uniqueIndex:uq_menu_ingredient" json:"ingredient_id"`
	Amount       float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	Menu         Menu       `gorm:"foreignKey:MenuID" json:"-"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID" json:"-"`
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (MenuIngredientItem) TableName() string {
	return "menu_ingredient_items"
}