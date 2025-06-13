import React, { createContext, useState, ReactNode, useCallback } from 'react';
import { Meal, Ingredient, CreateShoppingPlanRequest } from '../types';
import * as shoppingPlanApi from '../api/shoppingPlan';

// Contextが提供する値の型定義をexportする
export interface ShoppingPlanContextType {
  shoppingPlanId: string | null;
  meals: Meal[];
  ingredients: Ingredient[];
  isLoading: boolean;
  error: string | null;
  createPlan: (plannedMeals: CreateShoppingPlanRequest['planned_meals']) => Promise<void>;
  fetchPlanData: (planId: string) => Promise<void>;
  toggleIngredientBought: (itemId: string, currentStatus: boolean) => Promise<void>;
}

// Contextオブジェクトの作成 (非公開)
export const ShoppingPlanContext = createContext<ShoppingPlanContextType | undefined>(undefined);

// Contextを提供するためのProviderコンポーネント
export const ShoppingPlanProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  // ... (内部の実装は変更なし)
  const [shoppingPlanId, setShoppingPlanId] = useState<string | null>(null);
  const [meals, setMeals] = useState<Meal[]>([]);
  const [ingredients, setIngredients] = useState<Ingredient[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const createPlan = useCallback(async (plannedMeals: CreateShoppingPlanRequest['planned_meals']) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await shoppingPlanApi.createNewPlan({ planned_meals: plannedMeals });
      setShoppingPlanId(response.shopping_plan_id);
      setMeals(response.meals);
      setIngredients(response.ingredients);
    } catch (err) {
      setError('計画の作成に失敗しました。');
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  }, []);

  const fetchPlanData = useCallback(async (planId: string) => {
    setIsLoading(true);
    setError(null);
    try {
      const [menuResponse, ingredientResponse] = await Promise.all([
        shoppingPlanApi.getMenuList(planId),
        shoppingPlanApi.getIngredientList(planId),
      ]);
      setShoppingPlanId(planId);
      setMeals(menuResponse.meals);
      setIngredients(ingredientResponse.ingredients);
    } catch (err) {
      setError('データの取得に失敗しました。');
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  }, []);

  const toggleIngredientBought = useCallback(async (itemId: string, currentStatus: boolean) => {
    const originalIngredients = [...ingredients];
    const newIngredients = ingredients.map(ing =>
      ing.id === itemId ? { ...ing, bought: !currentStatus } : ing
    );
    setIngredients(newIngredients);

    try {
      await shoppingPlanApi.updateIngredientStatus(itemId, !currentStatus);
    } catch (err) {
      setError('更新に失敗しました。');
      console.error(err);
      setIngredients(originalIngredients);
    }
  }, [ingredients]);
  
  const value = {
    shoppingPlanId,
    meals,
    ingredients,
    isLoading,
    error,
    createPlan,
    fetchPlanData,
    toggleIngredientBought,
  };

  return (
    <ShoppingPlanContext.Provider value={value}>
      {children}
    </ShoppingPlanContext.Provider>
  );
};