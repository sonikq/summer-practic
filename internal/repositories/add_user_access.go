package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type AddAccessRepo struct {
	db *db.DB
}

func NewAddAccessRepo(db *db.DB) *AddAccessRepo {
	return &AddAccessRepo{
		db: db,
	}
}

func (r *AddAccessRepo) AddAccessUser(request models.AddUserAccessRequest) models.AddUserAccessResponse {
	if err := r.db.AddItem(
		request.UserID, request.OKUDS, request.ReportIDS, request.ChapterIDS, request.AllowedOPS,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.AddUserAccessResponse{
			Code:   400,
			Status: addAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.AddUserAccessResponse{
		Code:   201,
		Status: addAccessUserSuccess,
		Error:  nil,
	}
}
