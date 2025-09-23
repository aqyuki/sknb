package usecase

import (
	"testing"

	"github.com/aqyuki/sknb/internal/appbase"
	"github.com/aqyuki/sknb/pkg/task/model/task"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUpdateDescription_Invoke(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := task.NewMockRepository(ctrl)
		repo.
			EXPECT().
			FindByCodeForUpdate(gomock.Eq(t.Context()), gomock.Eq(0)).
			Return(&task.Task{
				Header: &task.Header{ID: "id", UserID: "userID"},
				Body:   &task.Body{Title: "title", Description: "description", Status: task.StatusToDo},
				Code:   0,
			}, nil)
		repo.
			EXPECT().
			Save(gomock.Eq(t.Context()), &task.Task{
				Header: &task.Header{ID: "id", UserID: "userID"},
				Body:   &task.Body{Title: "title", Description: "new description", Status: task.StatusToDo},
				Code:   0,
			}).
			Return(nil)

		input := &UpdateDescriptionInput{
			Code:        0,
			Description: "new description",
		}
		useCase := NewUpdateDescription(&appbase.MockTx{}, repo)

		err := useCase.Invoke(t.Context(), input)
		require.NoError(t, err)
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		t.Run("指定されたコードのタスクが存在しない", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := task.NewMockRepository(ctrl)
			repo.
				EXPECT().
				FindByCodeForUpdate(gomock.Eq(t.Context()), gomock.Eq(0)).
				Return(nil, task.ErrTaskNotFound)

			input := &UpdateDescriptionInput{
				Code:        0,
				Description: "new description",
			}
			useCase := NewUpdateDescription(&appbase.MockTx{}, repo)

			err := useCase.Invoke(t.Context(), input)
			require.Error(t, err)
		})

		t.Run("更新に失敗", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := task.NewMockRepository(ctrl)
			repo.
				EXPECT().
				FindByCodeForUpdate(gomock.Eq(t.Context()), gomock.Eq(0)).
				Return(&task.Task{
					Header: &task.Header{ID: "id", UserID: "userID"},
					Body:   &task.Body{Title: "title", Description: "description", Status: task.StatusToDo},
					Code:   0,
				}, nil)
			repo.
				EXPECT().
				Save(gomock.Eq(t.Context()), &task.Task{
					Header: &task.Header{ID: "id", UserID: "userID"},
					Body:   &task.Body{Title: "title", Description: "new description", Status: task.StatusToDo},
					Code:   0,
				}).
				Return(errors.New("some error"))

			input := &UpdateDescriptionInput{
				Code:        0,
				Description: "new description",
			}
			useCase := NewUpdateDescription(&appbase.MockTx{}, repo)

			err := useCase.Invoke(t.Context(), input)
			require.Error(t, err)
		})
	})
}
