package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	"gorm.io/gorm"
)

type RepositoryDelete interface {
	DeleteTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError)
}

type repositoryDelete struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repositoryDelete {
	return &repositoryDelete{db: db}
}

func (r *repositoryDelete) DeleteTaskRepository (input *models.ModelTask) (*entities.Task, models.DatabaseError) {
	var task entities.Task
	db := r.db.Model(&task)
	errorCode := make(chan models.DatabaseError, 1)

	task.ID = input.ID

	checkTaskId := db.Debug().First(&task)

	if checkTaskId.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_01",
		}
		return &task, <-errorCode
	}

	deleteTaskId := db.Debug().Delete(&task)

	if deleteTaskId.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_02",
		}
		return &task, <-errorCode
	}
	errorCode <- models.DatabaseError{}
	return &task, <-errorCode
}