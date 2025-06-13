package repository

import (
	"context"

	"meal-compass/backend/internal/domain/model"
)

// MenuRepository は、メニューに関連する永続化を担当するリポジトリです。
type MenuRepository interface {
	// FindRandomMenus は、指定された件数分のメニューをランダムに取得します。
	// 各メニューに必要な食材情報も合わせてEager Loadingすることを想定します。
	FindRandomMenus(ctx context.Context, count int) ([]*model.Menu, error)
}