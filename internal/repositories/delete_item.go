package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type DeleteItemRepo struct {
	db *db.DB
}

func NewDeleteItemRepo(db *db.DB) *DeleteItemRepo {
	return &DeleteItemRepo{
		db: db,
	}
}

func (r *DeleteItemRepo) DeleteItem(request models.DeleteItemRequest) models.DeleteItemResponse {
	if err := r.db.DeleteItem(request.ItemID); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.DeleteItemResponse{
			Code:   500,
			Status: deleteAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.DeleteItemResponse{
		Code:   200,
		Status: deleteAccessUserSuccess,
		Error:  nil,
	}
}
