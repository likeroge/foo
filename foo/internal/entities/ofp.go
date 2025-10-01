package entities

type OFP struct {
	FileName     string   `json:"fileName"`
	IcaoFrom     string   `json:"icaoFrom"`
	IcaoTo       string   `json:"icaoTo"`
	ETD          string   `json:"etd"`
	ATD          string   `json:"atd"`
	ETA          string   `json:"eta"`
	ATA          string   `json:"ata"`
	FlightNumber string   `json:"flightNumber"`
	DOF          string   `json:"dof"`
	AllAirports  []string `json:"allAirports"`
	AllFirs      []string `json:"allFirs"`
	RegNumber    string   `json:"regNumber"`
}
