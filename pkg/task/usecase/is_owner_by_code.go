package usecase

import (
	"context"

	"github.com/aqyuki/sknb/pkg/task/model/task"
)

type IsOwnerByCode struct {
	repo task.Repository
}

func NewIsOwnerByCode(repo task.Repository) *IsOwnerByCode {
	return &IsOwnerByCode{
		repo: repo,
	}
}

func (u *IsOwnerByCode) Invoke(ctx context.Context, userID string, code int) error {
	t, err := u.repo.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	if t.UserID != userID {
		return task.ErrNoTaskOwner
	}

	return nil
}
