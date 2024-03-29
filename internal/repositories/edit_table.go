package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type EditTableRepo struct {
	db *db.DB
}

func NewEditTableRepo(db *db.DB) *EditTableRepo {
	return &EditTableRepo{
		db: db,
	}
}

func (r *EditTableRepo) EditTable(request models.EditTableRequest) models.EditTableResponse {
	if err := r.db.EditTable(request.NewObj, request.Type); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.EditTableResponse{
			Code:   500,
			Status: editTableFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.EditTableResponse{
		Code:   200,
		Status: editTableSuccess,
		Error:  nil,
	}
}
