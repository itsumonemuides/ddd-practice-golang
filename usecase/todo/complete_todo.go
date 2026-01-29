package todo

import (
	"context"
	"fmt"
	"go-practice/domain/todo"
	"time"
)

// CompleteTodoInput はTodo完了のための入力データ
type CompleteTodoInput struct {
	ID string
}

// CompleteTodoUseCase はTodo完了のユースケース
type CompleteTodoUseCase struct {
	todoRepo todo.ITodoRepository
}

// NewCompleteTodoUseCase はCompleteTodoUseCaseのコンストラクタ
func NewCompleteTodoUseCase(todoRepo todo.ITodoRepository) *CompleteTodoUseCase {
	return &CompleteTodoUseCase{
		todoRepo: todoRepo,
	}
}

// Execute はTodo完了を実行する
func (uc *CompleteTodoUseCase) Execute(ctx context.Context, input CompleteTodoInput) error {
	// 1. IDのバリデーション
	todoID, err := todo.NewTodoIDFromString(input.ID)
	if err != nil {
		return fmt.Errorf("invalid todo id: %w", err)
	}

	// 2. リポジトリからTodoを取得
	t, err := uc.todoRepo.FindByID(ctx, todoID)
	if err != nil {
		return fmt.Errorf("failed to find todo: %w", err)
	}

	// 3. ドメインロジックを実行(完了処理)
	now := time.Now()
	if err := t.Complete(now); err != nil {
		return fmt.Errorf("failed to complete todo: %w", err)
	}

	// 4. 変更を永続化
	if err := uc.todoRepo.Save(ctx, t); err != nil {
		return fmt.Errorf("failed to save todo: %w", err)
	}

	return nil
}
