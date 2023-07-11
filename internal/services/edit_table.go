package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type EditTableService struct {
	repo repositories.IEditTableRepo
}

func NewEditTableService(repo repositories.IEditTableRepo) *EditTableService {
	return &EditTableService{
		repo: repo,
	}
}

func (s *EditTableService) EditTable(request models.EditTableRequest, response chan models.EditTableResponse) {
	result := s.repo.EditTable(request)

	response <- models.EditTableResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
