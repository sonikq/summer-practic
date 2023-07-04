package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type UpdateAccessRepo struct {
	db *db.DB
}

func NewUpdateAccessRepo(db *db.DB) *UpdateAccessRepo {
	return &UpdateAccessRepo{
		db: db,
	}
}

func (r *UpdateAccessRepo) UpdateAccessUser(request models.UpdateUserAccessRequest) models.UpdateUserAccessResponse {
	if err := r.db.UpdateUserAccess(
		request.UserID, request.OKUDS, request.ReportIDS, request.ChapterIDS, request.AllowedOPS,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.UpdateUserAccessResponse{
			Code:   500,
			Status: updateAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.UpdateUserAccessResponse{
		Code:   200,
		Status: updateAccessUserSuccess,
		Error:  nil,
	}
}
