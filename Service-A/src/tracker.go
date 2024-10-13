package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type BitcoinPrice struct {
	Bitcoin map[string]float64 `json:"bitcoin"`
}

type BitcoinPriceTracker struct {
	Prices      []float64
	MaxHistory  int
	MinuteCount int
	LastPrice   float64
	LastAverage float64
	Currency    string
	BaseURL     string
}

func NewBitcoinPriceTracker() *BitcoinPriceTracker {
	maxHistory := viper.GetInt("bitcoin.max_history")
	currency := viper.GetString("bitcoin.currency")
	baseURL := viper.GetString("bitcoin.base_url")
	fetchIntervalMin := viper.GetInt("bitcoin.fetch_interval_minutes")

	tracker := &BitcoinPriceTracker{
		MaxHistory: maxHistory,
		Prices:     make([]float64, 0),
		Currency:   currency,
		BaseURL:    baseURL,
	}

	go tracker.startPriceFetcher(fetchIntervalMin)
	return tracker
}

func (b *BitcoinPriceTracker) startPriceFetcher(intervalMin int) {
	b.fetchBitcoinPrice()
	ticker := time.NewTicker(time.Duration(intervalMin) * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		b.fetchBitcoinPrice()
	}
}

func (b *BitcoinPriceTracker) fetchBitcoinPrice() {
	apiURL := fmt.Sprintf("%s?ids=bitcoin&vs_currencies=%s", b.BaseURL, b.Currency)
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error fetching Bitcoin price: %v", err)
		return
	}
	defer resp.Body.Close()

	var result BitcoinPrice
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding Bitcoin price: %v", err)
		return
	}

	price := result.Bitcoin[b.Currency]
	b.LastPrice = price
	log.Printf("Current Bitcoin Price in %s: %.2f", b.Currency, price)

	b.Prices = append(b.Prices, price)
	if len(b.Prices) > b.MaxHistory {
		b.Prices = b.Prices[1:]
	}

	b.MinuteCount++
	if b.MinuteCount == b.MaxHistory {
		b.calculateAveragePrice()
		b.MinuteCount = 0
	}
}

func (b *BitcoinPriceTracker) calculateAveragePrice() {
	sum := 0.0
	for _, price := range b.Prices {
		sum += price
	}
	if len(b.Prices) > 0 {
		b.LastAverage = sum / float64(len(b.Prices))
		log.Printf("Average Bitcoin Price (Last %d Minutes): %.2f", b.MaxHistory, b.LastAverage)
	}
}

func (b *BitcoinPriceTracker) GetLastPrice() float64 {
	return b.LastPrice
}

func (b *BitcoinPriceTracker) GetLastAverage() float64 {
	if len(b.Prices) == 0 {
		return math.NaN()
	}
	return b.LastAverage
}

func (b *BitcoinPriceTracker) GetCurrency() string {
	return b.Currency
}
