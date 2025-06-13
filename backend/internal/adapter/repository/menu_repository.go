package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"meal-compass/backend/internal/domain/model"
	"meal-compass/backend/internal/domain/repository"
)

type menuRepository struct {
	db *gorm.DB
}

// NewMenuRepository は新しい menuRepository のインスタンスを生成します。
func NewMenuRepository(db *gorm.DB) repository.MenuRepository {
	return &menuRepository{db: db}
}

// FindRandomMenus は、指定された件数分のメニューをランダムに取得します。
func (r *menuRepository) FindRandomMenus(ctx context.Context, count int) ([]*model.Menu, error) {
	var menus []*model.Menu

	// 注意: ORDER BY RAND() はテーブルサイズが大きくなるとパフォーマンスが低下する可能性があります。
	// MVPではシンプルさを優先しますが、将来的にデータ件数が増える場合は、
	// 全IDを取得してからランダムにIDを選び、IN句で取得するなどの代替案を検討する必要があります。
	err := r.db.WithContext(ctx).
		Order("RAND()").
		Limit(count).
		Preload("MenuIngredientItems.Ingredient"). // レシピと食材情報も合わせて取得
		Find(&menus).Error

	if err != nil {
		return nil, fmt.Errorf("ランダムなメニューの取得に失敗しました: %w", err)
	}
	return menus, nil
}