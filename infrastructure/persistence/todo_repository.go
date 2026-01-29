package persistence

import (
	"context"
	"errors"
	"fmt"
	"go-practice/domain/todo"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Save(ctx context.Context, t *todo.Todo) error {
	model := FromDomain(t)

	// Upsert (INSERT or UPDATE)
	result := r.db.WithContext(ctx).Save(model)
	if result.Error != nil {
		return fmt.Errorf("failed to save todo: %w", result.Error)
	}

	return nil
}

// FindByID はIDでTodoを取得
func (r *TodoRepository) FindByID(ctx context.Context, id todo.TodoID) (*todo.Todo, error) {
	var model TodoModel

	result := r.db.WithContext(ctx).Where("id = ?", id.String()).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("todo not found: %s", id.String())
		}
		return nil, fmt.Errorf("failed to find todo: %w", result.Error)
	}

	return model.ToDomain()
}

// FindAll は全てのTodoを取得
func (r *TodoRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	var models []TodoModel

	result := r.db.WithContext(ctx).Order("created_at DESC").Find(&models)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find all todos: %w", result.Error)
	}

	todos := make([]*todo.Todo, 0, len(models))
	for _, model := range models {
		t, err := model.ToDomain()
		if err != nil {
			return nil, fmt.Errorf("failed to convert to domain: %w", err)
		}
		todos = append(todos, t)
	}

	return todos, nil
}

// Delete はTodoを削除
func (r *TodoRepository) Delete(ctx context.Context, id todo.TodoID) error {
	result := r.db.WithContext(ctx).Where("id = ?", id.String()).Delete(&TodoModel{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete todo: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("todo not found: %s", id.String())
	}

	return nil
}
