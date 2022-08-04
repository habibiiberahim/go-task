package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	"gorm.io/gorm"
)

type RepositoryCreate interface {
	CreateTaskRepository (input *models.ModelTask)(*entities.Task, models.DatabaseError)
}

type repositoryCreate struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repositoryCreate {
	return &repositoryCreate{
		db: db,
	}
}

func (r *repositoryCreate)CreateTaskRepository(input *models.ModelTask)(*entities.Task, models.DatabaseError)  {
	var task entities.Task
	db := r.db.Model(&task)
	errorCode := make(chan models.DatabaseError, 1)
	
	task.Description = input.Description
	task.Assignn = input.Assignn
	task.IsDone = false
	task.Deadline = input.Deadline


	addNewTask := db.Debug().Create(&task).Commit() 

	if addNewTask.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_forbidden",
		}
		return &task, <-errorCode
	}

	errorCode <- models.DatabaseError{}
	return &task, <-errorCode
}