package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shuntaka9576/go_api_sqlite/entity"
)

type AddTask struct {
	DB   *sqlx.DB
	Repo TaskAdder
}

func (at *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := at.Repo.AddTask(ctx, at.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return t, nil
}
