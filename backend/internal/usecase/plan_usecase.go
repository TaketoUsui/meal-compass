package usecase

import (
	"context"
	"fmt"
	"time"

	"meal-compass/backend/internal/domain/model"
	"meal-compass/backend/internal/domain/repository"
)

// --- DTO (Data Transfer Object) Definitions ---
// UsecaseのInput/Outputとして使用する構造体。APIのI/Oに近しい形となる。

type PlannedMealInput struct {
	DateOffset int
	MealPeriod string
}

type CreatePlanOutput struct {
	ShoppingPlanID string                  `json:"shopping_plan_id"`
	Meals          []*MenuOutput           `json:"meals"`
	Ingredients    []*IngredientListOutput `json:"ingredients"`
}

type MenuOutput struct {
	Date         string                `json:"date"`
	MealPeriod   string                `json:"meal_period"`
	MenuName     string                `json:"menu_name"`
	Ingredients  []*MenuIngredientInfo `json:"ingredients"`
}

type MenuIngredientInfo struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type IngredientListOutput struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
	Bought bool    `json:"bought"`
}

type UpdateShoppingIngredientItemInput struct {
	ItemID string
	Bought bool
}

// --- Usecase Interface ---

// PlanUsecase は、計画に関するビジネスロジックのインターフェースです。
type PlanUsecase interface {
	CreatePlan(ctx context.Context, input []PlannedMealInput) (*CreatePlanOutput, error)
	GetMenuList(ctx context.Context, planID string) ([]*MenuOutput, error)
	GetIngredientList(ctx context.Context, planID string) ([]*IngredientListOutput, error)
	UpdateShoppingIngredientItem(ctx context.Context, input UpdateShoppingIngredientItemInput) (*IngredientListOutput, error)
}

// --- Usecase Implementation ---

// planUsecase は PlanUsecase インターフェースの実装です。
type planUsecase struct {
	planRepo       repository.PlanRepository
	menuRepo       repository.MenuRepository
	ingredientRepo repository.IngredientRepository
}

// NewPlanUsecase は新しい planUsecase のインスタンスを生成します。
func NewPlanUsecase(planRepo repository.PlanRepository, menuRepo repository.MenuRepository, ingredientRepo repository.IngredientRepository) PlanUsecase {
	return &planUsecase{
		planRepo:       planRepo,
		menuRepo:       menuRepo,
		ingredientRepo: ingredientRepo,
	}
}

