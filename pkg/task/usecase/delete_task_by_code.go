package usecase

import (
	"context"

	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type DeleteTaskByCode struct {
	repo task.Repository
}

func NewDeleteTask(repo task.Repository) *DeleteTaskByCode {
	return &DeleteTaskByCode{repo: repo}
}

func (u *DeleteTaskByCode) Invoke(ctx context.Context, id string) error {
	return u.repo.DeleteByCode(ctx, id)
}
