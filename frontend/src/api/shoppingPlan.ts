import apiClient from './apiClient';
import type {
  CreateShoppingPlanRequest,
  ShoppingPlanResponse,
  MenuListResponse,
  IngredientListResponse,
  Ingredient
} from '../types';

/**
 * 新しい買い物計画を作成する
 * @param requestBody - ユーザーが選択した食事の予定
 * @returns 作成された買い物計画の情報
 */
export const createNewPlan = async (requestBody: CreateShoppingPlanRequest): Promise<ShoppingPlanResponse> => {
  const response = await apiClient.post<ShoppingPlanResponse>('/api/create-new-plan', requestBody);
  return response.data;
};

/**
 * 指定されたIDの献立リストを取得する
 * @param shoppingPlanId - 買い物計画のID
 * @returns 献立リスト
 */
export const getMenuList = async (shoppingPlanId: string): Promise<MenuListResponse> => {
  const response = await apiClient.get<MenuListResponse>(`/api/menu-list/${shoppingPlanId}`);
  return response.data;
};

/**
 * 指定されたIDの買い物リストを取得する
 * @param shoppingPlanId - 買い物計画のID
 * @returns 買い物リスト
 */
export const getIngredientList = async (shoppingPlanId: string): Promise<IngredientListResponse> => {
  const response = await apiClient.get<IngredientListResponse>(`/api/ingredient-list/${shoppingPlanId}`);
  return response.data;
};

/**
 * 買い物リストの材料の購入状況を更新する
 * @param itemId - 材料アイテムのID
 * @param bought - 新しい購入状況 (true: 購入済み, false: 未購入)
 * @returns 更新された材料の情報
 */
export const updateIngredientStatus = async (itemId: string, bought: boolean): Promise<Ingredient> => {
  const response = await apiClient.patch<Ingredient>(`/api/shopping_ingredient_items/${itemId}`, { bought });
  return response.data;
};