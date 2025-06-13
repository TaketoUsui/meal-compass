package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config は、アプリケーションの設定を保持する構造体です。
// envconfigタグを使って、対応する環境変数を指定します。
type Config struct {
	// Ginフレームワークの動作モード (debug, release, test)
	GinMode string `envconfig:"GIN_MODE" default:"debug"`

	// バックエンドアプリケーションがリッスンするポート
	GoAppPort string `envconfig:"GO_APP_PORT" default:"8080"`

	// データベース接続設定
	DBHost      string `envconfig:"DB_HOST" required:"true"`
	DBPort      string `envconfig:"DB_PORT" required:"true"`
	DBUser      string `envconfig:"DB_USER" required:"true"`
	DBPassword  string `envconfig:"DB_PASSWORD" required:"true"`
	DBName      string `envconfig:"DB_NAME" required:"true"`
	DBDsnParams string `envconfig:"DB_DSN_PARAMS" required:"true"`
}

// Load は、環境変数を読み込み、Config構造体にマッピングして返します。
func Load() (*Config, error) {
	// .env.localファイルが存在すれば、そちらを優先して読み込む
	// 存在しない場合でもエラーにはならず、次に進む
	_ = godotenv.Load(".env.local")

	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("'.env' file not found, loading from environment variables")
	}

	var cfg Config
	// 環境変数をConfig構造体にマッピング
	// プレフィックスなしで処理するため、第一引数は空文字列
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	log.Println("Configuration loaded successfully")
	return &cfg, nil
}