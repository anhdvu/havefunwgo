package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gorilla/websocket"
)

type Delta struct {
	StartTime int64   `json:"st"`
	EndTime   int64   `json:"et"`
	Price     float64 `json:"p"`
	TotalVol  float64 `json:"tv"`
	ActualVol float64 `json:"av"`
	Ratio     float64 `json:"r"`
	Trades    int     `json:"n"`
}

func main() {
	// var sentry float64 = 18165.79
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:7878",
		Path:   "/ws",
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", msg)

		// d := makeDelta(msg)
		// f, _ := os.OpenFile("dynamic_delta.txt", os.O_WRONLY, 755)
		// defer f.Close()
		// d.ToJSON(f)
	}

	// _, msg, err := c.ReadMessage()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(msg))

	// startTime := time.Now().UnixNano()
	// ps := &kline{}
	// _ = json.Unmarshal(msg, ps)
	// price1, _ := strconv.ParseFloat(ps.Kline.Close, 64)
	// fmt.Println(price1)
	// duration := time.Now().UnixNano() - startTime
	// fmt.Println(duration)

	// startTime = time.Now().UnixNano()
	// price2, _ := getPrice(msg)
	// fmt.Println(price2)
	// duration = time.Now().UnixNano() - startTime
	// fmt.Println(duration)

}

func makeDelta(msg []byte) *Delta {
	re := regexp.MustCompile(`"[tTcvnV]":"?(\d+\.?\d+)`)
	matches := re.FindAllSubmatch(msg, -1)
	startTime, _ := strconv.ParseInt(string(matches[0][1]), 10, 64)
	endTime, _ := strconv.ParseInt(string(matches[1][1]), 10, 64)
	price, _ := strconv.ParseFloat(string(matches[2][1]), 64)
	totalVol, _ := strconv.ParseFloat(string(matches[3][1]), 64)
	actualVol, _ := strconv.ParseFloat(string(matches[5][1]), 64)
	ratio := actualVol / totalVol
	trades, _ := strconv.Atoi(string(matches[4][1]))
	return &Delta{
		StartTime: startTime,
		EndTime:   endTime,
		Price:     price,
		TotalVol:  totalVol,
		ActualVol: actualVol,
		Ratio:     ratio,
		Trades:    trades,
	}
}

func (delta *Delta) ToJSON(w io.Writer) {
	d := json.NewEncoder(w)
	err := d.Encode(delta)
	if err != nil {
		log.Panic(err)
	}
}
