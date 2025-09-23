package usecase

import (
	"context"
	"testing"

	"github.com/aqyuki/sknb/pkg/task/model/task"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateTask_Invoke(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := task.NewMockRepository(ctrl)
		repo.
			EXPECT().
			CreateAndReturn(gomock.Eq(t.Context()), gomock.Any()).
			DoAndReturn(func(_ context.Context, draft *task.DraftTask) (*task.Task, error) {
				return &task.Task{
					Header: draft.Header,
					Body:   draft.Body,
					Code:   0,
				}, nil
			})

		input := &CreateTaskInput{
			UserID:      "user_id",
			Title:       "title",
			Description: "description",
		}
		useCase := NewCreateTask(repo)

		actual, err := useCase.Invoke(t.Context(), input)
		require.NoError(t, err)
		require.NotEmpty(t, actual.ID)
		require.Equal(t, "user_id", actual.UserID)
		require.Equal(t, "title", actual.Title)
		require.Equal(t, "description", actual.Description)
		require.Equal(t, 0, actual.Code)
		require.Equal(t, task.StatusToDo, actual.Status)
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := task.NewMockRepository(ctrl)
		repo.
			EXPECT().
			CreateAndReturn(gomock.Eq(t.Context()), gomock.Any()).
			Return(nil, errors.New("some error"))

		input := &CreateTaskInput{
			UserID:      "user_id",
			Title:       "title",
			Description: "description",
		}
		useCase := NewCreateTask(repo)

		actual, err := useCase.Invoke(t.Context(), input)
		require.Error(t, err)
		require.Nil(t, actual)
	})
}
