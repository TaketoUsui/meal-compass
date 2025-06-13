import React, { useEffect, useMemo } from 'react';
import { useParams, Link } from 'react-router-dom';
import { useShoppingPlan } from '../hooks/useShoppingPlan';
import { Checkbox } from '../components/ui/Checkbox';
import { formatToJapaneseDate } from '../utils/dateFormatter';
import styles from './ResultPage.module.css';
import { Meal, Ingredient } from '../types';

export const ResultPage: React.FC = () => {
  const { shoppingPlanId: planIdFromUrl } = useParams<{ shoppingPlanId: string }>();
  const {
    shoppingPlanId: planIdFromContext,
    meals,
    ingredients,
    isLoading,
    error,
    fetchPlanData,
    toggleIngredientBought,
  } = useShoppingPlan();

  useEffect(() => {
    // URLのIDが存在し、ContextのIDと異なる場合 (直接アクセス/リロード時) はデータを取得
    if (planIdFromUrl && planIdFromUrl !== planIdFromContext) {
      fetchPlanData(planIdFromUrl);
    }
  }, [planIdFromUrl, planIdFromContext, fetchPlanData]);

  // 献立を日付でグループ化する
  const groupedMeals = useMemo(() => {
    return meals.reduce<Record<string, Meal[]>>((acc, meal) => {
      const date = meal.date;
      if (!acc[date]) {
        acc[date] = [];
      }
      acc[date].push(meal);
      return acc;
    }, {});
  }, [meals]);

  // 買い物リストの材料を種類(type)でグループ化する
  const groupedIngredients = useMemo(() => {
    return ingredients.reduce<Record<string, Ingredient[]>>((acc, ingredient) => {
      const type = ingredient.type || 'その他';
      if (!acc[type]) {
        acc[type] = [];
      }
      acc[type].push(ingredient);
      return acc;
    }, {});
  }, [ingredients]);

  if (isLoading) {
    return <div className={styles.status}>読み込んでいます...</div>;
  }

  if (error) {
    return <div className={styles.statusError}>{error} <Link to="/">トップに戻る</Link></div>;
  }
  
  if (!planIdFromContext && !isLoading) {
      return <div className={styles.status}>計画が見つかりません。 <Link to="/">トップページから計画を作成してください。</Link></div>
  }

  return (
    <div className={styles.container}>
      <div className={styles.section}>
        <h2 className={styles.sectionTitle}>あなたの献立</h2>
        {Object.entries(groupedMeals).map(([date, dailyMeals]) => (
          <div key={date} className={styles.dateGroup}>
            <h3 className={styles.dateHeader}>{formatToJapaneseDate(date)}</h3>
            {dailyMeals.map((meal, index) => (
              <div key={index} className={styles.mealCard}>
                <p className={styles.mealPeriod}>{meal.meal_period === 'MORNING' ? '朝' : meal.meal_period === 'LUNCH' ? '昼' : '夜'}ごはん</p>
                <p className={styles.menuName}>{meal.menu_name}</p>
                <ul className={styles.mealIngredients}>
                  {meal.ingredients.map((ing, i) => (
                    <li key={i}>{ing.name} ({ing.amount}{ing.unit})</li>
                  ))}
                </ul>
              </div>
            ))}
          </div>
        ))}
      </div>

      <div className={styles.section}>
        <h2 className={styles.sectionTitle}>買い物リスト</h2>
        <div className={styles.ingredientList}>
          {/* グループ化されたオブジェクトを元に描画 */}
          {Object.entries(groupedIngredients).map(([type, items]) => (
            <div key={type} className={styles.ingredientGroup}>
              <h3 className={styles.ingredientTypeTitle}>{type}</h3>
              <div className={styles.ingredientItems}>
                {items.map((ingredient) => (
                  <Checkbox
                    key={ingredient.id}
                    label={
                      <span className={ingredient.bought ? styles.boughtItem : ''}>
                        {ingredient.name} ({ingredient.amount}{ingredient.unit})
                      </span>
                    }
                    checked={ingredient.bought}
                    onChange={() => toggleIngredientBought(ingredient.id, ingredient.bought)}
                  />
                ))}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};