package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type UpdateItemRepo struct {
	db *db.DB
}

func NewUpdateItemRepo(db *db.DB) *UpdateItemRepo {
	return &UpdateItemRepo{
		db: db,
	}
}

func (r *UpdateItemRepo) UpdateItem(request models.UpdateItemRequest) models.UpdateItemResponse {
	if err := r.db.UpdateItem(
		request.Name, request.Designation, request.Link, request.Quantity, request.Places.DIN, request.Places.U, request.Ports.AIN, request.Ports.AOT, request.Ports.DIN, request.Ports.DOT,
		request.Ports.EHT, request.PCI, request.Connectors.RJ, request.Connectors.PIN20, request.Connectors.PIN50, request.Connectors.DB37, request.Connectors.DB62,
		request.Connectors.SNC55, request.Wires.MGTF, request.Wires.CAT6, request.Length, request.Bracings.Washer, request.Bracings.Nut, request.Bracings.Bolt,
		request.Power, request.Voltage, request.Price,
	); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.UpdateItemResponse{
			Code:   500,
			Status: updateAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.UpdateItemResponse{
		Code:   200,
		Status: updateAccessUserSuccess,
		Error:  nil,
	}
}
