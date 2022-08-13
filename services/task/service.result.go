package services

import (
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
)

type ServiceResult interface {
	ResultTaskService(input *models.ModelTask) (*entities.Task, models.DatabaseError)
}
type serviceResult struct {
	repository repositories.RepositoryResult
}

func NewServiceResult(repository repositories.RepositoryResult) *serviceResult {
	return &serviceResult{
		repository: repository,
	}
}

func (s *serviceResult)ResultTaskService(input *models.ModelTask)(*entities.Task, models.DatabaseError)  {
	var task models.ModelTask
	task.ID = input.ID
	res, err := s.repository.ResultTaskRepository(&task)
	return res, err
}