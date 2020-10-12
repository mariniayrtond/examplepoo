package internal

import "time"

type Report struct {
	Id             string
	CreatedAt      time.Time
	LazyDevelopers int
	Details        []Detail
}

type Detail struct {
	DeveloperName  string
	Seniority      Seniority
	TaskStatus     TaskStatus
	LastActivityAt time.Time
}
