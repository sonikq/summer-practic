package repositories

import (
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/models"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/db"
	"strings"
)

type AddItemRepo struct {
	db *db.DB
}

func NewAddItemRepo(db *db.DB) *AddItemRepo {
	return &AddItemRepo{
		db: db,
	}
}

func (r *AddItemRepo) AddItem(request models.AddItemRequest) models.AddItemResponse {
	if err := r.db.AddItem(
		request.Name, request.Designation, request.Link, request.Quantity, request.Places.DIN, request.Places.U, request.Ports.AIN, request.Ports.AOT, request.Ports.DIN, request.Ports.DOT,
		request.Ports.EHT, request.PCI, request.Connectors.RJ, request.Connectors.PIN20, request.Connectors.PIN50, request.Connectors.DB37, request.Connectors.DB62,
		request.Connectors.SNC55, request.Wires.MGTF, request.Wires.CAT6, request.Length, request.Bracings.Washer, request.Bracings.Nut, request.Bracings.Bolt,
		request.Power, request.Voltage, request.Price); err != nil {
		e := strings.Split(err.Error(), ": ")
		src, msg := e[0], e[1]
		return models.AddItemResponse{
			Code:   400,
			Status: addAccessUserFail,
			Error: &models.Err{
				Source:  src,
				Message: msg,
			},
		}
	}

	return models.AddItemResponse{
		Code:   201,
		Status: addAccessUserSuccess,
		Error:  nil,
	}
}
