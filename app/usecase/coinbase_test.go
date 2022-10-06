package usecase

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_repository "tradingpairs/app/repository/mocks"
	"tradingpairs/domain/models"
	"tradingpairs/utils/constants"
)

func TestCoinbaseUseCase_GetTradingPairs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := httptest.NewRecorder()
	ctx := GetTestGinContext(r)

	type args struct {
		ctx *gin.Context
	}

	tradingPairsRepo := mock_repository.NewMockTradingPairsRepository(ctrl)

	w := NewCoinbaseUsecase(tradingPairsRepo)

	tests := []struct {
		name       string
		args       args
		beforeTest func()
		want       string
		err        error
		wantErr    string
	}{

		{
			name: "successfully created coinbase-pairs file",
			args: args{
				ctx: ctx,
			},
			beforeTest: func() {
				tradingPairsRepo.EXPECT().GetExchangePair(constants.CoinbaseExchangePairURL).Return([]models.TradingPairs{}, nil)
				tradingPairsRepo.EXPECT().WriteToFile([]models.TradingPairs{}, "coinbase-pairs.txt").Return("./coinbase-pairs.txt", nil)
			},
			want: "./coinbase-pairs.txt",
			err:  nil,
		}, {
			name: "error in getting exchange pair",
			args: args{
				ctx: ctx,
			},
			beforeTest: func() {
				tradingPairsRepo.EXPECT().GetExchangePair(constants.CoinbaseExchangePairURL).Return([]models.TradingPairs{}, errors.New(""))
			},
			want: "",
			err:  errors.New("something unexpected happen. error: "),
		}, {
			name: "error in writing the file",
			args: args{
				ctx: ctx,
			},
			beforeTest: func() {
				tradingPairsRepo.EXPECT().GetExchangePair(constants.CoinbaseExchangePairURL).Return([]models.TradingPairs{}, nil)
				tradingPairsRepo.EXPECT().WriteToFile([]models.TradingPairs{}, "coinbase-pairs.txt").Return("", errors.New(""))
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
