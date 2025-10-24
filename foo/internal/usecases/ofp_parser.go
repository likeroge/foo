package usecases

import (
	"fmt"
	"regexp"
	"strconv"

	"ego.dev21/greetings/internal/entities"
)

type OFPParser struct {
	ofpString string
}

func NewOFPParser(ofpString string) OFPParser {
	return OFPParser{ofpString: ofpString}
}

func (o *OFPParser) ParseOfp() (*entities.OFP, error) {
	result, err := ParseFlightData(o.ofpString)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	return result, nil
}

func ParseFlightData(content string) (*entities.OFP, error) {
	// Get flight Metadata from OFP
	patternMetaData := `([A-Z]{2,3}\d{1,4})\s+([A-Z0-9]{1,7})\s+(\d{2}-\d{2}-\d{2})\s+([A-Z]{4})-([A-Z]{4})\s+(\d{4})/(\d{4})\s+(\d{4})/(\d{4})\s+`

	re, err := regexp.Compile(patternMetaData)
	if err != nil {
		return nil, err
	}
	matchesMetaData := re.FindStringSubmatch(content)
	if matchesMetaData == nil {
		return nil, fmt.Errorf("неверный формат OFP")
	}

	// Get distance, wind and fuel flow
	patternDistWindFlow := `DIST/AIR\s+(\d+)\s+AVG W\/C\s+([P,M]\d+)\s+AVG FF (\d+)`
	re, err = regexp.Compile(patternDistWindFlow)
	if err != nil {
		return nil, err
	}
	matchesDistWindFlow := re.FindStringSubmatch(content)
	if matchesDistWindFlow == nil {
		return nil, fmt.Errorf("неверный формат OFP")
	}

	//Get trip fuel and time
	patternTripFuelTime := `TRIP \w\w\w\w\s+(\d+)\s+(\d+)`
	re, err = regexp.Compile(patternTripFuelTime)
	if err != nil {
		return nil, err
	}
	matchesTripFuelTime := re.FindStringSubmatch(content)
	if matchesTripFuelTime == nil {
		return nil, fmt.Errorf("неверный формат OFP")
	}

	// Get Alternate Airports
	patternAltnAirports := `ALT[1-5]:\s*([A-Z]{4})`

	re, err = regexp.Compile(patternAltnAirports)
	if err != nil {
		return nil, err
	}
	matchesAltnAirports := re.FindAllStringSubmatch(content, -1)
	var altnAirports []string
	for _, match := range matchesAltnAirports {
		altnAirports = append(altnAirports, match[1])
	}

	// Get Route
	patternRte := `ROUTE:\s([A-Z0-9\s]+)`

	re, err = regexp.Compile(patternRte)
	if err != nil {
		return nil, err
	}

	matchesRte := re.FindAllStringSubmatch(content, -1)

	if matchesRte == nil {
		return nil, fmt.Errorf("неверный формат OFP")
	}

	dist, err := strconv.Atoi(matchesDistWindFlow[1])
	if err != nil {
		return nil, err
	}

	fuelFlow, err := strconv.Atoi(matchesDistWindFlow[3])
	if err != nil {
		return nil, err
	}

	tripFuel, err := strconv.Atoi(matchesTripFuelTime[1])
	if err != nil {
		return nil, err
	}

	ofp := &entities.OFP{
		IcaoFrom: matchesMetaData[4],
		IcaoTo:   matchesMetaData[5],
		ETD:      matchesMetaData[6],
		// ATD:          "",
		ETA: matchesMetaData[8],
		// ATA:          "",
		FlightNumber: matchesMetaData[1],
		DOF:          matchesMetaData[3],
		AllAirports:  append(altnAirports, matchesMetaData[4], matchesMetaData[5]),
		AllFirs:      nil,
		RegNumber:    matchesMetaData[2],
		AltAirports:  altnAirports,
		Rte:          matchesRte[0][1],
		Distance:     dist,
		Wind:         matchesDistWindFlow[2],
		FuelFlow:     fuelFlow,
		TripFuel:     tripFuel,
		FlightTime:   matchesTripFuelTime[2],
	}
	return ofp, nil

}
