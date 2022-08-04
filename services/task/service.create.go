package services

import (
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
)

type ServiceCreate interface {
	CreateTaskService (input *models.ModelTask)(*entities.Task, models.DatabaseError)
}
type serviceCreate struct {
	repository repositories.RepositoryCreate
}

func NewServiceCreate(repository repositories.RepositoryCreate) *serviceCreate {
	return &serviceCreate{
		repository: repository,
	}
}

func (s *serviceCreate)CreateTaskService(input *models.ModelTask)(*entities.Task, models.DatabaseError)  {
	var task models.ModelTask
	task.Description = input.Description
	task.Assignn = input.Assignn
	task.Deadline = input.Deadline
	
	res, err := s.repository.CreateTaskRepository(&task)
	return res, err
}