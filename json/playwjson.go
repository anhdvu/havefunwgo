package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	_ "os"
	"time"
)

type Sentry struct {
	Time       string
	Prediction float64 `json:"pred_price"`
	Actual     float64 `json:"actual_price;omitempty"`
}

type SentryPrediction struct {
	Time       string
	Prediction float64 `json:"predict"`
}

type SentryResponse struct {
	Time  int64
	Value float64
}

type WsResponse struct {
	M string `json:"m"`
	D struct {
		T int64   `json:"t"`
		V float64 `json:"v"`
	} `json:"d"`
}

func main() {
	raw, err := ioutil.ReadFile("checker.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]SentryPrediction, 0)
	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Panic(err)
	}

	response := &WsResponse{}
	response.M = "sentry"
	response.D.T = parseTime(data[0].Time)
	response.D.V = data[0].Prediction

	err = json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		log.Panic(err)
	}
}

func parseTime(ts string) int64 {
	// Jan 2 15:04:05 2006 MST

	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, ts)
	if err != nil {
		log.Panic(err)
	}
	return t.Unix()
}
