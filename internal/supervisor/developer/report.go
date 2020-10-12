package developer

import (
	"time"
	"users_example/internal/supervisor/developer/task"
)

type Report struct {
	Id             string
	CreatedAt      time.Time
	LazyDevelopers int
	Details        []ReportDetail
}

type ReportDetail struct {
	DeveloperName  string
	Seniority      Seniority
	TaskStatus     task.Status
	LastActivityAt time.Time
}
