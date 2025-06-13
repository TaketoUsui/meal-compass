import axios from 'axios';

// Viteの機能を使って環境変数からAPIのベースURLを取得
const VITE_API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

// axiosのインスタンスを作成
const apiClient = axios.create({
  // 環境変数で設定されたAPIのベースURL
  // 例: http://localhost:8080
  baseURL: VITE_API_BASE_URL,
  // デフォルトのヘッダー設定
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
  // リクエストタイムアウトをミリ秒で設定 (例: 10秒)
  timeout: 10000,
});

// エラーハンドリングの共通化などもここに記述可能
apiClient.interceptors.response.use(
  response => response,
  error => {
    // ネットワークエラーやタイムアウトなどのハンドリング
    console.error('API Error:', error.response?.data || error.message);
    // エラーを呼び出し元に伝播させる
    return Promise.reject(error);
  }
);


export default apiClient;