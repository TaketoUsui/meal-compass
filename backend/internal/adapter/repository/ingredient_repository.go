package repository

import (
	"context"

	"gorm.io/gorm"

	"meal-compass/backend/internal/domain/model"
	"meal-compass/backend/internal/domain/repository"
)

type ingredientRepository struct {
	db *gorm.DB
}

// NewIngredientRepository は新しい ingredientRepository のインスタンスを生成します。
func NewIngredientRepository(db *gorm.DB) repository.IngredientRepository {
	return &ingredientRepository{db: db}
}

func (r *ingredientRepository) CreateIngredientTypes(ctx context.Context, ingredientTypes []*model.IngredientType) error {
	// GORMのCreateにスライスを渡すことで、バルクインサートが実行されます。
	return r.db.WithContext(ctx).Create(ingredientTypes).Error
}

func (r *ingredientRepository) CreateIngredients(ctx context.Context, ingredients []*model.Ingredient) error {
	return r.db.WithContext(ctx).Create(ingredients).Error
}