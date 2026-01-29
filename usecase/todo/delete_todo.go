package todo

import (
	"context"
	"fmt"
	"go-practice/domain/todo"
)

// DeleteTodoInput はTodo削除のための入力データ
type DeleteTodoInput struct {
	ID string
}

// DeleteTodoUseCase はTodo削除のユースケース
type DeleteTodoUseCase struct {
	todoRepo todo.ITodoRepository
}

// NewDeleteTodoUseCase はDeleteTodoUseCaseのコンストラクタ
func NewDeleteTodoUseCase(todoRepo todo.ITodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		todoRepo: todoRepo,
	}
}

// Execute はTodo削除を実行する
func (uc *DeleteTodoUseCase) Execute(ctx context.Context, input DeleteTodoInput) error {
	// 1. IDのバリデーション
	todoID, err := todo.NewTodoIDFromString(input.ID)
	if err != nil {
		return fmt.Errorf("invalid todo id: %w", err)
	}

	// 2. 存在確認(オプション)
	_, err = uc.todoRepo.FindByID(ctx, todoID)
	if err != nil {
		return fmt.Errorf("todo not found: %w", err)
	}

	// 3. 削除実行
	if err := uc.todoRepo.Delete(ctx, todoID); err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	return nil
}
