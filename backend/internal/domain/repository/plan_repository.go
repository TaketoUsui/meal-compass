package repository

import (
	"context"

	"meal-compass/backend/internal/domain/model"
)

// PlanRepository は、買い物計画に関連する永続化を担当するリポジトリです。
type PlanRepository interface {
	// Transaction は、引数で受け取った関数をトランザクション内で実行します。
	Transaction(ctx context.Context, fn func(txRepo PlanRepository) error) error

	// CreateShoppingPlan は、新しい買い物計画を保存します。
	CreateShoppingPlan(ctx context.Context, plan *model.ShoppingPlan) error
	// CreatePlanningMealItems は、複数の食事予定を保存します。
	CreatePlanningMealItems(ctx context.Context, meals []*model.PlanningMealItem) error
	// CreateShoppingIngredientItems は、複数の買い物リストアイテムを保存します。
	CreateShoppingIngredientItems(ctx context.Context, ingredients []*model.ShoppingIngredientItem) error

	// FindMealsByPlanID は、指定された計画IDに紐づく食事予定のリストを取得します。メニュー情報もEager Loadingします。
	FindMealsByPlanID(ctx context.Context, planID string) ([]*model.PlanningMealItem, error)
	// FindShoppingIngredientsByPlanID は、指定された計画IDに紐づく買い物リストを取得します。食材情報もEager Loadingします。
	FindShoppingIngredientsByPlanID(ctx context.Context, planID string) ([]*model.ShoppingIngredientItem, error)

	// FindShoppingIngredientItemByID は、指定された買い物アイテムIDでアイテムを1件取得します。
	FindShoppingIngredientItemByID(ctx context.Context, itemID string) (*model.ShoppingIngredientItem, error)
	// UpdateShoppingIngredientItem は、買い物リストのアイテム情報（主に'bought'フラグ）を更新します。
	UpdateShoppingIngredientItem(ctx context.Context, item *model.ShoppingIngredientItem) error
}