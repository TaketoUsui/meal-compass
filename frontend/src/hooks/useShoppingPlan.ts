import { useContext } from 'react';
import {
  ShoppingPlanContext,
  ShoppingPlanContextType,
} from '../contexts/ShoppingPlanContext';

/**
 * ShoppingPlanContextから値を取得するためのカスタムフック。
 *
 * このフックは、コンポーネントが必ず ShoppingPlanProvider の子孫であることを保証します。
 * もしそうでなければ、開発中にエラーをスローして問題を知らせます。
 *
 * @returns {ShoppingPlanContextType} shoppingPlanId, meals, ingredients, isLoading, error, および各種操作関数を含むオブジェクト。
 */
export const useShoppingPlan = (): ShoppingPlanContextType => {
  // useContextフックを使って、Contextから現在の値を取得
  const context = useContext(ShoppingPlanContext);

  // contextがundefinedの場合、Providerの外でフックが使われていることを意味する
  if (context === undefined) {
    throw new Error(
      'useShoppingPlan must be used within a ShoppingPlanProvider'
    );
  }

  return context;
};