package services

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/repositories"
)

type IAddAccessService interface {
	AddAccessUser(request models.AddUserAccessRequest, response chan models.AddUserAccessResponse)
}

type ICheckAccessService interface {
	CheckAccessUser(request models.CheckUserAccessRequest, response chan models.CheckUserAccessResponse)
}

type IDeleteAccessService interface {
	DeleteAccessUser(request models.DeleteUserAccessRequest, response chan models.DeleteUserAccessResponse)
}

type IUpdateAccessService interface {
	UpdateAccessUser(request models.UpdateUserAccessRequest, response chan models.UpdateUserAccessResponse)
}

type Service struct {
	IAddAccessService
	ICheckAccessService
	IDeleteAccessService
	IUpdateAccessService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IAddAccessService:    NewAddAccessService(repos.IAddAccessRepo),
		ICheckAccessService:  NewCheckAccessService(repos.ICheckAccessRepo),
		IDeleteAccessService: NewDeleteAccessService(repos.IDeleteAccessRepo),
		IUpdateAccessService: NewUpdateAccessService(repos.IUpdateAccessRepo),
	}
}
