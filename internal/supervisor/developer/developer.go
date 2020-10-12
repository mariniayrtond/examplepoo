package developer

import (
	"users_example/internal/supervisor/developer/task"
)

type Repository interface {
	Add(d *Developer) error
	Get(id string) (*Developer, error)
	Update(id string, d *Developer) error
	Delete(id string) error
}

type Developer struct {
	ID        string
	Name      string
	Team      string
	Seniority Seniority
	Task      task.Task
}

func (d Developer) IsBusy() bool {
	return d.Task.Status == task.Pending
}

type Seniority int

const (
	Senior Seniority = iota
	SemiSenior
	Analyst
	Junior
)

func (s Seniority) String() string {
	return [...]string{"senior", "semi_senior", "analyst", "junior"}[s]
}
