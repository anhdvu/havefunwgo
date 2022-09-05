package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

type msg struct {
	ID      int
	Content string
}

func main() {
	fmt.Println("API service.")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	var err error
	channels, err := rdb.PubSubChannels(ctx, "*").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(channels)

	var count int
	var m msg
	var n []byte
	var code int64

	pairs := []string{"btc-usdt", "eth-usdt", "luna-usdt", "one-usdt", "near-usdt", "mina-usdt"}
	for {
		m = msg{ID: count, Content: fmt.Sprintf("msgID %d - %s.", count, randomPair(pairs))}
		n, _ = json.Marshal(m)

		code, err = rdb.Publish(ctx, "trade-requests", n).Result()
		if err != nil {
			fmt.Println(err)
			return
		}

		if code == 0 {
			fmt.Println("channel has been closed.")
			return
		}

		time.Sleep(3 * time.Second)
		count++
	}
}

func randomPair(pairs []string) string {
	return pairs[rand.Intn(len(pairs))]
}
