package furyapp

import (
	"github.com/gin-gonic/gin"
	"users_example/cmd/devapi/furyapp/handlers"
	"users_example/internal/supervisor"
)

func Build(dep *Dependencies) *gin.Engine {
	r := gin.Default()

	// use cases
	reportsUseCase := supervisor.NewLazyReporterUseCase(dep.ReportRepository, dep.ReportNotifier, dep.DeveloperRepository)
	tasksUseCase := supervisor.NewTaskManagementUseCase(dep.TaskPublisher, dep.DeveloperRepository)
	crudUseCase := supervisor.NewCRUDUseCase(dep.DeveloperRepository)

	// controller adapters
	tasksHandler := handlers.NewTasksManagementHandler(tasksUseCase)
	reportHandler := handlers.NewReportHandler(reportsUseCase)
	crudHandler := handlers.NewCRUDHandler(crudUseCase)

	// crud endpoints
	r.POST("/supervisor/developer", crudHandler.HandleCreate)
	r.GET("/supervisor/developer/:id", crudHandler.HandleRead)
	r.PUT("/supervisor/developer/:id", crudHandler.HandleUpdate)
	r.DELETE("/supervisor/developer/:id", crudHandler.HandleDelete)

	// tasks endpoints
	r.POST("/supervisor/developer/:id/tasks/schedule", tasksHandler.HandleScheduleTask)
	r.POST("/supervisor/developer/:id/tasks/complete", tasksHandler.HandleCompleteTask)

	// report endpoints
	r.POST("/supervisor/report/lazy/generate", reportHandler.HandleLazyReport)

	return r
}
