package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type IAddItemService interface {
	AddItem(request models.AddItemRequest, response chan models.AddItemResponse)
}

//type ICheckAccessService interface {
//	CheckAccessUser(request models.CheckUserAccessRequest, response chan models.CheckUserAccessResponse)
//}

type IDeleteItemService interface {
	DeleteItem(request models.DeleteItemRequest, response chan models.DeleteItemResponse)
}

type IUpdateItemService interface {
	UpdateItem(request models.UpdateItemRequest, response chan models.UpdateItemResponse)
}

type Service struct {
	IAddItemService
	//ICheckAccessService
	IDeleteItemService
	IUpdateItemService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IAddItemService: NewAddItemService(repos.IAddItemRepo),
		//ICheckAccessService:  NewCheckAccessService(repos.ICheckAccessRepo),
		IDeleteItemService: NewDeleteItemService(repos.IDeleteAccessRepo),
		IUpdateItemService: NewUpdateItemService(repos.IUpdateAccessRepo),
	}
}
