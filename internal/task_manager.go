package internal

type TaskManager interface {
	Schedule(id string, task Task) DevApiError
	Complete(id string, task Task) DevApiError
}
