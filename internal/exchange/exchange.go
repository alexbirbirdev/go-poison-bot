package exchange

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func GetCNYRate() (float64, error) {
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
			rate := strings.Replace(v.Value, ",", ".", 1)
			return strconv.ParseFloat(rate, 64)
		}
	}

	return 0, fmt.Errorf("курс юаня не найден :(")
}
