package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

const BINANCE_URL = "https://api.binance.com/api/v3/ticker/price"

type BinanceResponce struct {
	Symbol string
	Price  string
}

func GetBinancePrice(pairs *string) (string, error) {
	pairsArr := strings.Split(*pairs, ",")
	result := make(map[string]float64)
	for _, v := range pairsArr {
		var br BinanceResponce
		pair := strings.Split(v, "-")
		url := BINANCE_URL + "?symbol=" + pair[0] + pair[1]
		rate, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer rate.Body.Close()
		err = json.NewDecoder(rate.Body).Decode(&br)
		if err != nil {
			return "", err
		}
		floatPrice, err := strconv.ParseFloat(br.Price, 64)
		if err != nil {
			return "", err
		}
		result[v] = floatPrice
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(jsonResult), nil
}
