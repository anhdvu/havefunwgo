package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

var data = `<?xml version="1.0"?>
<methodCall>
  <methodName>DeductReversal</methodName>
  <params>
    <param>
      <value>
        <string>0008323836</string>
      </value>
    </param>
    <param>
      <value>
        <string>T2P900000000258</string>
      </value>
    </param>
    <param>
      <value>
        <int>12300</int>
      </value>
    </param>
    <param>
      <value>
        <string>DEDUCT REVERSAL        HCM           VNM</string>
      </value>
    </param>
    <param>
      <value>
        <string></string>
      </value>
    </param>
    <param>
      <value>
        <string>880704</string>
      </value>
    </param>
    <param>
      <value>
        <dateTime.iso8601>20200325T10:15:24</dateTime.iso8601>
      </value>
    </param>
    <param>
      <value>
        <string>1234</string>
      </value>
    </param>
    <param>
      <value>
        <dateTime.iso8601>20200325T16:00:00</dateTime.iso8601>
      </value>
    </param>
    <param>
      <value>
        <string>893A554F9A23B81C688C3DBB23C59523316835AA</string>
      </value>
    </param>
  </params>
</methodCall>`

// Payload contains the outer structure
type Payload struct {
	MethodName methodName `xml:"methodName"`
	Parameters *paramDict `xml:",any"`
}

type methodName string

// paramDict contains data in form of "key: value" pairs
type paramDict map[string]string

var keys []string

// UnmarshalXML is custom
func (mn *methodName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Local != "methodName" {
		return fmt.Errorf("wanted \"methodName\" as start element but got %q", start.Name.Local)
	}

	var c string
	if err := d.DecodeElement(&c, &start); err != nil {
		return err
	}

	switch c {
	case "AdministrativeMessage":
		keys = []string{"terminal", "reference", "messageType", "klv", "txnID", "txnDate"}
	case "Deduct":
		keys = []string{"terminal", "reference", "amount", "narrative", "txnType", "klv", "txnID", "txnDate"}
	case "DeductAdjustment", "LoadAdjustment", "LoadReversal", "DeductReversal":
		keys = []string{"termianl", "reference", "amount", "narrative", "klv", "refTxnID", "refTxnDate", "txnID", "txnDate"}
	default:
		return fmt.Errorf("unknown method name, %q", c)
	}

	*mn = methodName(c)
	return nil
}

// UnmarshalXML is custom
func (params *paramDict) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Local != "params" {
		return fmt.Errorf("wanted \"params\" as start element but got %q", start.Name.Local)
	}

	values := make([]string, 0)
	for {
		t, err := d.Token()
		if t == nil {
			break
		}
		if err != nil {
			return err
		}

		switch elem := t.(type) {
		case xml.StartElement:
			tag := elem.Name.Local
			switch tag {
			case "string", "int", "i4", "dateTime.iso8601":
				var c string
				if err := d.DecodeElement(&c, &elem); err != nil {
					return err
				}
				values = append(values, c)

			case "params", "param", "value":
			default:
				return fmt.Errorf("unknown element %q", tag)
			}
		}
	}
	out, err := zip(keys, values)
	if err != nil {
		return err
	}
	*params = out

	return nil
}

func main() {
	pl := &Payload{}
	err := xml.Unmarshal([]byte(data), pl)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	// fmt.Printf("%+v\n\n", pl)
	// fmt.Println((*pl.Parameters)["klv"])

	jsonpayload, err := json.MarshalIndent(pl, "", "  ")
	fmt.Println(string(jsonpayload))
}

func zip(s1, s2 []string) (map[string]string, error) {
	length := 0
	if len(s1) > len(s2) {
		length = len(s2)
	} else {
		length = len(s1)
	}

	if length == 0 {
		return nil, fmt.Errorf("an input slice has length of %v", length)
	}
	result := make(map[string]string)
	for i := 0; i < length; i++ {
		result[s1[i]] = s2[i]
	}

	return result, nil
}
