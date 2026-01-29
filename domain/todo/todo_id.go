package todo

import (
	"errors"

	"github.com/google/uuid"
)

type TodoID struct {
	value string
}

// NewTodoID creates a new TodoID
func NewTodoID() TodoID {
	return TodoID{value: uuid.New().String()}
}

// NewTodoIDFromString creates TodoID from existing string
func NewTodoIDFromString(id string) (TodoID, error) {
	if id == "" {
		return TodoID{}, errors.New("todo id cannot be empty")
	}
	// uuidのバリデーション
	if _, err := uuid.Parse(id); err != nil {
		return TodoID{}, errors.New("invalid todo id format")
	}
	return TodoID{value: id}, nil
}

// String returns the string representations
func (id TodoID) String() string {
	return id.value
}

// Equals checks equality with another TodoID
func (id TodoID) Equals(other TodoID) bool {
	return id.value == other.value
}
