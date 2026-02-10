package registry

import (
	"fmt"
	"os"

	dt "go-practice/domain/todo"
	"go-practice/infrastructure/database"
	"go-practice/infrastructure/persistence"
	ut "go-practice/usecase/todo"
)

type CLIContainer struct {
	Uncomplete *ut.UncompleteTodoUseCase
}

func NewCLIContainer(store string) (*CLIContainer, error) {
	var repo dt.ITodoRepository

	switch store {
	case "memory":
		repo = persistence.NewInMemoryTodoRepository()
	case "db":
		dbConfig := database.Config{
			Driver:   getEnv("DB_DRIVER", "sqlite"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "todo.db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		}
		db, err := database.NewDB(dbConfig) // ※あなたのプロジェクトのDB初期化に合わせて調整にゃ
		if err != nil {
			return nil, err
		}
		repo = persistence.NewTodoRepository(db)
	default:
		return nil, fmt.Errorf("unknown store: %s (use memory|db)", store)
	}

	return &CLIContainer{
		Uncomplete: ut.NewUncompleteTodoUseCase(repo),
	}, nil
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
