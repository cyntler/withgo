package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	coingeckolib "github.com/superoo7/go-gecko/v3"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	coingecko := coingeckolib.NewClient(httpClient)
	ping, err := coingecko.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ping.GeckoSays)
}
