package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"

	"tradingpairs/app/repository"
	"tradingpairs/utils/constants"
)

type binanceUseCase struct {
	tradingPairsRepository repository.TradingPairsRepository
}

type BinanceUsecase interface {
	GetTradingPairs(ctx *gin.Context) (string, error)
}

func NewBinanceUsecase(tradingPairsRepository repository.TradingPairsRepository) BinanceUsecase {
	return &binanceUseCase{
		tradingPairsRepository: tradingPairsRepository,
	}
}

func (b binanceUseCase) GetTradingPairs(_ *gin.Context) (string, error) {
	url := constants.BinanceExchangePairURL

	pairs, err := b.tradingPairsRepository.GetExchangePair(url)
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("something unexpected happen. error: %v", err)
	}

	name, err := b.tradingPairsRepository.WriteToFile(pairs, "binance-pairs.txt")
	if err != nil {
		return "", fmt.Errorf("error in getting the file name. error: %v", err)
	}

	return name, nil
}
