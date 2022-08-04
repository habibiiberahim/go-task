package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/habibiiiberahim/go-task/handlers/task"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
	services "github.com/habibiiiberahim/go-task/services/task"
	"gorm.io/gorm"
)

func InitTaskRoutes(db *gorm.DB, route *fiber.App)  {
	createTaskRepository := repositories.NewRepositoryCreate(db)
	createTaskService := services.NewServiceCreate(createTaskRepository)
	createTaskHandler := handlers.NewHandlerCreateTask(createTaskService)

	resultsTaskRepository := repositories.NewRepositoryResults(db)
	resultsTaskService := services.NewServiceResults(resultsTaskRepository)
	resultsTaskHandler := handlers.NewHandlerResultsTask(resultsTaskService)

	resultTaskRepository := repositories.NewRepositoryResults(db)
	resultTaskService := services.NewServiceResults(resultTaskRepository)
	resultTaskHandler := handlers.NewHandlerResultsTask(resultTaskService)

	deleteTaskRepository := repositories.NewRepositoryDelete(db)
	deleteTaskService := services.NewServiceDelete(deleteTaskRepository)
	deleteTaskHandler := handlers.NewHandlerDeleteTask(deleteTaskService)

	updateTaskRepository := repositories.NewRepositoryUpdate(db)
	updateTaskService := services.NewServiceUpdate(updateTaskRepository)
	updateTaskHandler := handlers.NewHandlerUpdateTask(updateTaskService)

	doneTaskRepository := repositories.NewRepositoryDone(db)
	doneTaskService := services.NewServiceDone(doneTaskRepository)
	doneTaskHandler := handlers.NewHandlerDoneTask(doneTaskService)
	
	groupRoute := route.Group("/api/v1")	
	groupRoute.Post("/task", createTaskHandler.CreateTaskHandler)
	groupRoute.Get("/task", resultsTaskHandler.ResultsTaskHandler)
	groupRoute.Get("/task/:id", resultTaskHandler.ResultsTaskHandler)
	groupRoute.Get("/task/:id/done", doneTaskHandler.DoneTaskHandler)
	groupRoute.Delete("/task/:id", deleteTaskHandler.DeleteTaskHandler)
	groupRoute.Put("/task/:id", updateTaskHandler.UpdateTaskHandler)
}