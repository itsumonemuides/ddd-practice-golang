package todo

import (
	"errors"
	"time"
)

type Todo struct {
	id          TodoID
	title       Title
	completed   bool
	createdAt   time.Time
	completedAt *time.Time // ポインタ型にすることでnilを扱える
}

func NewTodo(title Title, createdAt time.Time) *Todo {
	return &Todo{
		id:        NewTodoID(),
		title:     title,
		completed: false,
		createdAt: createdAt,
	}
}

func Reconstruct(id TodoID, title Title, completed bool, createdAt time.Time, completedAt *time.Time) *Todo {
	return &Todo{
		id:          id,
		title:       title,
		completed:   completed,
		createdAt:   createdAt,
		completedAt: completedAt,
	}
}

func (t *Todo) Complete(now time.Time) error {
	if t.completed {
		return errors.New("todo is already completed")
	}
	t.completed = true
	t.completedAt = &now
	return nil
}

func (t *Todo) ChangeTitle(newTitle Title) {
	t.title = newTitle
}

// Getters
func (t *Todo) ID() TodoID              { return t.id }
func (t *Todo) Title() Title            { return t.title }
func (t *Todo) IsCompleted() bool       { return t.completed }
func (t *Todo) CreatedAt() time.Time    { return t.createdAt }
func (t *Todo) CompletedAt() *time.Time { return t.completedAt }
