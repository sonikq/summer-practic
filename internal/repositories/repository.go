package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
)

type IAddAccessRepo interface {
	AddAccessUser(request models.AddUserAccessRequest) models.AddUserAccessResponse
}

type ICheckAccessRepo interface {
	CheckAccessUser(request models.CheckUserAccessRequest) models.CheckUserAccessResponse
}

type IDeleteAccessRepo interface {
	DeleteAccessUser(request models.DeleteUserAccessRequest) models.DeleteUserAccessResponse
}

type IUpdateAccessRepo interface {
	UpdateAccessUser(request models.UpdateUserAccessRequest) models.UpdateUserAccessResponse
}

type Repository struct {
	IAddAccessRepo
	ICheckAccessRepo
	IDeleteAccessRepo
	IUpdateAccessRepo
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		IAddAccessRepo:    NewAddAccessRepo(db),
		ICheckAccessRepo:  NewCheckAccessRepo(db),
		IDeleteAccessRepo: NewDeleteAccessRepo(db),
		IUpdateAccessRepo: NewUpdateAccessRepo(db),
	}
}
