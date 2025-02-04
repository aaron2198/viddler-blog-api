package main

import (
	"flag"
	"log"

	"gitlab.aaronhess.xyz/viddler/viddler-blog-api/internal/viddler"
)

var videoURL *string = flag.String("url", "", "YouTube video URL")

func main() {
	app, err := viddler.New()
	if err != nil {
		log.Fatal(err)
	}
	app.Cli(*videoURL)
}
