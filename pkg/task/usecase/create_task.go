package usecase

import (
	"context"

	"github.com/aqyuki/sknb/pkg/task/model/task"
	"github.com/rs/xid"
)

const defaultStatus = task.StatusToDo

type CreateTaskInput struct {
	UserID      string
	Title       string
	Description string
}

type CreateTask struct {
	repo task.Repository
}

func NewCreateTask(repo task.Repository) *CreateTask {
	return &CreateTask{repo: repo}
}

func (u *CreateTask) Invoke(ctx context.Context, input *CreateTaskInput) (*task.Task, error) {
	header := &task.Header{
		ID:     xid.New().String(),
		UserID: input.UserID,
	}
	body := &task.Body{
		Title:       input.Title,
		Description: input.Description,
		Status:      defaultStatus,
	}
	draft := task.NewDraftTask(header, body)

	return u.repo.CreateAndReturn(ctx, draft)
}
