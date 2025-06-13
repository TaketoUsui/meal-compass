# Stage 1: ビルド環境 (Builder)
# Node.jsのLTS (Long Term Support) バージョンをベースイメージとして使用します。
FROM node:20-alpine AS builder

# compose.yml の build.args からビルド時引数を受け取る
ARG VITE_API_BASE_URL
# 受け取った引数をビルドプロセスで利用可能な環境変数として設定する
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL

# 作業ディレクトリを設定します。
WORKDIR /app

# まず、依存関係の定義ファイルのみをコピーします。
# これにより、依存関係に変更がない限り、npm installのレイヤーはキャッシュが利用され、ビルドが高速化されます。
COPY package*.json ./

# 依存関係をインストールします。
RUN npm install

# アプリケーションのソースコードをすべてコピーします。
COPY . .

# /app ディレクトリに index.html が存在するか確認する
RUN ls -la

# TypeScriptの型チェックとViteによる本番ビルドを実行します。
# 生成された静的ファイルは /app/dist ディレクトリに格納されます。
RUN npm run build

# Stage 2: 実行環境 (Runner)
# 軽量なWebサーバーであるNginxの安定版をベースイメージとして使用します。
FROM nginx:1.25-alpine

# ビルダー環境からビルド成果物 (distディレクトリの中身) を
# Nginxが静的ファイルを配信するためのデフォルトのディレクトリにコピーします。
COPY --from=builder /app/dist /usr/share/nginx/html

# (任意) カスタムのNginx設定ファイルをコピーしたい場合は、以下の行を有効にします。
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# コンテナがリッスンするポートを80番に設定します。
EXPOSE 80

# Nginxサーバーをフォアグラウンドで起動します。
CMD ["nginx", "-g", "daemon off;"]
