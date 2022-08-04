package repositories

import (
	"net/http"

	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	"gorm.io/gorm"
)

type RepositoryDone interface {
	DoneTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError)
}

type repositoryDone struct {
	db *gorm.DB
}

func NewRepositoryDone(db *gorm.DB) *repositoryDone {
	return &repositoryDone{db: db}
}

func (r *repositoryDone) DoneTaskRepository(input *models.ModelTask) (*entities.Task, models.DatabaseError) {

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

	task.IsDone = true

	doneTask := db.Debug().Where("id = ?", input.ID).Updates(&task)

	if doneTask.RowsAffected < 1 {
		errorCode <- models.DatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &task, <-errorCode
	}
	errorCode <- models.DatabaseError{}
	return &task, <-errorCode
}