package main

import (
	"fmt"
	"go-practice/infrastructure/database"
	"go-practice/infrastructure/persistence"
	"go-practice/interface/router"
	"go-practice/registry"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 環境変数を読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// データベース接続設定
	dbConfig := database.Config{
		Driver:   getEnv("DB_DRIVER", "sqlite"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvAsInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "todo.db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// データベース接続
	db, err := database.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// マイグレーション実行
	if err := db.AutoMigrate(&persistence.TodoModel{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")

	// Registry作成(依存性注入のコンテナ)
	reg := registry.NewRegistry(db)
	defer reg.Close()

	// Handlerを生成
	todoHandler := reg.NewTodoHandler()

	// ルーター設定
	r := router.SetupRouter(todoHandler)

	// サーバー起動
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv は環境変数を取得(デフォルト値あり)
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt は環境変数を整数として取得
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intValue int
		if _, err := fmt.Sscan(value, &intValue); err == nil {
			return intValue
		}
	}
	return defaultValue
}
