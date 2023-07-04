package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type DeleteAccessService struct {
	repo repositories.IDeleteAccessRepo
}

func NewDeleteAccessService(repo repositories.IDeleteAccessRepo) *DeleteAccessService {
	return &DeleteAccessService{
		repo: repo,
	}
}

func (s *DeleteAccessService) DeleteAccessUser(request models.DeleteUserAccessRequest, response chan models.DeleteUserAccessResponse) {
	result := s.repo.DeleteAccessUser(request)

	response <- models.DeleteUserAccessResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
