package seeder

import (
	"log"

	"gorm.io/gorm"

	"meal-compass/backend/internal/domain/model"
)

// IngredientsSeeder は、食材分類と食材のマスターデータをデータベースに投入します。
func IngredientsSeeder(db *gorm.DB) error {
	log.Println("Seeding ingredient types...")
	if err := createIngredientTypes(db); err != nil {
		return err
	}

	log.Println("Seeding ingredients...")
	if err := createIngredients(db); err != nil {
		return err
	}

	return nil
}

// MenusSeeder は、メニューと、それに必要な食材（レシピ）のマスターデータをデータベースに投入します。
// 必ず IngredientsSeeder の後に実行する必要があります。
func MenusSeeder(db *gorm.DB) error {
	log.Println("Seeding menus and recipes...")
	return createMenusWithRecipes(db)
}

// --- Private Functions ---

func createIngredientTypes(db *gorm.DB) error {
	types := []model.IngredientType{
		{Name: "生鮮食品"},   // 肉類、魚介類
		{Name: "野菜/果物"}, // 野菜、果物
		{Name: "乾物類"},     // 穀物、麺類、粉物
		{Name: "パン類"},
		{Name: "乳製品・卵"},
		{Name: "調味料"},
		{Name: "その他"},
	}

	for _, t := range types {
		// Nameをキーに、存在しなければ作成
		if err := db.FirstOrCreate(&t, model.IngredientType{Name: t.Name}).Error; err != nil {
			return err
		}
	}
	return nil
}

