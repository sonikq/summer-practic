package models

type EditTableRequest struct {
	ItemID int    `json:"item_id"`
	NewObj string `json:"new_obj"`
}

type EditTableResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
