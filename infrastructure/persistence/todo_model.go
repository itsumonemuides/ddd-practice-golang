package persistence

import (
	"go-practice/domain/todo"
	"time"
)

type TodoModel struct {
	ID          string     `gorm:"primarykey;type:varchar(36)"`
	Title       string     `gorm:"not null;type:varchar(100)"`
	Completed   bool       `gorm:"not null;default:false"`
	CreatedAt   time.Time  `gorm:"not null"`
	CompletedAt *time.Time `gorm:"type:timestamp"`
}

func (TodoModel) TableName() string {
	return "todos"
}

func (m *TodoModel) ToDomain() (*todo.Todo, error) {
	id, err := todo.NewTodoIDFromString(m.ID)
	if err != nil {
		return nil, err
	}

	title, err := todo.NewTitle(m.Title)
	if err != nil {
		return nil, err
	}

	return todo.Reconstruct(id, title, m.Completed, m.CreatedAt, m.CompletedAt), nil
}

func FromDomain(t *todo.Todo) *TodoModel {
	return &TodoModel{
		ID:          t.ID().String(),
		Title:       t.Title().String(),
		Completed:   t.IsCompleted(),
		CreatedAt:   t.CreatedAt(),
		CompletedAt: t.CompletedAt(),
	}
}
