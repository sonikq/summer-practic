package models

type DeleteItemRequest struct {
	ItemID int `json:"item_id"`
}

type DeleteItemResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
