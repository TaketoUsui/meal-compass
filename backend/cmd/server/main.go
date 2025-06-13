package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"meal-compass/backend/internal/adapter/handler"
	"meal-compass/backend/internal/adapter/repository"
	"meal-compass/backend/internal/config"
	"meal-compass/backend/internal/seeder"
	"meal-compass/backend/internal/usecase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}
	log.Println("データベースへの接続に成功しました。")

	if cfg.GinMode == "debug" {
		if err := runSeeder(db); err != nil {
			log.Fatalf("ダミーデータの投入に失敗しました: %v", err)
		}
	}

	planRepo := repository.NewPlanRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	ingredientRepo := repository.NewIngredientRepository(db)

	planUsecase := usecase.NewPlanUsecase(planRepo, menuRepo, ingredientRepo)

	planHandler := handler.NewPlanHandler(planUsecase)
	ingredientHandler := handler.NewIngredientHandler(planUsecase)

	router := handler.NewRouter(planHandler, ingredientHandler)

	port := os.Getenv("GO_APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("サーバーを起動します: http://localhost:%s\n", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}

// runSeeder はダミーデータ投入処理を実行します。
func runSeeder(db *gorm.DB) error {
	// 依存関係を考慮し、Ingredient -> Menu の順で実行
	if err := seeder.IngredientsSeeder(db); err != nil {
		return fmt.Errorf("IngredientsSeederの実行に失敗: %w", err)
	}
	if err := seeder.MenusSeeder(db); err != nil {
		return fmt.Errorf("MenusSeederの実行に失敗: %w", err)
	}

	log.Println("ダミーデータの投入が正常に完了しました。")
	return nil
}