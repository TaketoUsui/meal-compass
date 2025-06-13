package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel は、多くのモデルで共通するフィールドを定義します。
type BaseModel struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);onUpdate:CURRENT_TIMESTAMP(6)" json:"updated_at"`
}

// BeforeCreate は、GORMのフックで、レコード作成前に呼び出されます。
// IDが空の場合に新しいUUIDをセットします。
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}