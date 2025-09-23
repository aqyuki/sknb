package task

import "context"

type Repository interface {
	CreateAndReturn(ctx context.Context, draft *DraftTask) (*Task, error)
}
