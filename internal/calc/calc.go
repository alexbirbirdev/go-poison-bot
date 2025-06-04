package calc

import (
	"os"
	"strconv"
)

func YuanToRub(yuan, rate float64) float64 {
	deltaStr := os.Getenv("EXCHANGE_YUAN_DELTA")
	delta, err := strconv.ParseFloat(deltaStr, 64)
	if err != nil {
		delta = 0.0
	}
	delStr := os.Getenv("CALC_DELIVERY")
	delivery, err := strconv.ParseFloat(delStr, 64)
	if err != nil {
		delivery = 0.0
	}
	comStr := os.Getenv("CALC_COMMISSION")
	commission, err := strconv.ParseFloat(comStr, 64)
	if err != nil {
		commission = 0.0
	}
	exchangeRate := rate + delta

	rubPrice := yuan*exchangeRate + delivery + commission
	return rubPrice
}
