package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/snowchest67/todo-api/internal/model"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

func NewPostgresRepo(ctx context.Context, connString string) (*PostgresRepo, error) {
	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
    return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresRepo{db: db}, nil

}

func (r *PostgresRepo) Close() {
	r.db.Close()
}

func (r *PostgresRepo) CreateTodo(ctx context.Context, title string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, "INSERT INTO todos (title) VALUES ($1) RETURNING id", title).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return -1, fmt.Errorf("couldn't insert todo: %w", err)
		}
		return -1, fmt.Errorf("database error: %w", err)
	}
	return id, nil
}

func (r *PostgresRepo) GetAll(ctx context.Context) ([]model.Todo, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM todos")
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	defer rows.Close()

	var todos []model.Todo

	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title) 
		if err != nil {
			return nil, fmt.Errorf("database error: %w", err)
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
  }

	return todos, nil
}

func(r *PostgresRepo) GetByID(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo

	err := r.db.QueryRow(ctx, "SELECT id, title FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title)
	if err != nil {
		if err == pgx.ErrNoRows{
			return nil, fmt.Errorf("todo id=%d not found", id)
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return &todo, nil
}

func(r *PostgresRepo) DeleteByID(ctx context.Context, id int) (error) {
	result, err := r.db.Exec(ctx, "DELETE FROM todos WHERE id = $1",  id)

	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	if result.RowsAffected() == 0 {
    return fmt.Errorf("todo id=%d not found", id)
  }

	return nil
}