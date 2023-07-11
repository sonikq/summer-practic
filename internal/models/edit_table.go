package models

type EditTableRequest struct {
	NewObj string      `json:"new_obj"`
	Type   interface{} `json:"type"`
}

type EditTableResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
