package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	"gorm.io/gorm"
)

type RepositoryResult interface {
	ResultTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError)
}

type repositoryResult struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repositoryResult {
	return &repositoryResult{db: db}
}

func (r *repositoryResult) ResultTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError) {

	var task entities.Task	
	db := r.db.Model(&task)
	errorCode := make(chan models.DatabaseError, 1)

	task.ID = input.ID

	resultTask := db.Debug().Find(&task)

	if resultTask.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_01",
		}
		return &task, <-errorCode
	}
	errorCode <- models.DatabaseError{}
	return &task, <-errorCode
}