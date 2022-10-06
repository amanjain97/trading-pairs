package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"tradingpairs/app/usecase"
)

type tradingPairHandler struct {
	binanceUsecase  usecase.BinanceUsecase
	coinbaseUsecase usecase.CoinbaseUsecase
}

func NewTradingPairHandler(e *gin.RouterGroup, binance usecase.BinanceUsecase, coinbase usecase.CoinbaseUsecase) {
	handler := tradingPairHandler{
		binanceUsecase:  binance,
		coinbaseUsecase: coinbase,
	}
	e.GET("/binance/trading-pairs", handler.GetBinanceTradingPairs)
	e.GET("/coinbase/trading-pairs", handler.GetCoinbaseTradingPairs)
}

func (handler tradingPairHandler) GetBinanceTradingPairs(ctx *gin.Context) {
	filename, err := handler.binanceUsecase.GetTradingPairs(ctx)
	if filename == "" {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("unexpected error occurred. error: %v", err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, filename)
}

func (handler tradingPairHandler) GetCoinbaseTradingPairs(ctx *gin.Context) {
	filename, err := handler.coinbaseUsecase.GetTradingPairs(ctx)
	if filename == "" {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("unexpected error occurred. error: %v", err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, filename)
}
