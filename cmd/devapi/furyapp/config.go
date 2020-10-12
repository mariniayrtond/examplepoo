package furyapp

import (
	"users_example/internal/platform/environment"
	"users_example/internal/platform/localdb"
	"users_example/internal/supervisor/developer"
	"users_example/internal/supervisor/developer/task"
	"users_example/internal/supervisor/report"
)

type Dependencies struct {
	DeveloperRepository developer.Repository
	ReportRepository    report.Repository
	ReportNotifier      report.Notifier
	TaskPublisher       task.Publisher
}

func BuildDependencies(env environment.Environment) (*Dependencies, error) {
	switch env {
	case environment.Beta, environment.Production, environment.Development:
		localDb := localdb.NewLocalDB()

		// infra adapters
		reportNotifier := report.NewFakeNotifier()
		reportRepo := report.NewLocalRepo(localDb)
		devRepo := developer.NewLocalRepo(localDb)
		taskPublisher := task.NewFakePublisher()

		return &Dependencies{
			DeveloperRepository: devRepo,
			ReportRepository:    reportRepo,
			ReportNotifier:      reportNotifier,
			TaskPublisher:       taskPublisher,
		}, nil
	}

	return nil, nil
}
