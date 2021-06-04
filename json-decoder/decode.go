package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Result struct {
	IsoRequest             string
	IsoResponse            string
	IsoResponsePacket      map[string]string
	ResultCode             int
	ResultText             string
	ReversalIsoRequest     string
	ReversalIsoResponse    string
	ReversalWalletRequest  string
	ReversalWalletResponse string
	WalletRequest          string
	WalletResponse         string
}

func (r *Result) FromJSON(b io.Reader) error {
	d := json.NewDecoder(b)
	return d.Decode(r)
}

func main() {
	raw := `{
		"isoRequest": "<?xml version=\"1.0\"?><methodCall><methodName>deductAuth</methodName><params><param><value><int>2100</int></value></param><param><value><int>2804826</int></value></param><param><value><string>PROD_MSTR_1A</string></value></param><param><value><string>iso_var_mastercard_cis</string></value></param><param><value><struct><member><name>11</name><value><string>449210</string></value></member><member><name>22</name><value><string>05</string></value></member><member><name>99</name><value><string>3</string></value></member><member><name>55</name><value><string>1</string></value></member><member><name>12</name><value><string>20210530190949</string></value></member><member><name>24</name><value><string>100</string></value></member><member><name>35</name><value><string>5338485258218895=301112203395</string></value></member><member><name>26</name><value><string>7299</string></value></member><member><name>37</name><value><string>72582614060</string></value></member><member><name>39</name><value><string>0</string></value></member><member><name>2</name><value><string>5338485258218895</string></value></member><member><name>3</name><value><string>003000</string></value></member><member><name>4</name><value><string>8402000000000500</string></value></member><member><name>5</name><value><string>8402000000000500</string></value></member><member><name>6</name><value><string>8402000000000500</string></value></member><member><name>7</name><value><string>0530190949</string></value></member><member><name>41</name><value><string>12345678</string></value></member><member><name>52</name><value><string>95024690E3FA2230</string></value></member><member><name>42</name><value><string>12345</string></value></member><member><name>32</name><value><string>2202</string></value></member><member><name>43</name><value><string>47002C52003232504F532044454455435420524556455253414C2020203133434F4D50414E494F4E20202020415220</string></value></member></struct></value></param></params></methodCall>",
		"isoResponse": "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<methodResponse><params><param><value><struct><member><name>11</name><value><string>449210</string></value></member><member><name>22</name><value><string>05</string></value></member><member><name>99</name><value><string>3</string></value></member><member><name>55</name><value><string>1</string></value></member><member><name>12</name><value><string>20210530190949</string></value></member><member><name>24</name><value><string>100</string></value></member><member><name>26</name><value><string>7299</string></value></member><member><name>37</name><value><string>72582614060</string></value></member><member><name>39</name><value><string>0000</string></value></member><member><name>2</name><value><string>5338485258218895</string></value></member><member><name>3</name><value><string>003000</string></value></member><member><name>4</name><value><string>8402000000000500</string></value></member><member><name>5</name><value><string>8402000000000500</string></value></member><member><name>6</name><value><string>8402000000000500</string></value></member><member><name>7</name><value><string>0530190949</string></value></member><member><name>41</name><value><string>12345678</string></value></member><member><name>42</name><value><string>12345</string></value></member><member><name>32</name><value><string>2202</string></value></member><member><name>43</name><value><string>47002C52003232504F532044454455435420524556455253414C2020203133434F4D50414E494F4E20202020415220</string></value></member><member><name>38</name><value><string>528700</string></value></member><member><name>54</name><value><string>3002D8402000000000000</string></value></member></struct></value></param></params></methodResponse>",
		"isoResponsePacket": {
			"11": "R449210",
			"12": "R20210530191041",
			"2": "R5338485258218895",
			"22": "R05",
			"24": "R400",
			"26": "R7299",
			"3": "R003000",
			"32": "R2202",
			"37": "R72582614060",
			"39": "R4000",
			"4": "R8402000000000500",
			"41": "R12345678",
			"42": "R12345",
			"43": "R47002C52003232504F532044454455435420524556455253414C2020203133434F4D50414E494F4E20202020415220",
			"5": "R8402000000000500",
			"55": "R1",
			"56": "R00000000004492102021053019094900000002202",
			"6": "R8402000000000500",
			"7": "R0530190949",
			"99": "R3"
		},
		"resultCode": 1,
		"resultText": "Success",
		"reversalIsoRequest": "<?xml version=\"1.0\"?><methodCall><methodName>deductAuthReverse</methodName><params><param><value><int>2100</int></value></param><param><value><int>6380045</int></value></param><param><value><string>PROD_MSTR_1A</string></value></param><param><value><string>iso_var_mastercard_cis</string></value></param><param><value><struct><member><name>11</name><value><string>449210</string></value></member><member><name>22</name><value><string>05</string></value></member><member><name>99</name><value><string>3</string></value></member><member><name>55</name><value><string>1</string></value></member><member><name>12</name><value><string>20210530191041</string></value></member><member><name>56</name><value><string>00000000004492102021053019094900000002202</string></value></member><member><name>24</name><value><string>400</string></value></member><member><name>26</name><value><string>7299</string></value></member><member><name>37</name><value><string>72582614060</string></value></member><member><name>39</name><value><string>0</string></value></member><member><name>2</name><value><string>5338485258218895</string></value></member><member><name>3</name><value><string>003000</string></value></member><member><name>4</name><value><string>8402000000000500</string></value></member><member><name>5</name><value><string>8402000000000500</string></value></member><member><name>6</name><value><string>8402000000000500</string></value></member><member><name>7</name><value><string>0530190949</string></value></member><member><name>41</name><value><string>12345678</string></value></member><member><name>52</name><value><string>95024690E3FA2230</string></value></member><member><name>42</name><value><string>12345</string></value></member><member><name>32</name><value><string>2202</string></value></member><member><name>43</name><value><string>47002C52003232504F532044454455435420524556455253414C2020203133434F4D50414E494F4E20202020415220</string></value></member></struct></value></param></params></methodCall>",
		"reversalIsoResponse": "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<methodResponse><params><param><value><struct><member><name>11</name><value><string>449210</string></value></member><member><name>22</name><value><string>05</string></value></member><member><name>99</name><value><string>3</string></value></member><member><name>55</name><value><string>1</string></value></member><member><name>12</name><value><string>20210530191041</string></value></member><member><name>56</name><value><string>00000000004492102021053019094900000002202</string></value></member><member><name>24</name><value><string>400</string></value></member><member><name>26</name><value><string>7299</string></value></member><member><name>37</name><value><string>72582614060</string></value></member><member><name>39</name><value><string>4000</string></value></member><member><name>2</name><value><string>5338485258218895</string></value></member><member><name>3</name><value><string>003000</string></value></member><member><name>4</name><value><string>8402000000000500</string></value></member><member><name>5</name><value><string>8402000000000500</string></value></member><member><name>6</name><value><string>8402000000000500</string></value></member><member><name>7</name><value><string>0530190949</string></value></member><member><name>41</name><value><string>12345678</string></value></member><member><name>42</name><value><string>12345</string></value></member><member><name>32</name><value><string>2202</string></value></member><member><name>43</name><value><string>47002C52003232504F532044454455435420524556455253414C2020203133434F4D50414E494F4E20202020415220</string></value></member></struct></value></param></params></methodResponse>",
		"reversalWalletRequest": "<?xml version=\"1.0\"?><methodCall><methodName>DeductReversal</methodName><params><param><value><string>0061796788</string></value></param><param><value><string>test.1234</string></value></param><param><value><int>500</int></value></param><param><value><string>POS DEDUCT REVERSAL    COMPANION     AR </string></value></param><param><value><string></string></value></param><param><value><string>449210</string></value></param><param><value><dateTime.iso8601>20210530T19:09:49</dateTime.iso8601></value></param><param><value><string>14827</string></value></param><param><value><dateTime.iso8601>20210530T21:08:31</dateTime.iso8601></value></param><param><value><string>4C6E620113DC14EC6ECB1E7B6331118D1C334D69841459F3A8A4A4C297CED86E</string></value></param></params></methodCall>",
		"reversalWalletResponse": "<methodResponse><params><param><value><struct><member><name>resultCode</name><value><int>1</int></value></member></struct></value></param></params></methodResponse>",
		"walletRequest": "<?xml version=\"1.0\"?><methodCall><methodName>Deduct</methodName><params><param><value><string>0061796788</string></value></param><param><value><string>test.1234</string></value></param><param><value><int>500</int></value></param><param><value><string>POS DEDUCT REVERSAL    COMPANION     AR </string></value></param><param><value><string>00</string></value></param><param><value><string>00215889739500000002004120000000005000100002604729903711725826140600410812345678042051234504800049038400500230052000850025003EMV25110MasterCard252010253048895254002550025801025903ARG26101027000</string></value></param><param><value><string>449210</string></value></param><param><value><dateTime.iso8601>20210530T19:09:49</dateTime.iso8601></value></param><param><value><string>66BFCB015B69118F711FCD1FE161107CC3393E55327E773F333C3755AFEBF39D</string></value></param></params></methodCall>",
		"walletResponse": "<methodResponse><params><param><value><struct><member><name>resultCode</name><value><int>1</int></value></member></struct></value></param></params></methodResponse>"
	}`

	buf := bytes.NewBufferString(raw)

	r := &Result{}

	r.FromJSON(buf)

	fmt.Printf("%+v\n", r)
}
