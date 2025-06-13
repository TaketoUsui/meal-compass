package model

// Ingredient は、個別の食材情報を表すモデルです。
type Ingredient struct {
	BaseModel
	TypeID                 string          `gorm:"type:char(36);not null" json:"type_id"`
	Name                   string          `gorm:"type:varchar(255);not null;unique" json:"name"`
	BaseAmount             float64         `gorm:"type:decimal(10,2);not null" json:"base_amount"`
	Unit                   string          `gorm:"type:varchar(50);not null" json:"unit"`
	ShelfLifeDaysUnopened *int            `gorm:"default:null" json:"shelf_life_days_unopened"`
	ShelfLifeDaysOpened   *int            `gorm:"default:null" json:"shelf_life_days_opened"`
	IngredientType        IngredientType  `gorm:"foreignKey:TypeID" json:"type"` // Ingredient belongs to IngredientType
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (Ingredient) TableName() string {
	return "ingredients"
}