package models

type DeleteItemRequest struct {
	Designation string `json:"designation"`
}

type DeleteItemResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
