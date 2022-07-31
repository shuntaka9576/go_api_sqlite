package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shuntaka9576/go_api_sqlite/clock"
	"github.com/shuntaka9576/go_api_sqlite/config"
	"github.com/shuntaka9576/go_api_sqlite/handler"
	"github.com/shuntaka9576/go_api_sqlite/service"
	"github.com/shuntaka9576/go_api_sqlite/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{Service: &service.AddTask{DB: db, Repo: &r}, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Service: &service.ListTask{DB: db, Repo: &r}, Validator: v}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux, cleanup, nil
}
