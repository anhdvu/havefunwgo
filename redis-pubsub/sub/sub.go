package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Logic service.")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	channels := []string{
		"pair:binance:btc-usdt",
		"pair:binance:eth-usdt",
		"pair:binance:ada-usdt",
		"pair:binance:luna-usdt",
		"pair:binance:sol-usdt",
		"pair:binance:near-usdt",
		"pair:binance:bullshit-usdt",
	}
	// Subscribe to a channel
	sub := rdb.Subscribe(ctx, channels...)

	m, err := sub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)

	var count int
	redisChannel := sub.Channel()
	go func() {
		time.Sleep(1 * time.Minute)

		err := sub.Unsubscribe(context.TODO(), "pair:binance:eth-usdt", "pair:binance:ada-usdt")
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		time.Sleep(2 * time.Minute)

		err := sub.Subscribe(context.TODO(), "pair:binance:bnb-usdt", "pair:binance:atom-usdt")
		if err != nil {
			log.Println(err)
		}
	}()

	for msg := range redisChannel {
		fmt.Println(msg.Payload)
		count++
		if count == 10000 {
			fmt.Println("I've processed 100 trade requests.")
			_ = sub.Close()
			return
		}
	}
}
