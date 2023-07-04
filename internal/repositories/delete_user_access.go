package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type DeleteAccessRepo struct {
	db *db.DB
}

func NewDeleteAccessRepo(db *db.DB) *DeleteAccessRepo {
	return &DeleteAccessRepo{
		db: db,
	}
}

func (r *DeleteAccessRepo) DeleteAccessUser(request models.DeleteUserAccessRequest) models.DeleteUserAccessResponse {
	if err := r.db.DeleteUserAccess(request.UserID); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.DeleteUserAccessResponse{
			Code:   500,
			Status: deleteAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.DeleteUserAccessResponse{
		Code:   200,
		Status: deleteAccessUserSuccess,
		Error:  nil,
	}
}
