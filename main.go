package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"

	"tradingpairs/app/interface/httphandler"
	"tradingpairs/app/repository"
	"tradingpairs/app/usecase"
	logger "tradingpairs/utils/log"
)

func main() {
	//Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.New()

	newLogger, err := logger.NewLogger("debug")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] Failed to setup newLogger: %s\n", err)
	}
	defer func() {
		_ = newLogger.Sync()
	}()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(newLogger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(newLogger, true))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1 := r.Group("/api/v1")

	tradingPairsRepo := repository.NewTradingPairsRepository()
	binanceUsecase := usecase.NewBinanceUsecase(tradingPairsRepo)
	coinbaseUsecase := usecase.NewCoinbaseUsecase(tradingPairsRepo)
	httphandler.NewTradingPairHandler(v1, binanceUsecase, coinbaseUsecase)

	log.Println("Server listening on port 5000")
	srv := &http.Server{
		Addr:    ":5000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
