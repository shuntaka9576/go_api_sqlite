package handler

import (
	"context"

	"github.com/shuntaka9576/go_api_sqlite/entity"
)

type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}
