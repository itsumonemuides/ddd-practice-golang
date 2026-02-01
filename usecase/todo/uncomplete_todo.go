package todo

import (
	"context"
	"fmt"
	"go-practice/domain/todo"
)

type UncompleteInput struct {
	ID string
}

type UncompleteTodoUseCase struct {
	todoRepo todo.ITodoRepository
}

func NewUncompleteTodoUseCase(todoRepo todo.ITodoRepository) *UncompleteTodoUseCase {
	return &UncompleteTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *UncompleteTodoUseCase) Execute(ctx context.Context, input UncompleteInput) error {
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

	// 3. ドメインロジックを実行(未完了にする)
	if err := t.UnComplete(); err != nil {
		return fmt.Errorf("failed to complete todo: %w", err)
	}

	// 4. 変更を永続化
	if err := uc.todoRepo.Save(ctx, t); err != nil {
		return fmt.Errorf("failed to save todo: %w", err)
	}

	return nil
}
