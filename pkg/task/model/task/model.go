package task

import "github.com/cockroachdb/errors"

var ErrTaskNotFound = errors.New("指定されたタスクが見つかりません")

type Task struct {
	*Header
	*Body
	Code int
}

type DraftTask struct {
	*Header
	*Body
}

type Header struct {
	ID     string
	UserID string
}

type Body struct {
	Title       string
	Description string
	Status      Status
}

func NewTask(header *Header, body *Body, code int) *Task {
	return &Task{
		Header: header,
		Body:   body,
		Code:   code,
	}
}

func NewDraftTask(header *Header, body *Body) *DraftTask {
	return &DraftTask{
		Header: header,
		Body:   body,
	}
}
