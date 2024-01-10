package main

import (
	"log"

	"github.com/MrAinslay/OpapAPI/internal/opap"
)

func main() {
	client := opap.NewClient(nil)

	draw, _, err := client.Draws.Latest(opap.Pοwerspin)
	if err != nil {
		log.Println(err)
	}

	log.Println(draw)
}
