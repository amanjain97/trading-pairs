package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"tradingpairs/app/usecase"
)

type tradingPairHandler struct {
	binanceUsecase  usecase.BinanceUsecase
	coinbaseUsecase usecase.CoinbaseUsecase
}

func NewTradingPairHandler(e *echo.Group, binance usecase.BinanceUsecase, coinbase usecase.CoinbaseUsecase) {
	handler := tradingPairHandler{
		binanceUsecase:  binance,
		coinbaseUsecase: coinbase,
	}
	e.GET("/binance/trading-pairs", handler.GetBinanceTradingPairs)
	e.GET("/coinbase/trading-pairs", handler.GetCoinbaseTradingPairs)
}

func (handler tradingPairHandler) GetBinanceTradingPairs(ctx echo.Context) error {
	filename, err := handler.binanceUsecase.GetTradingPairs(ctx)
	if filename == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unexpected error occurred. error: %v", err))
	}
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, filename)
}

func (handler tradingPairHandler) GetCoinbaseTradingPairs(ctx echo.Context) error {
	filename, err := handler.coinbaseUsecase.GetTradingPairs(ctx)
	if filename == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unexpected error occurred. error: %v", err))
	}
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, filename)
}
