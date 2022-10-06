package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"tradingpairs/domain/models"
	"tradingpairs/utils/constants"
)

var crPool = sync.Pool{
	New: func() interface{} { return new([]models.CoinbaseTradingPair) },
}

var prPool = sync.Pool{
	New: func() interface{} { return new(models.BinanceTradingPairs) },
}

type tradingPairsRepository struct{}

type TradingPairsRepository interface {
	GetExchangePair(url string) ([]models.TradingPairs, error)
	WriteToFile(pairs []models.TradingPairs, filename string) (string, error)
}

func NewTradingPairsRepository() TradingPairsRepository {
	return &tradingPairsRepository{}
}

func (repo *tradingPairsRepository) GetExchangePair(url string) ([]models.TradingPairs, error) {
	pairs := make([]models.TradingPairs, 0)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("can't proceed with the request at this time. error: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("something unexpected happen. error: %v", err)
	}
	defer res.Body.Close()

	switch url {
	case constants.CoinbaseExchangePairURL:
		tradingPairs := crPool.Get().(*[]models.CoinbaseTradingPair)
		defer crPool.Put(tradingPairs)

		err = json.NewDecoder(res.Body).Decode(&tradingPairs)
		if err != nil {
			return nil, fmt.Errorf("error decoding the body. error: %v", err)
		}

		for _, tp := range *tradingPairs {
			pairs = append(pairs, models.TradingPairs{
				BaseCurrency:  tp.BaseCurrency,
				QuoteCurrency: tp.QuoteCurrency,
			})
		}
	case constants.BinanceExchangePairURL:
		tradingPairs := prPool.Get().(*models.BinanceTradingPairs)
		defer prPool.Put(tradingPairs)

		err = json.NewDecoder(res.Body).Decode(&tradingPairs)
		if err != nil {
			return nil, fmt.Errorf("error decoding the body. error: %v", err)
		}

		for _, tp := range tradingPairs.Symbols {
			pairs = append(pairs, models.TradingPairs{
				BaseCurrency:  tp.Symbol[:3],
				QuoteCurrency: tp.Symbol[3:],
			})
		}
	default:
		return nil, fmt.Errorf("exchange pair url not supported")
	}

	return pairs, nil
}

func (repo *tradingPairsRepository) WriteToFile(pairs []models.TradingPairs, filename string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting the current working directory. error: %v", err)
	}

	var f *os.File
	defer f.Close()

	filename = currentDir + "/" + filename

	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		err := os.MkdirAll(currentDir, 0700)
		if err != nil {
			return "", fmt.Errorf("error in creating the new directory. error: %v", err)
		}
	}

	f, err = os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("can't create file at this time. error: %v", err)
	}

	w := bufio.NewWriter(f)
	for _, p := range pairs {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("%s/%s", p.BaseCurrency, p.QuoteCurrency))
	}
	if err = w.Flush(); err != nil {
		return "", fmt.Errorf("error in writing a file. error: %v", err)
	}

	return filename, nil
}
