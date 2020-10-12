package task

import "time"

type Task struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status
}

type Status int

const (
	Completed Status = iota
	Pending
	Cancelled
	None
)

func (s Status) String() string {
	return [...]string{"none", "pending", "cancelled", "completed"}[s]
}
