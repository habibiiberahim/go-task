package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	"gorm.io/gorm"
)

type RepositoryResults interface {
	ResultsTaskRepository() (*[]entities.Task, models.DatabaseError)
}

type repositoryResults struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repositoryResults {
	return &repositoryResults{db: db}
}

func (r *repositoryResults) ResultsTaskRepository() (*[]entities.Task, models.DatabaseError) {

	var task []entities.Task	
	db := r.db.Model(&task)
	errorCode := make(chan models.DatabaseError, 1)

	resultsTask := db.Debug().Find(&task)

	if resultsTask.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_01",
		}
		return &task, <-errorCode
	}
	errorCode <- models.DatabaseError{}
	return &task, <-errorCode
}