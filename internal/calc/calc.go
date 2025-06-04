package calc

import (
	"os"
	"strconv"
)

func YuanToRub(yuan, rate float64) float64 {
	deltaStr := os.Getenv("EXCHANGE_YUAN_DELTA")
	delta, err := strconv.ParseFloat(deltaStr, 64)
	if err != nil {
		delta = 0
	}
	exchangeRate := rate + delta
	delivery := 2000.0
	comission := 1000.0

	rubPrice := yuan*exchangeRate + delivery + comission
	return rubPrice
}
