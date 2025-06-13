package repository

import (
	"context"
	"gorm.io/gorm"

	"meal-compass/backend/internal/domain/model"
	"meal-compass/backend/internal/domain/repository"
)

// planRepository は repository.PlanRepository の実装です。
type planRepository struct {
	db *gorm.DB
}

// NewPlanRepository は新しい planRepository のインスタンスを生成します。
func NewPlanRepository(db *gorm.DB) repository.PlanRepository {
	return &planRepository{db: db}
}

// Transaction は、引数で受け取った関数をトランザクション内で実行します。
func (r *planRepository) Transaction(ctx context.Context, fn func(txRepo repository.PlanRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// トランザクション用の新しいリポジトリインスタンスを作成
		txRepo := NewPlanRepository(tx)
		return fn(txRepo)
	})
}

func (r *planRepository) CreateShoppingPlan(ctx context.Context, plan *model.ShoppingPlan) error {
	return r.db.WithContext(ctx).Create(plan).Error
}

func (r *planRepository) CreatePlanningMealItems(ctx context.Context, meals []*model.PlanningMealItem) error {
	return r.db.WithContext(ctx).Create(meals).Error
}

func (r *planRepository) CreateShoppingIngredientItems(ctx context.Context, ingredients []*model.ShoppingIngredientItem) error {
	return r.db.WithContext(ctx).Create(ingredients).Error
}

func (r *planRepository) FindMealsByPlanID(ctx context.Context, planID string) ([]*model.PlanningMealItem, error) {
	var meals []*model.PlanningMealItem
	err := r.db.WithContext(ctx).
		Preload("Menu.MenuIngredientItems.Ingredient"). // IngredientTypeのPreloadを削除
		Where("plan_id = ?", planID).
		Order("date ASC, meal_period ASC").
		Find(&meals).Error
	return meals, err
}

func (r *planRepository) FindShoppingIngredientsByPlanID(ctx context.Context, planID string) ([]*model.ShoppingIngredientItem, error) {
	var ingredients []*model.ShoppingIngredientItem
	err := r.db.WithContext(ctx).
		Preload("Ingredient.IngredientType").
		Where("plan_id = ?", planID).
		Find(&ingredients).Error
	return ingredients, err
}

func (r *planRepository) FindShoppingIngredientItemByID(ctx context.Context, itemID string) (*model.ShoppingIngredientItem, error) {
	var item model.ShoppingIngredientItem
	err := r.db.WithContext(ctx).
		Preload("Ingredient.IngredientType"). // レスポンス生成に必要な情報をPreload
		First(&item, "id = ?", itemID).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *planRepository) UpdateShoppingIngredientItem(ctx context.Context, item *model.ShoppingIngredientItem) error {
	// Saveは全フィールドを更新します。特定のフィールドのみ更新したい場合はUpdateを使用します。
	return r.db.WithContext(ctx).Save(item).Error
}