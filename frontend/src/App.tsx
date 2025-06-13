// src/App.tsx

import React from 'react';
import { Routes, Route } from 'react-router-dom';

// レイアウトコンポーネントとページコンポーネントをインポート
import { Header } from './components/layout/Header';
import { Footer } from './components/layout/Footer';
import { TopPage } from './pages/TopPage';
import { ResultPage } from './pages/ResultPage';

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
          <Route path="/plan/:shoppingPlanId" element={<ResultPage />} />
        </Routes>
      </main>
      <Footer />
    </>
  );
};

export default App;