func createIngredients(db *gorm.DB) error {
	// 型名をキーとして、登録済みのIngredientTypeのIDをマップに保持
	typeMap := make(map[string]string)
	var types []model.IngredientType
	if err := db.Find(&types).Error; err != nil {
		return err
	}
	for _, t := range types {
		typeMap[t.Name] = t.ID
	}

	ingredients := []struct {
		TypeName string
		model.Ingredient
	}{
		// --- 生鮮食品 ---
		{"生鮮食品", model.Ingredient{Name: "豚バラ肉", BaseAmount: 200, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "豚ロース肉", BaseAmount: 200, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "鶏もも肉", BaseAmount: 250, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "鶏むね肉", BaseAmount: 250, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "鶏ひき肉", BaseAmount: 200, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "合いびき肉", BaseAmount: 200, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "牛肉", BaseAmount: 200, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "ベーコン", BaseAmount: 80, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "鮭", BaseAmount: 1, Unit: "切れ"}},
		{"生鮮食品", model.Ingredient{Name: "エビ", BaseAmount: 100, Unit: "g"}},
		{"生鮮食品", model.Ingredient{Name: "アジ", BaseAmount: 1, Unit: "尾"}},
		{"生鮮食品", model.Ingredient{Name: "サバ", BaseAmount: 1, Unit: "切れ"}},
		// --- 野菜/果物 ---
		{"野菜/果物", model.Ingredient{Name: "玉ねぎ", BaseAmount: 3, Unit: "個"}},
		{"野菜/果物", model.Ingredient{Name: "じゃがいも", BaseAmount: 3, Unit: "個"}},
		{"野菜/果物", model.Ingredient{Name: "人参", BaseAmount: 2, Unit: "本"}},
		{"野菜/果物", model.Ingredient{Name: "キャベツ", BaseAmount: 1, Unit: "玉"}},
		{"野菜/果物", model.Ingredient{Name: "ピーマン", BaseAmount: 4, Unit: "個"}},
		{"野菜/果物", model.Ingredient{Name: "なす", BaseAmount: 3, Unit: "本"}},
		{"野菜/果物", model.Ingredient{Name: "トマト", BaseAmount: 3, Unit: "個"}},
		{"野菜/果物", model.Ingredient{Name: "きゅうり", BaseAmount: 3, Unit: "本"}},
		{"野菜/果物", model.Ingredient{Name: "レタス", BaseAmount: 1, Unit: "玉"}},
		{"野菜/果物", model.Ingredient{Name: "大根", BaseAmount: 1, Unit: "本"}},
		{"野菜/果物", model.Ingredient{Name: "長ねぎ", BaseAmount: 1, Unit: "本"}},
		{"野菜/果物", model.Ingredient{Name: "にんにく", BaseAmount: 1, Unit: "玉"}},
		{"野菜/果物", model.Ingredient{Name: "生姜", BaseAmount: 1, Unit: "個"}},
		{"野菜/果物", model.Ingredient{Name: "しめじ", BaseAmount: 1, Unit: "パック"}},
		// --- 乾物類 ---
		{"乾物類", model.Ingredient{Name: "米", BaseAmount: 5000, Unit: "g"}},
		{"乾物類", model.Ingredient{Name: "パスタ", BaseAmount: 500, Unit: "g"}},
		{"乾物類", model.Ingredient{Name: "うどん", BaseAmount: 3, Unit: "玉"}},
		{"乾物類", model.Ingredient{Name: "小麦粉", BaseAmount: 500, Unit: "g"}},
		{"乾物類", model.Ingredient{Name: "片栗粉", BaseAmount: 200, Unit: "g"}},
		{"乾物類", model.Ingredient{Name: "パン粉", BaseAmount: 100, Unit: "g"}},
		// --- パン類 ---
		{"パン類", model.Ingredient{Name: "食パン", BaseAmount: 6, Unit: "枚"}},
		// --- 乳製品・卵 ---
		{"乳製品・卵", model.Ingredient{Name: "卵", BaseAmount: 10, Unit: "個"}},
		{"乳製品・卵", model.Ingredient{Name: "牛乳", BaseAmount: 1000, Unit: "ml"}},
		{"乳製品・卵", model.Ingredient{Name: "バター", BaseAmount: 150, Unit: "g"}},
		{"乳製品・卵", model.Ingredient{Name: "チーズ", BaseAmount: 100, Unit: "g"}},
		// --- 調味料 ---
		{"調味料", model.Ingredient{Name: "醤油", BaseAmount: 1000, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "みりん", BaseAmount: 500, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "酒", BaseAmount: 500, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "酢", BaseAmount: 500, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "味噌", BaseAmount: 750, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "砂糖", BaseAmount: 1000, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "塩", BaseAmount: 200, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "こしょう", BaseAmount: 50, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "サラダ油", BaseAmount: 1000, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "ごま油", BaseAmount: 200, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "オリーブオイル", BaseAmount: 500, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "マヨネーズ", BaseAmount: 500, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "ケチャップ", BaseAmount: 500, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "コンソメ", BaseAmount: 50, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "鶏がらスープの素", BaseAmount: 50, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "豆板醤", BaseAmount: 50, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "オイスターソース", BaseAmount: 120, Unit: "g"}},
		// --- その他 ---
		{"その他", model.Ingredient{Name: "カレールー", BaseAmount: 1, Unit: "箱"}},
		{"その他", model.Ingredient{Name: "豆腐", BaseAmount: 1, Unit: "丁"}},
		{"その他", model.Ingredient{Name: "キムチ", BaseAmount: 200, Unit: "g"}},
	}

	for _, i := range ingredients {
		ing := i.Ingredient
		ing.TypeID = typeMap[i.TypeName]
		if err := db.Where(model.Ingredient{Name: ing.Name}).FirstOrCreate(&ing).Error; err != nil {
			return err
		}
	}
	return nil
}

