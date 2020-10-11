package internal

import "time"

type Task struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    TaskStatus
}

type TaskStatus int

const (
	Completed TaskStatus = iota
	Pending
	Cancelled
	None
)

func (s TaskStatus) String() string {
	return [...]string{"completed", "pending", "cancelled", "none"}[s]
}
