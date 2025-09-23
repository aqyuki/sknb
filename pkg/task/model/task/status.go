package task

type Status string

const (
	StatusToDo           = Status("todo")
	StatusWorkInProgress = Status("wip")
	StatusComplete       = Status("complete")
)
