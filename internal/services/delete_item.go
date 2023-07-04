package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type DeleteItemService struct {
	repo repositories.IDeleteAccessRepo
}

func NewDeleteItemService(repo repositories.IDeleteAccessRepo) *DeleteItemService {
	return &DeleteItemService{
		repo: repo,
	}
}

func (s *DeleteItemService) DeleteItem(request models.DeleteItemRequest, response chan models.DeleteItemResponse) {
	result := s.repo.DeleteItem(request)

	response <- models.DeleteItemResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
