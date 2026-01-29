package todo

import "context"

type ITodoRepository interface {
	Save(ctx context.Context, todo *Todo) error
	FindByID(ctx context.Context, id TodoID) (*Todo, error)
	FindAll(ctx context.Context) ([]*Todo, error)
	Delete(ctx context.Context, id TodoID) error
}
