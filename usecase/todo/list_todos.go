package todo

import (
	"context"
	"fmt"
	"go-practice/domain/todo"
	"time"
)

type TodoDTO struct {
	ID          string
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type ListTodosUseCase struct {
	todoRepo todo.ITodoRepository
}

func NewListTodoUseCase(todoRepo todo.ITodoRepository) *ListTodosUseCase {
	return &ListTodosUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *ListTodosUseCase) Execute(ctx context.Context) ([]TodoDTO, error) {
	todos, err := uc.todoRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}
	dtos := make([]TodoDTO, 0, len(todos))
	for _, t := range todos {
		dtos = append(dtos, TodoDTO{
			ID:          t.ID().String(),
			Title:       t.Title().String(),
			Completed:   t.IsCompleted(),
			CreatedAt:   t.CreatedAt(),
			CompletedAt: t.CompletedAt(),
		})
	}

	return dtos, nil
}
