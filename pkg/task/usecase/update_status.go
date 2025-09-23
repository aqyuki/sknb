package usecase

import (
	"context"

	"github.com/aqyuki/sknb/internal/appbase"
	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type UpdateStatusInput struct {
	Code   int
	Status task.Status
}

type UpdateStatus struct {
	tx   appbase.Tx
	repo task.Repository
}

func NewUpdateStatus(tx appbase.Tx, repo task.Repository) *UpdateStatus {
	return &UpdateStatus{
		tx:   tx,
		repo: repo,
	}
}

func (u *UpdateStatus) Invoke(ctx context.Context, input *UpdateStatusInput) error {
	return u.tx.Do(ctx, func(ctx context.Context) error {
		task, err := u.repo.FindByCodeForUpdate(ctx, input.Code)
		if err != nil {
			return err
		}

		task.Status = input.Status

		return u.repo.Save(ctx, task)
	})
}
