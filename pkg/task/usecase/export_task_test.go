package usecase

import (
	_ "embed"
	"io"
	"testing"

	"github.com/aqyuki/sknb/internal/appbase"
	"github.com/aqyuki/sknb/pkg/task/model/task"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

//go:embed testdata/export.csv
var embedCSV []byte

func TestExportTask_Invoke(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := task.NewMockRepository(ctrl)
	repo.
		EXPECT().
		Find(
			gomock.Eq(t.Context()),
			gomock.Eq(
				&task.Query{
					PaginateQuery: appbase.PaginateQuery{Offset: 0, Limit: 10},
					UserIDs:       []string{},
					Status:        []task.Status{},
				})).
		Return([]*task.Task{
			{
				Header: &task.Header{ID: "foo", UserID: "user1"},
				Body:   &task.Body{Title: "title1", Description: "description1", Status: task.StatusWorkInProgress},
				Code:   0,
			},
			{
				Header: &task.Header{ID: "bar", UserID: "user2"},
				Body:   &task.Body{Title: "title2", Description: "description2", Status: task.StatusToDo},
				Code:   1,
			},
			{
				Header: &task.Header{ID: "baz", UserID: "user3"},
				Body:   &task.Body{Title: "title3", Description: "description3", Status: task.StatusComplete},
				Code:   2,
			},
		}, nil)

	listTask := NewListTask(repo)
	useCase := NewExportTask(listTask)

	reader, err := useCase.Invoke(t.Context(), &ExportTaskInput{
		Query: &task.Query{
			PaginateQuery: appbase.PaginateQuery{Offset: 0, Limit: 10},
			UserIDs:       []string{},
			Status:        []task.Status{},
		},
	})
	require.NoError(t, err)

	actual, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.Equal(t, string(embedCSV), string(actual))
}
