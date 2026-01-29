package todo

import (
	"context"
	"fmt"
	"go-practice/domain/todo"
)

type CreateTodoInput struct {
	Title string
}

type CreateTodoOutput struct {
	ID    string
	Title string
}

type CreateTodoUseCase struct {
	todoRepo todo.ITodoRepository
}

func NewCreateTodoUseCase(todoRepo todo.ITodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		todoRepo: todoRepo,
	}
}

func (uc *CreateTodoUseCase) Execute(ctx context.Context, input CreateTodoInput) (*CreateTodoOutput, error) {
	title, err := todo.NewTitle(input.Title)
	if err != nil {
		return nil, fmt.Errorf("invalid title: %w", err)
	}

	newTodo := todo.NewTodo(title)

	if err := uc.todoRepo.Save(ctx, newTodo); err != nil {
		return nil, fmt.Errorf("failed to save todo: %w", err)
	}

	return &CreateTodoOutput{
		ID:    newTodo.ID().String(),
		Title: newTodo.Title().String(),
	}, nil
}
