package repositories

import (
	"net/http"

	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	"gorm.io/gorm"
)

type RepositoryUpdate interface {
	UpdateTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError)
}

type repositoryUpdate struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repositoryUpdate {
	return &repositoryUpdate{db: db}
}

func (r *repositoryUpdate) UpdateTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError) {

	var task entities.Task
	db := r.db.Model(&task)
	errorCode := make(chan models.DatabaseError, 1)

	task.ID = input.ID

	checkTaskId := db.Debug().First(&task)

	if checkTaskId.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &task, <-errorCode
	}

	task.Description = input.Description
	task.Assignn = input.Assignn
	task.IsDone = input.IsDone
	task.Deadline = input.Deadline

	updateTask := db.Debug().Where("id = ?", input.ID).Updates(&task)

	if updateTask.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &task, <-errorCode
	}
	errorCode <- models.DatabaseError{}
	return &task, <-errorCode
}