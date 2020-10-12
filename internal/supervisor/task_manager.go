package supervisor

import (
	"users_example/internal/supervisor/developer/task"
)

type TaskManagement interface {
	Schedule(id string, task task.Task) Error
	Complete(id string, task task.Task) Error
}
