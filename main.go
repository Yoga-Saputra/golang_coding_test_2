package main

import (
	"codingTest/routes"
	"time"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

func main() {
	overseer.Run(overseer.Config{
		Program: gracefulStart,
		Address: ":2001",
		Fetcher: &fetcher.HTTP{
			URL:      "http://localhost",
			Interval: 1 * time.Second,
		},
		Debug: true,
	})
}

func gracefulStart(state overseer.State) {
	routes.InitApi(state)
}
