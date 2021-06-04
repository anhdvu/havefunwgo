package main

import (
	"fmt"
	"net/url"
)

type ex struct {
	BaseUrl url.URL `json:"base_url,omitempty"`
}

func main() {
	queryParams := map[string][]string{
		"name":   {"bob", "alice"},
		"gender": {"trans"},
	}

	e := &ex{
		BaseUrl: url.URL{
			Scheme: "https",
			Host:   "api.binance.com",
		},
	}

	endpoint := e.BaseUrl
	endpoint.Path = "/api/v3/time"
	endpoint.ForceQuery = false
	endpoint.RawQuery = queryBuilder(queryParams)

	fmt.Printf("%#v\n", endpoint)

	fmt.Printf("Base URL --- %#v\n", e.BaseUrl)
	fmt.Printf("Endpoint --- %#v\n", endpoint)
	fmt.Println(endpoint.String())
}

func queryBuilder(m map[string][]string) string {
	var query string

	for k, v := range m {
		for _, e := range v {
			query += k + "=" + e
			query += "&"
		}

	}

	query = query[:len(query)-1]
	return query
}