func createMenusWithRecipes(db *gorm.DB) error {
	// レシピ情報を定義
	recipes := []struct {
		MenuName    string
		Ingredients []struct {
			Name   string
			Amount float64
		}
	}{
		// --- 定番料理 ---
		{
			MenuName: "豚の生姜焼き",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚ロース肉", Amount: 150}, {Name: "玉ねぎ", Amount: 0.5}, {Name: "生姜", Amount: 15}, {Name: "醤油", Amount: 30}, {Name: "みりん", Amount: 30}, {Name: "酒", Amount: 15}, {Name: "サラダ油", Amount: 10},
			},
		},
		{
			MenuName: "カレーライス",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚バラ肉", Amount: 100}, {Name: "じゃがいも", Amount: 1}, {Name: "人参", Amount: 0.5}, {Name: "玉ねぎ", Amount: 0.5}, {Name: "カレールー", Amount: 0.5}, {Name: "米", Amount: 150}, {Name: "サラダ油", Amount: 10},
			},
		},
		{
			MenuName: "親子丼",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "鶏もも肉", Amount: 100}, {Name: "玉ねぎ", Amount: 0.25}, {Name: "卵", Amount: 2}, {Name: "醤油", Amount: 20}, {Name: "みりん", Amount: 20}, {Name: "米", Amount: 150},
			},
		},
		{
			MenuName: "肉じゃが",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "牛肉", Amount: 100}, {Name: "じゃがいも", Amount: 2}, {Name: "人参", Amount: 0.5}, {Name: "玉ねぎ", Amount: 1}, {Name: "醤油", Amount: 45}, {Name: "砂糖", Amount: 20}, {Name: "みりん", Amount: 30},
			},
		},
		{
			MenuName: "鶏の唐揚げ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "鶏もも肉", Amount: 250}, {Name: "醤油", Amount: 30}, {Name: "酒", Amount: 15}, {Name: "にんにく", Amount: 10}, {Name: "生姜", Amount: 10}, {Name: "片栗粉", Amount: 30}, {Name: "サラダ油", Amount: 100},
			},
		},
		{
			MenuName: "ハンバーグ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "合いびき肉", Amount: 200}, {Name: "玉ねぎ", Amount: 0.5}, {Name: "卵", Amount: 1}, {Name: "パン粉", Amount: 20}, {Name: "牛乳", Amount: 30}, {Name: "塩", Amount: 2}, {Name: "こしょう", Amount: 0.5}, {Name: "サラダ油", Amount: 15}, {Name: "ケチャップ", Amount: 30},
			},
		},
		{
			MenuName: "とんかつ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚ロース肉", Amount: 150}, {Name: "小麦粉", Amount: 20}, {Name: "卵", Amount: 1}, {Name: "パン粉", Amount: 30}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5}, {Name: "サラダ油", Amount: 150},
			},
		},
		{
			MenuName: "牛丼",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "牛肉", Amount: 150}, {Name: "玉ねぎ", Amount: 0.5}, {Name: "醤油", Amount: 30}, {Name: "みりん", Amount: 30}, {Name: "砂糖", Amount: 10}, {Name: "酒", Amount: 15}, {Name: "米", Amount: 150},
			},
		},
		{
			MenuName: "豚汁",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚バラ肉", Amount: 80}, {Name: "大根", Amount: 50}, {Name: "人参", Amount: 30}, {Name: "長ねぎ", Amount: 0.25}, {Name: "豆腐", Amount: 0.25}, {Name: "味噌", Amount: 30}, {Name: "ごま油", Amount: 5},
			},
		},
		// --- 中華 ---
		{
			MenuName: "麻婆豆腐",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豆腐", Amount: 1}, {Name: "鶏ひき肉", Amount: 100}, {Name: "長ねぎ", Amount: 0.5}, {Name: "にんにく", Amount: 10}, {Name: "生姜", Amount: 10}, {Name: "豆板醤", Amount: 10}, {Name: "醤油", Amount: 15}, {Name: "鶏がらスープの素", Amount: 5}, {Name: "片栗粉", Amount: 10}, {Name: "ごま油", Amount: 10},
			},
		},
		{
			MenuName: "回鍋肉",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚バラ肉", Amount: 150}, {Name: "キャベツ", Amount: 150}, {Name: "ピーマン", Amount: 1}, {Name: "味噌", Amount: 20}, {Name: "砂糖", Amount: 10}, {Name: "醤油", Amount: 10}, {Name: "豆板醤", Amount: 5}, {Name: "ごま油", Amount: 10},
			},
		},
		{
			MenuName: "青椒肉絲",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "牛肉", Amount: 150}, {Name: "ピーマン", Amount: 2}, {Name: "醤油", Amount: 20}, {Name: "酒", Amount: 10}, {Name: "片栗粉", Amount: 10}, {Name: "オイスターソース", Amount: 15}, {Name: "ごま油", Amount: 10},
			},
		},
		{
			MenuName: "エビチリ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "エビ", Amount: 150}, {Name: "長ねぎ", Amount: 0.5}, {Name: "生姜", Amount: 10}, {Name: "にんにく", Amount: 10}, {Name: "ケチャップ", Amount: 45}, {Name: "豆板醤", Amount: 10}, {Name: "鶏がらスープの素", Amount: 5}, {Name: "片栗粉", Amount: 10},
			},
		},
		{
			MenuName: "チャーハン",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "米", Amount: 180}, {Name: "卵", Amount: 1}, {Name: "長ねぎ", Amount: 0.25}, {Name: "ベーコン", Amount: 20}, {Name: "醤油", Amount: 10}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5}, {Name: "ごま油", Amount: 10},
			},
		},
		{
			MenuName: "豚キムチ炒め",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚バラ肉", Amount: 150}, {Name: "キムチ", Amount: 100}, {Name: "玉ねぎ", Amount: 0.25}, {Name: "醤油", Amount: 5}, {Name: "ごま油", Amount: 10},
			},
		},
		// --- 洋食・パスタ ---
		{
			MenuName: "ミートソースパスタ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "パスタ", Amount: 100}, {Name: "合いびき肉", Amount: 100}, {Name: "玉ねぎ", Amount: 0.25}, {Name: "人参", Amount: 0.25}, {Name: "にんにく", Amount: 10}, {Name: "トマト", Amount: 1}, {Name: "ケチャップ", Amount: 30}, {Name: "コンソメ", Amount: 5}, {Name: "オリーブオイル", Amount: 10},
			},
		},
		{
			MenuName: "カルボナーラ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "パスタ", Amount: 100}, {Name: "ベーコン", Amount: 50}, {Name: "卵", Amount: 2}, {Name: "牛乳", Amount: 50}, {Name: "チーズ", Amount: 30}, {Name: "にんにく", Amount: 10}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5}, {Name: "オリーブオイル", Amount: 10},
			},
		},
		{
			MenuName: "オムライス",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "米", Amount: 150}, {Name: "鶏もも肉", Amount: 50}, {Name: "玉ねぎ", Amount: 0.25}, {Name: "ケチャップ", Amount: 45}, {Name: "卵", Amount: 2}, {Name: "牛乳", Amount: 15}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5}, {Name: "サラダ油", Amount: 10},
			},
		},
		{
			MenuName: "チキングラタン",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "鶏もも肉", Amount: 100}, {Name: "玉ねぎ", Amount: 0.25}, {Name: "しめじ", Amount: 0.5}, {Name: "小麦粉", Amount: 20}, {Name: "牛乳", Amount: 200}, {Name: "バター", Amount: 20}, {Name: "チーズ", Amount: 30}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5},
			},
		},
		{
			MenuName: "焼きうどん",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "うどん", Amount: 1}, {Name: "豚バラ肉", Amount: 50}, {Name: "キャベツ", Amount: 100}, {Name: "人参", Amount: 20}, {Name: "ピーマン", Amount: 0.5}, {Name: "醤油", Amount: 15}, {Name: "みりん", Amount: 10}, {Name: "サラダ油", Amount: 10},
			},
		},
		// --- 魚料理 ---
		{
			MenuName: "鮭の塩焼き",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "鮭", Amount: 1}, {Name: "塩", Amount: 2},
			},
		},
		{
			MenuName: "サバの味噌煮",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "サバ", Amount: 1}, {Name: "生姜", Amount: 10}, {Name: "味噌", Amount: 30}, {Name: "砂糖", Amount: 20}, {Name: "酒", Amount: 30}, {Name: "みりん", Amount: 15},
			},
		},
		{
			MenuName: "アジフライ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "アジ", Amount: 1}, {Name: "小麦粉", Amount: 15}, {Name: "卵", Amount: 0.5}, {Name: "パン粉", Amount: 20}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5}, {Name: "サラダ油", Amount: 100},
			},
		},
		// --- 簡単な一品 ---
		{
			MenuName: "バタートースト",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "食パン", Amount: 1}, {Name: "バター", Amount: 10},
			},
		},
		{
			MenuName: "目玉焼き",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "卵", Amount: 1}, {Name: "サラダ油", Amount: 5}, {Name: "塩", Amount: 0.5}, {Name: "こしょう", Amount: 0.2},
			},
		},
		{
			MenuName: "冷奴",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豆腐", Amount: 0.5}, {Name: "長ねぎ", Amount: 0.1}, {Name: "生姜", Amount: 5}, {Name: "醤油", Amount: 10},
			},
		},
		{
			MenuName: "きゅうりの塩昆布和え",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "きゅうり", Amount: 1}, {Name: "ごま油", Amount: 5},
			},
		},
		{
			MenuName: "トマトサラダ",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "トマト", Amount: 1}, {Name: "玉ねぎ", Amount: 0.1}, {Name: "酢", Amount: 15}, {Name: "オリーブオイル", Amount: 10}, {Name: "塩", Amount: 1}, {Name: "こしょう", Amount: 0.5},
			},
		},
		{
			MenuName: "鶏むね肉のレンジ蒸し",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "鶏むね肉", Amount: 250}, {Name: "酒", Amount: 15}, {Name: "塩", Amount: 2}, {Name: "こしょう", Amount: 0.5},
			},
		},
		{
			MenuName: "無限ピーマン",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "ピーマン", Amount: 3}, {Name: "ベーコン", Amount: 20}, {Name: "鶏がらスープの素", Amount: 3}, {Name: "ごま油", Amount: 5},
			},
		},
	}

	// レシピ情報を元にDBに登録
	for _, r := range recipes {
		menu := model.Menu{Name: r.MenuName}
		if err := db.Where(model.Menu{Name: menu.Name}).FirstOrCreate(&menu).Error; err != nil {
			return err
		}

		for _, ingInfo := range r.Ingredients {
			var ingredient model.Ingredient
			if err := db.First(&ingredient, "name = ?", ingInfo.Name).Error; err != nil {
				log.Printf("Ingredient not found: %s. Skipping for menu: %s. Error: %v", ingInfo.Name, r.MenuName, err)
				continue
			}

			item := model.MenuIngredientItem{
				MenuID:       menu.ID,
				IngredientID: ingredient.ID,
				Amount:       ingInfo.Amount, // Create時にAmountも設定
			}
			// 複合主キー(MenuID, IngredientID)で存在チェック
			var existingItem model.MenuIngredientItem
			result := db.Where("menu_id = ? AND ingredient_id = ?", item.MenuID, item.IngredientID).First(&existingItem)

			if result.Error != nil {
				if result.Error == gorm.ErrRecordNotFound {
					// 存在しないので作成
					if err := db.Create(&item).Error; err != nil {
						return err
					}
				} else {
					// その他のDBエラー
					return result.Error
				}
			} else {
				// 存在する場合、Amountが異なる場合のみ更新
				if existingItem.Amount != item.Amount {
					if err := db.Model(&existingItem).Update("amount", item.Amount).Error; err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
