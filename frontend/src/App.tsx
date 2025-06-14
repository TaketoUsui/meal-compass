// src/App.tsx

import React from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import { useShoppingPlan } from './hooks/useShoppingPlan';

// レイアウトコンポーネントとページコンポーネントをインポート
import { Header } from './components/layout/Header';
import { Footer } from './components/layout/Footer';
import { TopPage } from './pages/TopPage';
import { ResultPage } from './pages/ResultPage';

const ProtectedResultRoute: React.FC = () => {
  const { shoppingPlanId } = useShoppingPlan();

  // Contextに計画IDが存在すればResultPageを表示し、
  // 存在しなければトップページにリダイレクトする
  return shoppingPlanId ? <ResultPage /> : <Navigate to="/" replace />;
};

// アプリケーション全体のレイアウトとルーティングを定義するコンポーネント
const App: React.FC = () => {
  return (
    // #root要素に display: flex; flex-direction: column; が設定されているため、
    // この構造でヘッダー・メイン・フッターのレイアウトが実現される
    <>
      <Header />
      <main>
        {/*
         * Routesコンポーネントは、現在のURLにマッチする最初のRouteを
         * レンダリングします。
        */}
        <Routes>
          {/* ルートパス ("/") にはTopPageコンポーネントを割り当て */}
          <Route path="/" element={<TopPage />} />
          
          {/*
           * "/plan/:shoppingPlanId" のパスにはResultPageコンポーネントを割り当て。
           * ":shoppingPlanId" はURLパラメータとして扱われ、
           * ResultPage内で useParams フックを使って取得できます。
          */}
          <Route path="/plan/:shoppingPlanId" element={<ProtectedResultRoute />} />
        </Routes>
      </main>
      <Footer />
    </>
  );
};

export default App;
