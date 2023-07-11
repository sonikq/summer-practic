package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type IAddItemService interface {
	AddItem(request models.AddItemRequest, response chan models.AddItemResponse)
}

type IEditTableService interface {
	EditTable(request models.EditTableRequest, response chan models.EditTableResponse)
}

type IDeleteItemService interface {
	DeleteItem(request models.DeleteItemRequest, response chan models.DeleteItemResponse)
}

type IUpdateItemService interface {
	UpdateItem(request models.UpdateItemRequest, response chan models.UpdateItemResponse)
}

type Service struct {
	IAddItemService
	IEditTableService
	IDeleteItemService
	IUpdateItemService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IAddItemService:    NewAddItemService(repos.IAddItemRepo),
		IEditTableService:  NewEditTableService(repos.IEditTableRepo),
		IDeleteItemService: NewDeleteItemService(repos.IDeleteAccessRepo),
		IUpdateItemService: NewUpdateItemService(repos.IUpdateAccessRepo),
	}
}
