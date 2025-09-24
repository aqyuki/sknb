package task

import "context"

//go:generate go tool mockgen -source ./repository.go -destination ./repository_mock.go -package task -typed
type Repository interface {
	FindByCode(ctx context.Context, code int) (*Task, error)
	FindByCodeForUpdate(ctx context.Context, code int) (*Task, error)
	CreateAndReturn(ctx context.Context, draft *DraftTask) (*Task, error)
	Save(ctx context.Context, task *Task) error
}
