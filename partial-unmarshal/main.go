package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	var input struct {
		Side string `json:"side"`
		Rest json.RawMessage
	}

	pl := `{"side":"buy","user":"dace","cex":"kucoin"}`

	buf := bytes.NewBuffer([]byte(pl))

	err := json.NewDecoder(buf).Decode(&input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", input)
}
