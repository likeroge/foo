package entities

type Flight struct {
	Dep        string `json:"from"`
	Arr        string `json:"to"`
	Date       string `json:"date"`
	Rte        string `json:"rte"`
	Dist       int    `json:"dist"`
	TripFuel   int    `json:"tripFuel"`
	FlightTime string `json:"flightTime"`
}

func NewFlight() *Flight {
	return &Flight{}
}
