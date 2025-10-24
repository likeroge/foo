package entities

type OFP struct {
	Id       int    `json:"id"`
	IcaoFrom string `json:"icaoFrom"`
	IcaoTo   string `json:"icaoTo"`
	ETD      string `json:"etd"`
	// ATD          string   `json:"atd"`
	ETA string `json:"eta"`
	// ATA          string   `json:"ata"`
	FlightNumber string   `json:"flightNumber"`
	Rte          string   `json:"rte"`
	DOF          string   `json:"dof"`
	AllAirports  []string `json:"allAirports,omitempty"`
	AllFirs      []string `json:"allFirs,omitempty"`
	RegNumber    string   `json:"regNumber"`
	AltAirports  []string `json:"altAirports,omitempty"`
	CreatedAt    string   `json:"createdAt,omitempty"`
	UpdatedAt    string   `json:"updatedAt,omitempty"`
	Distance     int      `json:"dist"`
	Wind         string   `json:"wind"`
	FuelFlow     int      `json:"fuelFlow"`
	TripFuel     int      `json:"tripFuel"`
	FlightTime   string   `json:"flightTime"`
}
