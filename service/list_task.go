package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/shuntaka9576/go_api_sqlite/entity"
	"github.com/shuntaka9576/go_api_sqlite/store"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo *store.Repository
}

func (lt *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}

	return tasks, nil
}
