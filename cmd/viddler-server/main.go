package main

import (
	"flag"
	"log"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/viddler"
)

var port *int = flag.Int("port", 8080, "Port to run the server on")

func main() {
	app, err := viddler.New()
	if err != nil {
		log.Fatal(err)
	}
	app.Server(*port)
}
