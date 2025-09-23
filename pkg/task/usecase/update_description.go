package usecase

import (
	"context"

	"github.com/aqyuki/sknb/internal/appbase"
	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type UpdateDescriptionInput struct {
	Code        int
	Description string
}

type UpdateDescription struct {
	tx   appbase.Tx
	repo task.Repository
}

func NewUpdateDescription(tx appbase.Tx, repo task.Repository) *UpdateDescription {
	return &UpdateDescription{
		tx:   tx,
		repo: repo,
	}
}

func (u *UpdateDescription) Invoke(ctx context.Context, input *UpdateDescriptionInput) error {
	return u.tx.Do(ctx, func(ctx context.Context) error {
		task, err := u.repo.FindByCodeForUpdate(ctx, input.Code)
		if err != nil {
			return err
		}

		task.Description = input.Description

		return u.repo.Save(ctx, task)
	})
}
