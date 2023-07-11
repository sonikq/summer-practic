package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
)

type IAddItemRepo interface {
	AddItem(request models.AddItemRequest) models.AddItemResponse
}

type IEditTableRepo interface {
	EditTable(request models.EditTableRequest) models.EditTableResponse
}

type IDeleteAccessRepo interface {
	DeleteItem(request models.DeleteItemRequest) models.DeleteItemResponse
}

type IUpdateAccessRepo interface {
	UpdateItem(request models.UpdateItemRequest) models.UpdateItemResponse
}

type Repository struct {
	IAddItemRepo
	IEditTableRepo
	IDeleteAccessRepo
	IUpdateAccessRepo
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		IAddItemRepo:      NewAddItemRepo(db),
		IEditTableRepo:    NewEditTableRepo(db),
		IDeleteAccessRepo: NewDeleteItemRepo(db),
		IUpdateAccessRepo: NewUpdateItemRepo(db),
	}
}
