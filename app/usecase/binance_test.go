package usecase

import (
	"errors"
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

func TestBinanceUseCase_GetTradingPairs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := httptest.NewRecorder()
	ctx := GetTestGinContext(r)

	type args struct {
		ctx *gin.Context
	}

	tradingPairsRepo := mock_repository.NewMockTradingPairsRepository(ctrl)

	w := NewBinanceUsecase(tradingPairsRepo)

	tests := []struct {
		name       string
		args       args
		beforeTest func()
		want       string
		err        error
		wantErr    string
	}{

		{
			name: "successfully created binance-pairs file",
			args: args{
				ctx: ctx,
			},
			beforeTest: func() {
				tradingPairsRepo.EXPECT().GetExchangePair(constants.BinanceExchangePairURL).Return([]models.TradingPairs{}, nil)
				tradingPairsRepo.EXPECT().WriteToFile([]models.TradingPairs{}, "binance-pairs.txt").Return("./binance-pairs.txt", nil)
			},
			want: "./binance-pairs.txt",
			err:  nil,
		}, {
			name: "error in getting exchange pair",
			args: args{
				ctx: ctx,
			},
			beforeTest: func() {
				tradingPairsRepo.EXPECT().GetExchangePair(constants.BinanceExchangePairURL).Return([]models.TradingPairs{}, errors.New(""))
			},
			want: "",
			err:  errors.New("something unexpected happen. error: "),
		}, {
			name: "error in writing the file",
			args: args{
				ctx: ctx,
			},
			beforeTest: func() {
				tradingPairsRepo.EXPECT().GetExchangePair(constants.BinanceExchangePairURL).Return([]models.TradingPairs{}, nil)
				tradingPairsRepo.EXPECT().WriteToFile([]models.TradingPairs{}, "binance-pairs.txt").Return("", errors.New(""))
			},
			want: "",
			err:  errors.New("error in getting the file name. error: "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest()
			res, err := w.GetTradingPairs(tt.args.ctx)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Equal(t, tt.want, res)
			}
		})
	}
}
