package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"meal-compass/backend/internal/config"
)

// NewDB は、設定情報に基づいて新しいデータベース接続を確立し、GORMのDBインスタンスを返します。
func NewDB(cfg *config.Config) (*gorm.DB, error) {
	// DSN (Data Source Name) を構築
	// 例: "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBDsnParams,
	)

	// GORMのロガー設定
	var gormLogger logger.Interface
	if cfg.GinMode == "debug" {
		// 開発環境では、実行されたSQLがすべてログに出力されるようにする
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 本番環境では、エラー発生時のみログに出力する
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	// データベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("データベースへの接続に失敗しました: %w", err)
	}

	// 接続確認
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("DBインスタンスの取得に失敗しました: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("データベースへのPingに失敗しました: %w", err)
	}
	
	return db, nil
}