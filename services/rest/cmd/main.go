package main

import (
	"flag"
	"log"

	"github.com/condemo/game-organizer/services/common/fetch"
	"github.com/condemo/game-organizer/services/rest/api"
)

func main() {
	addr := flag.String("addr", ":7000", "service port")
	flag.Parse()

	// Login
	if err := fetch.IgdbLogin(); err != nil {
		log.Fatalf("twitch login failed: %s", err)
	}

	server := api.NewApiServer(*addr)
	server.Run()
}
