package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	bitcoinTracker := NewBitcoinPriceTracker()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lastPrice := bitcoinTracker.GetLastPrice()
		lastAverage := bitcoinTracker.GetLastAverage()

		response := map[string]interface{}{
			"lastPrice": fmt.Sprintf("%.2f", lastPrice),
			"currency":  bitcoinTracker.GetCurrency(),
		}

		if !math.IsNaN(lastAverage) {
			response["lastAverage"] = fmt.Sprintf("%.2f", lastAverage)
		} else {
			response["lastAverage"] = "Average price data not available yet"
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	log.Println("service-a is running !")
	log.Fatal(http.ListenAndServe(":80", nil))
}
