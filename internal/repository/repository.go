package repository

import (
	"context"

	"github.com/snowchest67/todo-api/internal/model"
)

type TodoRepository interface {
    CreateTodo(ctx context.Context, title string) (int, error)
		GetAll(ctx context.Context) ([]model.Todo, error)
		GetByID(ctx context.Context, id int) (*model.Todo, error)
		DeleteByID(ctx context.Context, id int) (error)
}
