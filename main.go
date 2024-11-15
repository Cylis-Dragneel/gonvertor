package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func api(currencyFrom string, currencyTo string, amount float64) float64 {
	apiKey := os.Getenv("API_KEY")
	url := "https://openexchangerates.org/api/latest.json?app_id=" + apiKey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}

	req.Header.Add("accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp.Body.Close()
	var Response struct {
		Disclaimer string
		License    string
		Timestamp  int64
		Base       string
		Rates      map[string]float64
	}
	err = json.NewDecoder(resp.Body).Decode(&Response)
	if err != nil {
		log.Fatal("Error:", err)
	}

	if currencyFrom == "USD" {
		rateFrom := 1.0
		rateTo := Response.Rates[currencyTo]
		converted := float64(math.Round(amount*(rateTo/rateFrom)*100)) / 100
		return converted
	} else {
		rateFrom := Response.Rates[currencyFrom]
		rateTo := Response.Rates[currencyTo]
		converted := float64(math.Round(amount*(rateTo/rateFrom)*100)) / 100
		return converted
	}
}

func main() {
	base := os.Args[1]
	amount, err := strconv.ParseFloat(base, 64)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	currencyFrom := os.Args[2]
	currencyTo := os.Args[3]
	fmt.Println(api(currencyFrom, currencyTo, amount))
}
