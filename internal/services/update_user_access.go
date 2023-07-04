package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type UpdateAccessService struct {
	repo repositories.IUpdateAccessRepo
}

func NewUpdateAccessService(repo repositories.IUpdateAccessRepo) *UpdateAccessService {
	return &UpdateAccessService{
		repo: repo,
	}
}

func (s *UpdateAccessService) UpdateAccessUser(request models.UpdateUserAccessRequest, response chan models.UpdateUserAccessResponse) {
	result := s.repo.UpdateAccessUser(request)

	response <- models.UpdateUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
