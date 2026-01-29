package persistence

import (
	"context"
	"errors"
	"go-practice/domain/todo"
	"sync"
)

type InMemoryTodoRepository struct {
	mu    sync.RWMutex
	todos map[string]*todo.Todo
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]*todo.Todo),
	}
}

// Save はTodoを保存
func (r *InMemoryTodoRepository) Save(ctx context.Context, t *todo.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.todos[t.ID().String()] = t
	return nil
}

// FindByID はIDでTodoを取得
func (r *InMemoryTodoRepository) FindByID(ctx context.Context, id todo.TodoID) (*todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, exists := r.todos[id.String()]
	if !exists {
		return nil, errors.New("todo not found")
	}

	return t, nil
}

// FindAll は全てのTodoを取得
func (r *InMemoryTodoRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos := make([]*todo.Todo, 0, len(r.todos))
	for _, t := range r.todos {
		todos = append(todos, t)
	}

	return todos, nil
}

// Delete はTodoを削除
func (r *InMemoryTodoRepository) Delete(ctx context.Context, id todo.TodoID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id.String()]; !exists {
		return errors.New("todo not found")
	}

	delete(r.todos, id.String())
	return nil
}
