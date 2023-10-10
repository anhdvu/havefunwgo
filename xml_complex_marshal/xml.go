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
		Int    string   `xml:"int,omitempty"`
		String string   `xml:"string,omitempty"`
		Value  []*value `xml:"array>data>value,omitempty"`
	} `xml:"value,omitempty"`
}

type value struct {
	Member [2]struct {
		Name  string `xml:"name"`
		Value string `xml:"value"`
	} `xml:"struct>member"`
}

func main() {
	r := methodResponse{}

	m1 := member{Name: "resultCode"}
	m1.Value.Int = "1"
	m1.Value.Value = nil

	m2 := member{Name: "activationMethods"}
	v1 := value{}
	v1.Member[0].Name = "type"
	v1.Member[0].Value = "1"
	v1.Member[1].Name = "value"
	v1.Member[1].Value = "+1(###) ### 4567"

	v2 := value{}
	v2.Member[0].Name = "type"
	v2.Member[0].Value = "2"
	v2.Member[1].Name = "value"
	v2.Member[1].Value = "joh***n@anymail.com"

	m2.Value.Value = append(m2.Value.Value, &v1, &v2)

	r.Params.Member = append(r.Params.Member, m1, m2)

	b, err := xml.Marshal(r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", b)
}
