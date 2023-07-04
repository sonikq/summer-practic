package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type AddItemService struct {
	repo repositories.IAddItemRepo
}

func NewAddItemService(repo repositories.IAddItemRepo) *AddItemService {
	return &AddItemService{
		repo: repo,
	}
}

func (s *AddItemService) AddItem(request models.AddItemRequest, response chan models.AddItemResponse) {
	result := s.repo.AddItem(request)

	response <- models.AddItemResponse{
		Code:   result.Code,
		Status: result.Status,
		Error:  result.Error,
	}
}
