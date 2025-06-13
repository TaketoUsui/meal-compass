package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter は、ハンドラーを受け取り、Ginのルーターエンジンをセットアップして返します。
func NewRouter(planHandler *PlanHandler, ingredientHandler *IngredientHandler) *gin.Engine {
	// gin.Default() は Logger と Recovery ミドルウェアを搭載したルーターを生成します
	router := gin.Default()

	// CORS (Cross-Origin Resource Sharing) の設定
	// フロントエンドのURL (Viteのデフォルト開発サーバー) からのアクセスを許可します
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"} // フロントエンドのオリジン
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// ヘルスチェック用のエンドポイント
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// APIのルートグループ
	api := router.Group("/api")
	{
		// 計画作成
		api.POST("/create-new-plan", planHandler.CreateNewPlan)

		// メニューリスト取得
		api.GET("/menu-list/:shopping_plan_id", planHandler.GetMenuList)

		// 買い物リスト取得
		api.GET("/ingredient-list/:shopping_plan_id", planHandler.GetIngredientList)

		// 買い物リストのアイテム更新 (購入済みチェック)
		api.PATCH("/shopping_ingredient_items/:item_id", ingredientHandler.UpdateShoppingIngredientItem)
	}

	return router
}