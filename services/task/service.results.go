package services

import (
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/models"
	repositories "github.com/habibiiiberahim/go-task/repositories/task"
)

type ServiceResults interface {
	ResultsTaskService() (*[]entities.Task, models.DatabaseError)
}
type serviceResults struct {
	repository repositories.RepositoryResults
}

func NewServiceResults(repository repositories.RepositoryResults) *serviceResults {
	return &serviceResults{
		repository: repository,
	}
}

func (s *serviceResults)ResultsTaskService()(*[]entities.Task, models.DatabaseError)  {
	res, err := s.repository.ResultsTaskRepository()
	return res, err
}