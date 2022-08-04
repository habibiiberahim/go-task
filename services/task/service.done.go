package services

import (
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
)

type ServiceDone interface {
	DoneTaskService (input *models.ModelTask)(*entities.Task, models.DatabaseError)
}
type serviceDone struct {
	repository repositories.RepositoryDone
}

func NewServiceDone(repository repositories.RepositoryDone) *serviceDone {
	return &serviceDone{
		repository: repository,
	}
}

func (s *serviceDone)DoneTaskService(input *models.ModelTask)(*entities.Task, models.DatabaseError)  {
	var task models.ModelTask
	task.ID = input.ID
	
	res, err := s.repository.DoneTaskRepository(&task)
	return res, err
}