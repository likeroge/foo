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

func (o *OFPParser) ParseOfp() entities.OFP {
	etd := ParseETD(o.ofpString)
	fmt.Println(etd)
	// log.Println(content)
	// re, err := regexp.Compile(`\D+`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // result := re.FindString(content)
	// result := re.FindAllString(content, -1)

	// fmt.Println(result)
	// // regexp.Match(re, []byte("Hello world"))
	return entities.OFP{}
}

// func ParseOfp(content string) entities.OFP {

// }

func ParseETD(content string) string {
	re, err := regexp.Compile(`\d\d\d\d\\\\d\d\d\d`)
	if err != nil {
		fmt.Println(err)
	}
	result := re.FindString(content)
	return result

}
