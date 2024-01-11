package opapapi

import (
	"github.com/MrAinslay/OpapAPI/internal/opap"
)

func GetDraws(date string) (*opap.DrawsByDate, error) {
	client := opap.NewClient(nil)

	draw, _, err := client.Draws.ByDate(opap.Pοwerspin, date)
	if err != nil {
		return &opap.DrawsByDate{}, err
	}

	return draw, nil
}
