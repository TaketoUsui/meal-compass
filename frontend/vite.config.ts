import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // Dockerコンテナの外部からアクセスを許可するために '0.0.0.0' ではなく 'true' を指定します。
    host: true,
    // 開発サーバーがリッスンするポート番号。
    // compose.ymlでホスト側のポートにマッピングされます。
    port: 5173,
    // ホットリロードを有効にするためのポーリング設定
    watch: {
      usePolling: true,
    }
  }
});
