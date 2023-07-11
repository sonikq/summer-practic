package models

type EditTableRequest struct {
	NewObj string      `json:"new_col"`
	Type   interface{} `json:"type"`
}

type EditTableResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
