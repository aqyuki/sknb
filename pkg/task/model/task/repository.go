package task

import "context"

//go:generate go tool mockgen -source ./repository.go -destination ./repository_mock.go -package task -typed
type Repository interface {
	CreateAndReturn(ctx context.Context, draft *DraftTask) (*Task, error)
}
