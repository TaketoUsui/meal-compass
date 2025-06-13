// src/main.tsx

import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import App from './App.tsx';
import { ShoppingPlanProvider } from './contexts/ShoppingPlanContext.tsx';

// アプリケーション全体に適用するグローバルスタイルをインポート
import './assets/styles/global.css';

// DOMのルート要素を取得
const rootElement = document.getElementById('root');
if (!rootElement) {
  throw new Error("Failed to find the root element with id 'root'");
}

// Reactアプリケーションをレンダリング
ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    {/*
     * ブラウザのURLとUIを同期させるためにBrowserRouterでラップします。
     * これにより、内部のコンポーネントでルーティング機能が使用可能になります。
    */}
    <BrowserRouter>
      {/*
       * アプリケーション全体で買い物計画の状態を共有するために
       * ShoppingPlanProviderでラップします。
      */}
      <ShoppingPlanProvider>
        <App />
      </ShoppingPlanProvider>
    </BrowserRouter>
  </React.StrictMode>,
);
