package models

type AddItemRequest struct {
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

type Places struct {
	DIN int `json:"placeDIN"`
	U   int `json:"placeU"`
}

type Ports struct {
	AIN int `json:"portAIN"`
	AOT int `json:"portAOT"`
	DIN int `json:"portDIN"`
	DOT int `json:"portDOT"`
	EHT int `json:"portEHT"`
}

type Connectors struct {
	RJ    int `json:"connRJ"`
	PIN20 int `json:"connPIN20"`
	PIN50 int `json:"connPIN50"`
	DB37  int `json:"connDB37"`
	DB62  int `json:"connDB62"`
	SNC55 int `json:"connSNC55"`
}

type Wires struct {
	MGTF int `json:"wireMGTF"`
	CAT6 int `json:"wireCAT6"`
}

type Bracings struct {
	Washer int `json:"braceWasher"`
	Nut    int `json:"braceNut"`
	Bolt   int `json:"braceBolt"`
}

type AddItemResponse struct {
	Code   int
	Status string `json:"status"`
	Error  *Err   `json:"error"`
}
