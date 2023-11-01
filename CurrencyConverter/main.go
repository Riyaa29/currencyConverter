package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
	"os"
	"strconv"
)

const APIEndpoint = "https://open.er-api.com/v6/latest/"

type ExchangeRatesResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func getExchangeRates(base string) (ExchangeRatesResponse, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return ExchangeRatesResponse{}, fmt.Errorf("API key not set")
	}
	reqURL := fmt.Sprintf("%s%s?apikey=%s", APIEndpoint, base, apiKey)
	resp, err := http.Get(reqURL)
	if err != nil {
		return ExchangeRatesResponse{}, err
	}
	defer resp.Body.Close()

	var rates ExchangeRatesResponse
	err = json.NewDecoder(resp.Body).Decode(&rates)
	return rates, err
}

func convertCurrency(c *gin.Context) {
	fromCurrency := c.DefaultQuery("from", "USD")
	toCurrency := c.DefaultQuery("to", "EUR")
	amountStr := c.DefaultQuery("amount", "1")

	amount, _ := strconv.ParseFloat(amountStr, 64)

	rates, err := getExchangeRates(fromCurrency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rate, found := rates.Rates[toCurrency]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
		return
	}

	convertedAmount := amount * rate
	c.JSON(http.StatusOK, gin.H{
		"from":   fromCurrency,
		"to":     toCurrency,
		"amount": amount,
		"result": convertedAmount,
	})
}
func getCurrencies(c *gin.Context) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not set"})
		return
	}
	reqURL := fmt.Sprintf("%s%s?apikey=%s", APIEndpoint, "USD", apiKey) // You can use any base currency to get the list
	resp, err := http.Get(reqURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var rates ExchangeRatesResponse
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currencies := make([]string, 0, len(rates.Rates))
	for currency := range rates.Rates {
		currencies = append(currencies, currency)
	}

	c.JSON(http.StatusOK, gin.H{
		"currencies": currencies,
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Setup CORS
	r.Use(cors.AllowAll())
	r.GET("/currencies", getCurrencies)
	r.GET("/convert", convertCurrency)
	r.Static("/static", "C:/Users/riyan/currency-converter")

	r.Run(":8081")

}