// CreatePlan は、新しい食事計画を作成する中心的なビジネスロジックです。
func (u *planUsecase) CreatePlan(ctx context.Context, input []PlannedMealInput) (*CreatePlanOutput, error) {
	mealCount := len(input)
	if mealCount == 0 {
		return nil, fmt.Errorf("自炊する食事が指定されていません")
	}
	menus, err := u.menuRepo.FindRandomMenus(ctx, mealCount)
	if err != nil {
		return nil, fmt.Errorf("メニューの取得に失敗しました: %w", err)
	}
	if len(menus) < mealCount {
		return nil, fmt.Errorf("十分な数のメニューが登録されていません")
	}

	shoppingListItems := make(map[string]*model.ShoppingIngredientItem)
	// レスポンス生成用に、集計した食材の完全なモデル情報も保持
	ingredientMap := make(map[string]*model.Ingredient) 

	for _, menu := range menus {
		for _, item := range menu.MenuIngredientItems {
			ingredientMap[item.IngredientID] = &item.Ingredient
			if existingItem, ok := shoppingListItems[item.IngredientID]; ok {
				existingItem.Amount += item.Amount
			} else {
				shoppingListItems[item.IngredientID] = &model.ShoppingIngredientItem{
					IngredientID: item.IngredientID,
					Amount:       item.Amount,
					Bought:       false,
				}
			}
		}
	}

	var newPlan model.ShoppingPlan
	var newMeals []*model.PlanningMealItem
	var newIngredients []*model.ShoppingIngredientItem

	for _, item := range shoppingListItems {
		// マップに保持したIngredientモデルを関連付ける
		item.Ingredient = *ingredientMap[item.IngredientID]
		newIngredients = append(newIngredients, item)
	}

	err = u.planRepo.Transaction(ctx, func(txRepo repository.PlanRepository) error {
		newPlan = model.ShoppingPlan{PeriodStartAt: time.Now()}
		if err := txRepo.CreateShoppingPlan(ctx, &newPlan); err != nil { return err }

		for i, mealInput := range input {
			newMeals = append(newMeals, &model.PlanningMealItem{
				PlanID: newPlan.ID,
				MenuID: menus[i].ID,
				Date: time.Now().AddDate(0, 0, mealInput.DateOffset),
				MealPeriod: model.MealPeriod(mealInput.MealPeriod),
				Menu: *menus[i],
			})
		}
		if err := txRepo.CreatePlanningMealItems(ctx, newMeals); err != nil { return err }

		for _, ing := range newIngredients {
			ing.PlanID = newPlan.ID
		}
		if err := txRepo.CreateShoppingIngredientItems(ctx, newIngredients); err != nil { return err }
		
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("計画の保存に失敗しました: %w", err)
	}

	// DBから再取得せず、作成したモデルからレスポンスを生成
	return &CreatePlanOutput{
		ShoppingPlanID: newPlan.ID,
		Meals:          toMenuOutput(newMeals),
		Ingredients:    toIngredientListOutput(newIngredients),
	}, nil
}


// GetMenuList は、指定された計画IDのメニューリストを取得します。
func (u *planUsecase) GetMenuList(ctx context.Context, planID string) ([]*MenuOutput, error) {
	meals, err := u.planRepo.FindMealsByPlanID(ctx, planID)
	if err != nil {
		return nil, err
	}
	return toMenuOutput(meals), nil
}

// GetIngredientList は、指定された計画IDの買い物リストを取得します。
func (u *planUsecase) GetIngredientList(ctx context.Context, planID string) ([]*IngredientListOutput, error) {
	ingredients, err := u.planRepo.FindShoppingIngredientsByPlanID(ctx, planID)
	if err != nil {
		return nil, err
	}
	return toIngredientListOutput(ingredients), nil
}

// UpdateShoppingIngredientItem は、買い物リストのアイテムの購入済み状態を更新します。
func (u *planUsecase) UpdateShoppingIngredientItem(ctx context.Context, input UpdateShoppingIngredientItemInput) (*IngredientListOutput, error) {
	// 1. 更新対象のアイテムを、レスポンスに必要な関連情報を含めて取得します。
	//    repository側でIngredientとIngredientTypeがPreloadされています。
	item, err := u.planRepo.FindShoppingIngredientItemByID(ctx, input.ItemID)
	if err != nil {
		// このエラーはhandlerで gorm.ErrRecordNotFound として扱われます
		return nil, err
	}

	// 2. 状態を更新します。
	item.Bought = input.Bought

	// 3. データベースに保存します。
	if err := u.planRepo.UpdateShoppingIngredientItem(ctx, item); err != nil {
		return nil, fmt.Errorf("アイテムの更新に失敗しました: %w", err)
	}

	// 4. 再取得は不要。更新したモデルオブジェクトを直接DTOに変換して返します。
	//    これにより、不要なDBアクセスがなくなり、ロジックもシンプルになります。
	return toSingleIngredientListOutput(item), nil
}


// --- DTO Converters ---
// ドメインモデルからOutput用のDTOへ変換するヘルパー関数

func toMenuOutput(meals []*model.PlanningMealItem) []*MenuOutput {
	output := make([]*MenuOutput, len(meals))
	for i, meal := range meals {
		ingredientsInfo := make([]*MenuIngredientInfo, len(meal.Menu.MenuIngredientItems))
		for j, item := range meal.Menu.MenuIngredientItems {
			ingredientsInfo[j] = &MenuIngredientInfo{
				Name:   item.Ingredient.Name,
				Amount: item.Amount,
				Unit:   item.Ingredient.Unit,
			}
		}
		output[i] = &MenuOutput{
			Date:         meal.Date.Format("2006-01-02"),
			MealPeriod:   string(meal.MealPeriod),
			MenuName:     meal.Menu.Name,
			Ingredients:  ingredientsInfo,
		}
	}
	return output
}

func toIngredientListOutput(ingredients []*model.ShoppingIngredientItem) []*IngredientListOutput {
	output := make([]*IngredientListOutput, len(ingredients))
	for i, ing := range ingredients {
		output[i] = &IngredientListOutput{
			ID:     ing.ID,
			Name:   ing.Ingredient.Name,
			Type:   ing.Ingredient.IngredientType.Name,
			Amount: ing.Amount,
			Unit:   ing.Ingredient.Unit,
			Bought: ing.Bought,
		}
	}
	return output
}

func toSingleIngredientListOutput(ing *model.ShoppingIngredientItem) *IngredientListOutput {
    return &IngredientListOutput{
        ID:     ing.ID,
        Name:   ing.Ingredient.Name,
        Type:   ing.Ingredient.IngredientType.Name,
        Amount: ing.Amount,
        Unit:   ing.Ingredient.Unit,
        Bought: ing.Bought,
    }
}