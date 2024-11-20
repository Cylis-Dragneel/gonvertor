package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func api(currencyFrom string, currencyTo string, amount float64) float64 {
	// Getting an environment variable for secret
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal(
			"Must set the environment variable of the name API_KEY. Get the API Key from openexchangerates.org",
		)
	}
	url := "https://openexchangerates.org/api/latest.json?app_id=" + apiKey

	// This sends a GET request to the api
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// With the request we send a header with info needed from the api's documentation
	req.Header.Add("accept", "application/json")

	// We then send a http request to get a http response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Defer ensures that this line of code always executes when the function returns
	// this is used to close connections and prevent resource leakage
	// this line of code is best practice whenever doing an API Call
	defer resp.Body.Close()

	// We then create a struct to store all the fields from the JSON received from the api
	var Response struct {
		Disclaimer string
		License    string
		Timestamp  int64
		Base       string
		Rates      map[string]float64
	}

	// Next we Decode resp.Body, Body is a field in response that contains the data for json decoding
	// We then put the decoded json into the Response struct
	err = json.NewDecoder(resp.Body).Decode(&Response)
	if err != nil {
		log.Fatal("Error:", err)
	}

	if currencyFrom == "USD" {
		rateFrom := 1.0
		rateTo := Response.Rates[currencyTo]
		// We multiply the result by 100 to shift decimal places
		// We use math.Round to round up to the nearest integer
		// We divide the answer by 100 to move back the decimal places
		// We use float64 to ensure result is a float64
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
	// We use os.Args[number] to get an argument passed to the program
	base := os.Args[1]
	// We use strconv.ParseFloat to convert the argument which is a string type to a float64
	amount, err := strconv.ParseFloat(base, 64)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	currencyFrom := strings.ToUpper(os.Args[2])
	currencyTo := strings.ToUpper(os.Args[3])
	fmt.Println(base, currencyFrom, "=", api(currencyFrom, currencyTo, amount), currencyTo)
}
