package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"tradingpairs/app/repository"
	"tradingpairs/utils/constants"
)

type coinbaseUseCase struct {
	tradingPairsRepository repository.TradingPairsRepository
}

type CoinbaseUsecase interface {
	GetTradingPairs(ctx *gin.Context) (string, error)
}

func NewCoinbaseUsecase(tradingPairsRepository repository.TradingPairsRepository) CoinbaseUsecase {
	return &coinbaseUseCase{
		tradingPairsRepository: tradingPairsRepository,
	}
}

func (b coinbaseUseCase) GetTradingPairs(_ *gin.Context) (string, error) {
	url := constants.CoinbaseExchangePairURL

	pairs, err := b.tradingPairsRepository.GetExchangePair(url)
	if err != nil {
		return "", fmt.Errorf("something unexpected happen. error: %v", err)
	}

	name, err := b.tradingPairsRepository.WriteToFile(pairs, "coinbase-pairs.txt")
	if err != nil {
		return "", fmt.Errorf("error in getting the file name. error: %v", err)
	}

	return name, nil
}
