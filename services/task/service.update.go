package services

import (
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
)

type ServiceUpdate interface {
	UpdateTaskService (input *models.ModelTask)(*entities.Task, models.DatabaseError)
}
type serviceUpdate struct {
	repository repositories.RepositoryUpdate
}

func NewServiceUpdate(repository repositories.RepositoryUpdate) *serviceUpdate {
	return &serviceUpdate{
		repository: repository,
	}
}

func (s *serviceUpdate)UpdateTaskService(input *models.ModelTask)(*entities.Task, models.DatabaseError)  {
	var task models.ModelTask
	task.ID = input.ID
	task.Description = input.Description
	task.Assignn = input.Assignn
	task.Deadline = input.Deadline
	
	res, err := s.repository.UpdateTaskRepository(&task)
	return res, err
}