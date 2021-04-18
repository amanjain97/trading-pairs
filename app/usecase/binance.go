package usecase

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"tradingpairs/domain/models"
)

const binanceExchangePairURL = "https://api.binance.com/api/v3/exchangeInfo"

type binanceUseCase struct {
}

type BinanceUsecase interface {
	GetTradingPairs(ctx echo.Context) (string, error)
}

func NewBinanceUsecase() BinanceUsecase {
	return &binanceUseCase{}
}

func (b binanceUseCase) GetTradingPairs(ctx echo.Context) (string, error) {
	url := binanceExchangePairURL
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

	tradingPairs := models.BinanceTradingPairs{}
	err = json.NewDecoder(res.Body).Decode(&tradingPairs)
	if err != nil {
		return "", err
	}

	pairs := make([]models.TradingPairs, 0)
	for _, tp := range tradingPairs.Symbols {
		pairs = append(pairs, models.TradingPairs{
			BaseCurrency:  tp.Symbol[:3],
			QuoteCurrency: tp.Symbol[3:],
		})
	}

	name, err := writeToFile(pairs, "binance-pairs.txt")
	if err != nil {
		return "", err
	}
	return name, nil
}
