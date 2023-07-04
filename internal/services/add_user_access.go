package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type AddAccessService struct {
	repo repositories.IAddAccessRepo
}

func NewAddAccessService(repo repositories.IAddAccessRepo) *AddAccessService {
	return &AddAccessService{
		repo: repo,
	}
}

func (s *AddAccessService) AddAccessUser(request models.AddUserAccessRequest, response chan models.AddUserAccessResponse) {
	result := s.repo.AddAccessUser(request)

	response <- models.AddUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
