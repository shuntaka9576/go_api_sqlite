package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/shuntaka9576/go_api_sqlite/entity"
)

type ListTask struct {
	Service   ListTasksService
	Validator *validator.Validate
}

type task struct {
	ID      entity.TaskID     `json:"id"`
	Title   string            `json:"title"`
	Status  entity.TaskStatus `json:"status"`
	Created string            `json:"created"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Printf("RemoteAddr: %s\n", r.RemoteAddr)

	fwdAddress := r.Header.Get("X-Forwarded-For")
	if fwdAddress != "" {
		ips := strings.Split(fwdAddress, ",")

		for i, ip := range ips {
			log.Printf("X-Forwarded-For[%d]: %s", i, ip)
		}
	}

	tasks, err := lt.Service.ListTasks(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:      t.ID,
			Title:   t.Title,
			Status:  t.Status,
			Created: t.Created,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
