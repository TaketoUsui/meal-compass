import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useShoppingPlan } from '../hooks/useShoppingPlan';
import { Button } from '../components/ui/Button';
import { Checkbox } from '../components/ui/Checkbox';
import { MealPeriod } from '../types';
import styles from './TopPage.module.css';

// 選択肢を生成するための設定
const MEAL_PERIODS: { key: MealPeriod; label: string }[] = [
  { key: 'MORNING', label: '朝' },
  { key: 'LUNCH', label: '昼' },
  { key: 'DINNER', label: '夜' },
];
const DAYS_TO_SHOW = 7;

export const TopPage: React.FC = () => {
  const navigate = useNavigate();
  const { createPlan, isLoading, error, shoppingPlanId } = useShoppingPlan();
  
  // { '0-MORNING': true, '1-DINNER': false, ... } のような形式で選択状態を管理
  const [selectedMeals, setSelectedMeals] = useState<Record<string, boolean>>({});

  const handleCheckboxChange = (dateOffset: number, mealPeriod: MealPeriod) => {
    const key = `${dateOffset}-${mealPeriod}`;
    setSelectedMeals(prev => ({ ...prev, [key]: !prev[key] }));
  };

  const handleSubmit = async () => {
    const planned_meals = Object.entries(selectedMeals)
      .filter(([, isSelected]) => isSelected)
      .map(([key]) => {
        const [date_offset, meal_period] = key.split('-');
        return {
          date_offset: parseInt(date_offset, 10),
          meal_period: meal_period as MealPeriod,
        };
      });

    if (planned_meals.length === 0) {
      alert('自炊する食事を1つ以上選択してください。');
      return;
    }

    await createPlan(planned_meals);
  };

  // createPlanが成功し、shoppingPlanIdがセットされたら結果ページに遷移
  useEffect(() => {
    if (shoppingPlanId) {
      navigate(`/plan/${shoppingPlanId}`);
    }
  }, [shoppingPlanId, navigate]);

  // 日付ラベルを生成 (例: "6月13日 (金)")
  const dayLabels = Array.from({ length: DAYS_TO_SHOW }, (_, i) => {
    const date = new Date();
    date.setDate(date.getDate() + i);
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const dayOfWeek = ['日', '月', '火', '水', '木', '金', '土'][date.getDay()];
    const label = i === 0 ? '今日' : i === 1 ? '明日' : `${month}/${day}`;
    return {
      label: `${label} (${dayOfWeek})`,
      offset: i,
    };
  });

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>自炊するごはんを選ぼう</h1>
      <p className={styles.description}>
        下の表から、自炊したい食事にチェックを入れて「計画を作成する」ボタンを押してください。
      </p>

      <div className={styles.selectionGrid}>
        {dayLabels.map(({ label, offset }) => (
          <div key={offset} className={styles.dayColumn}>
            <div className={styles.dayHeader}>{label}</div>
            <div className={styles.mealCheckboxes}>
              {MEAL_PERIODS.map(({ key, label: mealLabel }) => (
                <Checkbox
                  key={key}
                  label={mealLabel}
                  checked={!!selectedMeals[`${offset}-${key}`]}
                  onChange={() => handleCheckboxChange(offset, key)}
                  disabled={isLoading}
                />
              ))}
            </div>
          </div>
        ))}
      </div>

      <div className={styles.actions}>
        {error && <p className={styles.error}>{error}</p>}
        <Button
          onClick={handleSubmit}
          isLoading={isLoading}
          disabled={Object.values(selectedMeals).every(v => !v)}
        >
          計画を作成する
        </Button>
      </div>
    </div>
  );
};