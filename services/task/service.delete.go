package services

import (
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
)

type ServiceDelete interface {
	DeleteTaskService (input *models.ModelTask)(*entities.Task, models.DatabaseError)
}
type serviceDelete struct {
	repository repositories.RepositoryDelete
}

func NewServiceDelete(repository repositories.RepositoryDelete) *serviceDelete {
	return &serviceDelete{
		repository: repository,
	}
}

func (s *serviceDelete)DeleteTaskService(input *models.ModelTask)(*entities.Task, models.DatabaseError)  {
	var task models.ModelTask
	task.ID = input.ID

	
	res, err := s.repository.DeleteTaskRepository(&task)
	return res, err
}