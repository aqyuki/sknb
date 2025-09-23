package usecase

import (
	"context"

	"github.com/aqyuki/sknb/internal/appbase"
	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type UpdateTitleInput struct {
	Code  int
	Title string
}

type UpdateTitle struct {
	tx   appbase.Tx
	repo task.Repository
}

func NewUpdateTitle(tx appbase.Tx, repo task.Repository) *UpdateTitle {
	return &UpdateTitle{
		tx:   tx,
		repo: repo,
	}
}

func (u *UpdateTitle) Invoke(ctx context.Context, input *UpdateTitleInput) error {
	return u.tx.Do(ctx, func(ctx context.Context) error {
		task, err := u.repo.FindByCodeForUpdate(ctx, input.Code)
		if err != nil {
			return err
		}

		task.Title = input.Title

		return u.repo.Save(ctx, task)
	})
}
