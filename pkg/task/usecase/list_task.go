package usecase

import (
	"context"

	"github.com/aqyuki/sknb/internal/appbase"
	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type ListTaskInput struct {
	appbase.PaginateQuery

	UserIDs []string
	Status  []task.Status
}

type ListTask struct {
	repo task.Repository
}

func NewListTask(repo task.Repository) *ListTask {
	return &ListTask{
		repo: repo,
	}
}

func (u *ListTask) Invoke(ctx context.Context, input *ListTaskInput) ([]*task.Task, error) {
	query := &task.Query{
		PaginateQuery: input.PaginateQuery,

		UserIDs: input.UserIDs,
		Status:  input.Status,
	}

	return u.repo.Find(ctx, query)
}
