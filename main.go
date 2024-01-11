package main

import (
	"log"

	"github.com/MrAinslay/OpapAPI/internal/opap"
)

func main() {
	client := opap.NewClient(nil)

	draw, _, err := client.Draws.ByDate(opap.Pοwerspin, "2024-01-11")
	if err != nil {
		log.Println(err)
	}

	log.Println(len(draw.Content))
}
