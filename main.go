package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"tradingpairs/app/interface/http"
	"tradingpairs/app/usecase"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1 := e.Group("/api/v1")

	binanceUsecase := usecase.NewBinanceUsecase()
	coinbaseUsecase := usecase.NewCoinbaseUsecase()
	http.NewTradingPairHandler(v1, binanceUsecase, coinbaseUsecase)

	// Should use a separate config file
	e.Logger.Debug(e.Start(":5000"))
}
