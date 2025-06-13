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
		{Name: "肉類"},
		{Name: "魚介類"},
		{Name: "野菜"},
		{Name: "果物"},
		{Name: "穀物"},
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
		{"肉類", model.Ingredient{Name: "豚バラ肉", BaseAmount: 200, Unit: "g"}},
		{"肉類", model.Ingredient{Name: "鶏もも肉", BaseAmount: 250, Unit: "g"}},
		{"野菜", model.Ingredient{Name: "玉ねぎ", BaseAmount: 1, Unit: "個"}},
		{"野菜", model.Ingredient{Name: "じゃがいも", BaseAmount: 1, Unit: "個"}},
		{"野菜", model.Ingredient{Name: "人参", BaseAmount: 1, Unit: "本"}},
		{"野菜", model.Ingredient{Name: "キャベツ", BaseAmount: 1, Unit: "玉"}},
		{"野菜", model.Ingredient{Name: "生姜", BaseAmount: 1, Unit: "かけ"}},
		{"穀物", model.Ingredient{Name: "ご飯", BaseAmount: 150, Unit: "g"}},
		{"パン類", model.Ingredient{Name: "食パン", BaseAmount: 6, Unit: "枚"}},
		{"乳製品・卵", model.Ingredient{Name: "バター", BaseAmount: 100, Unit: "g"}},
		{"乳製品・卵", model.Ingredient{Name: "卵", BaseAmount: 10, Unit: "個"}},
		{"調味料", model.Ingredient{Name: "醤油", BaseAmount: 1000, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "みりん", BaseAmount: 500, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "酒", BaseAmount: 500, Unit: "ml"}},
		{"調味料", model.Ingredient{Name: "塩", BaseAmount: 200, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "こしょう", BaseAmount: 50, Unit: "g"}},
		{"調味料", model.Ingredient{Name: "サラダ油", BaseAmount: 1000, Unit: "ml"}},
		{"その他", model.Ingredient{Name: "カレールー", BaseAmount: 1, Unit: "箱"}},
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
		{
			MenuName: "豚の生姜焼き",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚バラ肉", Amount: 150},
				{Name: "玉ねぎ", Amount: 0.5},
				{Name: "生姜", Amount: 0.5},
				{Name: "醤油", Amount: 30},
				{Name: "みりん", Amount: 30},
				{Name: "酒", Amount: 15},
				{Name: "サラダ油", Amount: 10},
			},
		},
		{
			MenuName: "カレーライス",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "豚バラ肉", Amount: 100},
				{Name: "じゃがいも", Amount: 1},
				{Name: "人参", Amount: 0.5},
				{Name: "玉ねぎ", Amount: 0.5},
				{Name: "カレールー", Amount: 1},
				{Name: "ご飯", Amount: 1},
				{Name: "サラダ油", Amount: 10},
			},
		},
		{
			MenuName: "バタートースト",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "食パン", Amount: 1},
				{Name: "バター", Amount: 10},
			},
		},
		{
			MenuName: "目玉焼き",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "卵", Amount: 1},
				{Name: "サラダ油", Amount: 5},
				{Name: "塩", Amount: 0.5},
				{Name: "こしょう", Amount: 0.2},
			},
		},
		{
			MenuName: "親子丼",
			Ingredients: []struct { Name string; Amount float64 }{
				{Name: "鶏もも肉", Amount: 100},
				{Name: "玉ねぎ", Amount: 0.25},
				{Name: "卵", Amount: 2},
				{Name: "醤油", Amount: 20},
				{Name: "みりん", Amount: 20},
				{Name: "ご飯", Amount: 1},
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
			if err := db.First(&ingredient, "name = ?", ingInfo.Name).Error; err != nil { /* ... */ }

			item := model.MenuIngredientItem{
				MenuID:       menu.ID,
				IngredientID: ingredient.ID,
			}
			if err := db.Where(&item).FirstOrCreate(&item).Error; err != nil {
				return err
			}
            if item.Amount != ingInfo.Amount {
                db.Model(&item).Update("amount", ingInfo.Amount)
            }
		}
	}
	return nil
}