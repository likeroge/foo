package ofp

import (
	"fmt"
	"regexp"

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

// func ParseOfp(content string) entities.OFP {

// }

func ParseFlightData(content string) (*entities.OFP, error) {
	// Get flight Metadata from OFP
	patternMetaData := `([A-Z]{2,3}\d{1,4})\s+([A-Z0-9]{1,7})\s+(\d{2}-\d{2}-\d{2})\s+([A-Z]{4})-([A-Z]{4})\s+(\d{4})/(\d{4})\s+(\d{4})/(\d{4})\s+`

	re, err := regexp.Compile(patternMetaData)
	if err != nil {
		return nil, err
	}
	matchesMetaData := re.FindStringSubmatch(content)
	// for i, match := range matchesMetaData {
	// 	fmt.Printf("Match %d: %s\n", i, match)
	// }

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
	// var ofp entities.OFP = entities.OFP{}
	ofp := &entities.OFP{
		FileName:     "",
		IcaoFrom:     matchesMetaData[4],
		IcaoTo:       matchesMetaData[5],
		ETD:          matchesMetaData[6],
		ATD:          "",
		ETA:          matchesMetaData[8],
		ATA:          "",
		FlightNumber: matchesMetaData[1],
		DOF:          matchesMetaData[3],
		AllAirports:  altnAirports,
		AllFirs:      nil,
		RegNumber:    matchesMetaData[2],
	}
	// fmt.Println(ofp)
	return ofp, nil

}
