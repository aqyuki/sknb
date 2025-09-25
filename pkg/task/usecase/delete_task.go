package usecase

import (
	"context"

	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type DeleteTask struct {
	repo task.Repository
}

func NewDeleteTask(repo task.Repository) *DeleteTask {
	return &DeleteTask{repo: repo}
}

func (u *DeleteTask) Invoke(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
