package usecase

import (
	"bytes"
	"context"
	"encoding/csv"
	"io"
	"strconv"

	"github.com/aqyuki/sknb/pkg/task/model/task"
	"github.com/cockroachdb/errors"
)

type ExportTaskInput struct {
	Query *task.Query
}

type ExportTask struct {
	listTask *ListTask
}

func NewExportTask(listTask *ListTask) *ExportTask {
	return &ExportTask{
		listTask: listTask,
	}
}

func (u *ExportTask) Invoke(ctx context.Context, input *ExportTaskInput) (io.Reader, error) {
	query := &ListTaskInput{
		PaginateQuery: input.Query.PaginateQuery,
		UserIDs:       input.Query.UserIDs,
		Status:        input.Query.Status,
	}

	tasks, err := u.listTask.Invoke(ctx, query)
	if err != nil {
		return nil, err
	}

	return toCSV(tasks)
}

func toCSV(tasks []*task.Task) (io.Reader, error) {
	var buffer bytes.Buffer

	writer := csv.NewWriter(&buffer)
	defer writer.Flush()

	for _, item := range tasks {
		row := []string{
			item.ID,
			strconv.Itoa(item.Code),
			item.UserID,
			string(item.Status),
			item.Title,
			item.Description,
		}
		if err := writer.Write(row); err != nil {
			return nil, errors.Wrap(err, "CSVへの変換に失敗しました")
		}
	}

	return &buffer, nil
}
