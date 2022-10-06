package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_repository "tradingpairs/app/repository/mocks"
	"tradingpairs/domain/models"
	"tradingpairs/utils/constants"
)

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func Test_GetTradingPairs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tradingPairsRepo := mock_repository.NewMockTradingPairsRepository(ctrl)

	w := NewBinanceUsecase(tradingPairsRepo)

	r := httptest.NewRecorder()
	ctx := GetTestGinContext(r)

	tradingPairsRepo.EXPECT().GetExchangePair(constants.BinanceExchangePairURL).Return([]models.TradingPairs{}, errors.New(""))

	_, err := w.GetTradingPairs(ctx)

	assert.EqualError(t, err, "something unexpected happen. error: ")
}

func Test_WriteFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tradingPairsRepo := mock_repository.NewMockTradingPairsRepository(ctrl)

	w := NewBinanceUsecase(tradingPairsRepo)

	r := httptest.NewRecorder()
	ctx := GetTestGinContext(r)
	tradingPairsRepo.EXPECT().GetExchangePair(constants.BinanceExchangePairURL).Return([]models.TradingPairs{}, nil)

	tradingPairsRepo.EXPECT().WriteToFile([]models.TradingPairs{}, "binance-pairs.txt").Return("", errors.New(""))

	_, err := w.GetTradingPairs(ctx)
	assert.EqualError(t, err, "error in getting the file name. error: ")
}

func Test_GetBinanceTradingPairs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tradingPairsRepo := mock_repository.NewMockTradingPairsRepository(ctrl)

	w := NewBinanceUsecase(tradingPairsRepo)

	r := httptest.NewRecorder()
	ctx := GetTestGinContext(r)
	tradingPairsRepo.EXPECT().GetExchangePair(constants.BinanceExchangePairURL).Return([]models.TradingPairs{}, nil)

	tradingPairsRepo.EXPECT().WriteToFile([]models.TradingPairs{}, "binance-pairs.txt").Return("./binance-pairs.txt", nil)

	res, err := w.GetTradingPairs(ctx)
	if err != nil {
		fmt.Println("got=", err)
		t.Errorf("tradingPairsRepository.GetExchangePair() error = %v, wantErr=%v", err, "something unexpected happen. error: h")
		return
	}

	assert.Equal(t, "./binance-pairs.txt", res)
}
