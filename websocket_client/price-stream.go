package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{
		Scheme: "wss",
		Host:   "stream.binance.com:9443",
		Path:   "/ws/btcusdt@kline_1m",
	}
	fmt.Println(u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(resp.Status)
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", string(msg))
	}
}
