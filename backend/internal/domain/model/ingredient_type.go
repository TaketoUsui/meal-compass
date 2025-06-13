package model

// IngredientType は、食材の分類を表すモデルです。
type IngredientType struct {
	BaseModel
	Name        string       `gorm:"type:varchar(255);not null;unique" json:"name"`
	Ingredients []Ingredient `gorm:"foreignKey:TypeID" json:"-"` // IngredientType has many Ingredients
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (IngredientType) TableName() string {
	return "ingredient_types"
}