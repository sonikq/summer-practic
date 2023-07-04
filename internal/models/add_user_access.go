package models

type AddUserAccessRequest struct {
	UserID     int      `json:"user_id"`
	OKUDS      []int    `json:"okuds"`
	ReportIDS  []int    `json:"report_ids"`
	ChapterIDS []int    `json:"chapter_ids"`
	AllowedOPS []string `json:"allowed_ops"`
}

type AddUserAccessResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
