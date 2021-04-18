package usecase

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"tradingpairs/domain/models"
)

const coinbaseExchangePairURL = "https://api-public.sandbox.pro.coinbase.com/products"

type coinbaseUseCase struct {
}

type CoinbaseUsecase interface {
	GetTradingPairs(ctx echo.Context) (string, error)
}

func NewCoinbaseUsecase() CoinbaseUsecase {
	return &coinbaseUseCase{}
}

func (b coinbaseUseCase) GetTradingPairs(ctx echo.Context) (string, error) {
	url := coinbaseExchangePairURL
	method := "GET"

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	tradingPairs := []models.CoinbaseTradingPair{}
	err = json.NewDecoder(res.Body).Decode(&tradingPairs)
	if err != nil {
		return "", err
	}

	pairs := make([]models.TradingPairs, 0)
	for _, tp := range tradingPairs {
		pairs = append(pairs, models.TradingPairs{
			BaseCurrency:  tp.BaseCurrency,
			QuoteCurrency: tp.QuoteCurrency,
		})
	}

	name, err := writeToFile(pairs, "coinbase-pairs.txt")
	if err != nil {
		return "", err
	}
	return name, nil
}
