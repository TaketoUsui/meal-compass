package model

// Menu は、料理のメニュー情報を表すモデルです。
type Menu struct {
	BaseModel
	Name                string               `gorm:"type:varchar(255);not null;unique" json:"name"`
	MenuIngredientItems []MenuIngredientItem `gorm:"foreignKey:MenuID" json:"-"` // Menu has many MenuIngredientItems
}

// TableName は、GORMにテーブル名を明示的に指定します。
func (Menu) TableName() string {
	return "menus"
}