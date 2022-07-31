package service

import (
	"context"

	"github.com/shuntaka9576/go_api_sqlite/entity"
	"github.com/shuntaka9576/go_api_sqlite/store"
)

type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
