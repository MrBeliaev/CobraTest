package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type GetPriceRequest struct {
	Pairs []string `json:"pairs"`
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	var pairsStr string
	if r.Body != nil && r.Method == "POST" {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		var gpr GetPriceRequest
		err = json.Unmarshal(bodyBytes, &gpr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, v := range gpr.Pairs {
			pairsStr += v
			if i < len(gpr.Pairs)-1 {
				pairsStr += ","
			}
		}
	}
	if pairsStr == "" && r.Method == "GET" {
		pairsStr = r.URL.Query().Get("pairs")
	}
	if pairsStr != "" {
		result, err := GetBinancePrice(&pairsStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, result)
		return
	}
	http.Error(w, "Not Found", http.StatusBadRequest)
	return
}
