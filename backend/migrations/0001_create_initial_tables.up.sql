-- ----------------------------------------------------------------
-- ingredient_types: 食材の分類（例：野菜、肉類、調味料）を管理
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `ingredient_types` (
  `id` CHAR(36) NOT NULL COMMENT '食材分類ID (UUID)',
  `name` VARCHAR(255) NOT NULL COMMENT '分類名',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------------------------------------------
-- ingredients: 個別の食材情報を管理
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `ingredients` (
  `id` CHAR(36) NOT NULL COMMENT '食材ID (UUID)',
  `type_id` CHAR(36) NOT NULL COMMENT '食材分類ID',
  `name` VARCHAR(255) NOT NULL COMMENT '食材名',
  `base_amount` DECIMAL(10, 2) NOT NULL COMMENT '基本量（例：1パックあたりの量）',
  `unit` VARCHAR(50) NOT NULL COMMENT '単位（例：g, ml, 個）',
  `shelf_life_days_unopened` INT DEFAULT NULL COMMENT '賞味期限（未開封・日数）',
  `shelf_life_days_opened` INT DEFAULT NULL COMMENT '賞味期限（開封後・日数）',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_name` (`name`),
  FOREIGN KEY (`type_id`) REFERENCES `ingredient_types` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------------------------------------------
-- menus: 料理のメニュー情報を管理
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `menus` (
  `id` CHAR(36) NOT NULL COMMENT 'メニューID (UUID)',
  `name` VARCHAR(255) NOT NULL COMMENT 'メニュー名',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------------------------------------------
-- menu_ingredient_items: メニューと食材の中間テーブル
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `menu_ingredient_items` (
  `id` CHAR(36) NOT NULL COMMENT 'ID (UUID)',
  `menu_id` CHAR(36) NOT NULL COMMENT 'メニューID',
  `ingredient_id` CHAR(36) NOT NULL COMMENT '食材ID',
  `amount` DECIMAL(10, 2) NOT NULL COMMENT '必要量（1人前）',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_menu_ingredient` (`menu_id`, `ingredient_id`),
  FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------------------------------------------
-- shopping_plans: 買い物計画全体を管理
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `shopping_plans` (
  `id` CHAR(36) NOT NULL COMMENT '買い物計画ID (UUID)',
  `period_start_at` DATE NOT NULL COMMENT '計画の開始日',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------------------------------------------
-- planning_meal_items: 計画された個々の食事（どの日に何を食べるか）を管理
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `planning_meal_items` (
  `id` CHAR(36) NOT NULL COMMENT '計画食事ID (UUID)',
  `plan_id` CHAR(36) NOT NULL COMMENT '買い物計画ID',
  `menu_id` CHAR(36) NOT NULL COMMENT 'メニューID',
  `date` DATE NOT NULL COMMENT '食事の日付',
  `meal_period` ENUM('MORNING', 'LUNCH', 'DINNER') NOT NULL COMMENT '食事の時間帯',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`plan_id`) REFERENCES `shopping_plans` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------------------------------------------
-- shopping_ingredient_items: 買い物リストの個々のアイテムを管理
-- ----------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `shopping_ingredient_items` (
  `id` CHAR(36) NOT NULL COMMENT '買い物アイテムID (UUID)',
  `plan_id` CHAR(36) NOT NULL COMMENT '買い物計画ID',
  `ingredient_id` CHAR(36) NOT NULL COMMENT '食材ID',
  `amount` DECIMAL(10, 2) NOT NULL COMMENT '購入量',
  `bought` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '購入済みフラグ',
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時',
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_plan_ingredient` (`plan_id`, `ingredient_id`),
  FOREIGN KEY (`plan_id`) REFERENCES `shopping_plans` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`ingredient_id`) REFERENCES `ingredients` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;