package exchange

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

var (
	cachedRate float64
	lastUpdate time.Time
	mu         sync.Mutex
)

func GetCNYRate() (float64, error) {
	mu.Lock()

	defer mu.Unlock()

	today := time.Now().Format("2006-01-02")
	last := lastUpdate.Format("2006-01-02")

	if today == last && cachedRate > 0 {
		return cachedRate, nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.cbr.ru/scripts/XML_daily.asp", nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; PoisonBot/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	decoder := xml.NewDecoder(strings.NewReader(string(body)))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return 0, err
	}

	for _, v := range valCurs.Valutes {
		if v.CharCode == "CNY" {
			rateStr := strings.Replace(v.Value, ",", ".", 1)
			rate, err := strconv.ParseFloat(rateStr, 64)
			if err != nil {
				return 0, err
			}
			rate = roundToNearest(rate, 0.05)
			cachedRate = rate
			lastUpdate = time.Now()

			return rate, nil
		}
	}

	return 0, fmt.Errorf("курс юаня не найден :(")
}

func roundToNearest(value, step float64) float64 {
	return float64(int((value/step+1)+0.5)) * step
}
