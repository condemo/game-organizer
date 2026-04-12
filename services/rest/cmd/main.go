package main

import (
	"flag"

	"github.com/condemo/game-organizer/services/rest/api"
)

func main() {
	addr := flag.String("addr", ":7000", "service port")
	flag.Parse()

	server := api.NewApiServer(*addr)
	server.Run()
}
