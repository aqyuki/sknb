package task

import (
	"context"

	"github.com/aqyuki/sknb/internal/appbase"
)

type Query struct {
	appbase.PaginateQuery

	// 指定されたユーザーのタスクを取得する
	UserIDs []string

	// 指定されたステータスのタスクを取得する
	Status []Status
}

//go:generate go tool mockgen -source ./repository.go -destination ./repository_mock.go -package task -typed
type Repository interface {
	Find(ctx context.Context, query *Query) ([]*Task, error)
	FindByCode(ctx context.Context, code int) (*Task, error)
	FindByCodeForUpdate(ctx context.Context, code int) (*Task, error)
	CreateAndReturn(ctx context.Context, draft *DraftTask) (*Task, error)
	Save(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
}
