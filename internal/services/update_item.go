package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type UpdateItemService struct {
	repo repositories.IUpdateAccessRepo
}

func NewUpdateItemService(repo repositories.IUpdateAccessRepo) *UpdateItemService {
	return &UpdateItemService{
		repo: repo,
	}
}

func (s *UpdateItemService) UpdateItem(request models.UpdateItemRequest, response chan models.UpdateItemResponse) {
	result := s.repo.UpdateItem(request)

	response <- models.UpdateItemResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
