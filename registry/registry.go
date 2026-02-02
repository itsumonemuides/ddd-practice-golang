package registry

import (
	"go-practice/infrastructure/persistence"
	"go-practice/interface/handler"
	"go-practice/usecase/todo"
	"log"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) *Registry {
	return &Registry{
		db: db,
	}
}

func (r *Registry) NewTodoHandler() *handler.TodoHandler {
	todoRepo := persistence.NewTodoRepository(r.db)

	createUseCase := todo.NewCreateTodoUseCase(todoRepo)
	listUseCase := todo.NewListTodoUseCase(todoRepo)
	completeUseCase := todo.NewCompleteTodoUseCase(todoRepo)
	deleteUseCase := todo.NewDeleteTodoUseCase(todoRepo)
	uncompleteUseCase := todo.NewUncompleteTodoUseCase(todoRepo)

	return handler.NewTodoHandler(createUseCase, listUseCase, completeUseCase, deleteUseCase, uncompleteUseCase)
}

func (r *Registry) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	log.Println("Closing database connection...")

	return sqlDB.Close()
}
