package repository

import (
	"context"

	"meal-compass/backend/internal/domain/model"
)

// IngredientRepository は、食材マスターに関連する永続化を担当するリポジトリです。
type IngredientRepository interface {
	// CreateIngredientTypes は、複数の食材分類を保存します。Seederでの利用を想定しています。
	CreateIngredientTypes(ctx context.Context, ingredientTypes []*model.IngredientType) error
	// CreateIngredients は、複数の食材を保存します。Seederでの利用を想定しています。
	CreateIngredients(ctx context.Context, ingredients []*model.Ingredient) error
}