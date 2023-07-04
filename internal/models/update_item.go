package models

type UpdateItemRequest struct {
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Link        string `json:"link"`
	Quantity    int    `json:"quantity"`
	Places      Places
	Ports       Ports
	PCI         int
	Connectors  Connectors
	Wires       Wires
	Length      int
	Bracings    Bracings
	Power       float64 `json:"power"`
	Voltage     string  `json:"voltage"`
	Price       float64 `json:"price"`
}

type UpdateItemResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
