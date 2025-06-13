/**
 * 食事の期間を表す型
 * 仕様書に基づき、"MORNING", "LUNCH", "DINNER" のいずれかを取る
 */
export type MealPeriod = "MORNING" | "LUNCH" | "DINNER";

/**
 * 献立に含まれる食材の型
 */
export interface MenuIngredient {
  name: string;
  amount: number;
  unit: string;
}

/**
 * 献立（食事）の型
 * APIレスポンスの "meals" 配列の要素に対応
 */
export interface Meal {
  date: string; // "YYYY-MM-DD" 形式
  meal_period: MealPeriod;
  menu_name: string;
  ingredients: MenuIngredient[];
}

/**
 * 買い物リストの材料の型
 * APIレスポンスの "ingredients" 配列の要素に対応
 */
export interface Ingredient {
  id: string; // UUID
  name: string;
  type: string;
  amount: number;
  unit: string;
  bought: boolean;
}

// --- API Request Types ---

/**
 * 買い物計画作成API (POST /api/create-new-plan) のリクエストBodyの型
 */
export interface CreateShoppingPlanRequest {
  planned_meals: {
    date_offset: number;
    meal_period: MealPeriod;
  }[];
}

/**
 * 材料の購入状況更新API (PATCH /api/shopping_ingredient_items/{item_id}) のリクエストBodyの型
 */
export interface UpdateIngredientRequest {
  bought: boolean;
}


// --- API Response Types ---

/**
 * 買い物計画作成API (POST /api/create-new-plan) のレスポンスの型
 */
export interface ShoppingPlanResponse {
  shopping_plan_id: string;
  meals: Meal[];
  ingredients: Ingredient[];
}

/**
 * 献立リスト取得API (GET /api/menu-list/{shopping_plan_id}) のレスポンスの型
 */
export interface MenuListResponse {
  meals: Meal[];
}

/**
 * 買い物リスト取得API (GET /api/ingredient-list/{shopping_plan_id}) のレスポンスの型
 */
export interface IngredientListResponse {
  ingredients: Ingredient[];
}