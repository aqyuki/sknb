package usecase

import (
	"testing"

	"github.com/aqyuki/sknb/pkg/task/model/task"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestIsOwnerByCode_Invoke(t *testing.T) {
	t.Parallel()

	t.Run("正常系", func(t *testing.T) {
		t.Parallel()

		t.Run("オーナーの場合", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := task.NewMockRepository(ctrl)
			repo.
				EXPECT().
				FindByCode(gomock.Eq(t.Context()), gomock.Eq(0)).
				Return(&task.Task{
					Header: &task.Header{ID: "id", UserID: "userID"},
					Body:   &task.Body{Title: "title", Description: "description", Status: task.StatusToDo},
					Code:   0,
				}, nil)

			useCase := NewIsOwnerByCode(repo)

			err := useCase.Invoke(t.Context(), "userID", 0)
			require.NoError(t, err)
		})

		t.Run("オーナーでない場合", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := task.NewMockRepository(ctrl)
			repo.
				EXPECT().
				FindByCode(gomock.Eq(t.Context()), gomock.Eq(0)).
				Return(&task.Task{
					Header: &task.Header{ID: "id", UserID: "userID"},
					Body:   &task.Body{Title: "title", Description: "description", Status: task.StatusToDo},
					Code:   0,
				}, nil)

			useCase := NewIsOwnerByCode(repo)

			err := useCase.Invoke(t.Context(), "_userID", 0)
			require.Error(t, err)
			require.ErrorIs(t, err, task.ErrNoTaskOwner)
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := task.NewMockRepository(ctrl)
		repo.
			EXPECT().
			FindByCode(gomock.Eq(t.Context()), gomock.Eq(0)).
			Return(nil, errors.New("some error"))

		useCase := NewIsOwnerByCode(repo)

		err := useCase.Invoke(t.Context(), "userID", 0)
		require.Error(t, err)
	})
}
