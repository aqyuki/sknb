package appbase

import "context"

type Tx interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type MockTx struct{}

func (tx *MockTx) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}
