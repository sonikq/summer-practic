package models

type DeleteUserAccessRequest struct {
	UserID int `json:"user_id"`
}

type DeleteUserAccessResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
