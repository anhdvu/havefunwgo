package main

import (
	"encoding/xml"
	"fmt"
)

type methodResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	Params  struct {
		Member []member `xml:"member"`
	} `xml:"params>param>value>struct"`
}

type member struct {
	Name  string `xml:"name"`
	Value struct {
		Int    string `xml:"int,omitempty"`
		String string `xml:"string,omitempty"`
	} `xml:"value"`
}

// GenerateResponse generates XML response.
func GenerateResponse(resultCode, resultMessage string) {
	response := &methodResponse{}

	member1 := member{}
	member1.Name = "resultCode"
	member1.Value.Int = resultCode

	member2 := member{}
	member2.Name = "resultText"
	member2.Value.String = resultMessage

	response.Params.Member = append(response.Params.Member, member1)
	response.Params.Member = append(response.Params.Member, member2)

	responseXML, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n******** Response ********\n%v\n", string(responseXML))
}

func main() {
	resCode := "-8"
	resMess := "Error: Authentication failed."

	GenerateResponse(resCode, resMess)
}
