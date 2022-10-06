package repository

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"tradingpairs/app/usecase"
)

func Test_GetTradingPairs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//c := NewMockTradingPairsRepository(ctrl)
	//NewHandler(c)
	//fmt.Println(c)
	////
	//type args struct {
	//	ctx *gin.Context
	//}
	//
	//tests := []struct {
	//	name       string
	//	args       args
	//	beforeTest func(tradingPairRepo *MockTradingPairsRepository)
	//	want       string
	//	wantErr    bool
	//}
	//	{
	//		name: "get exchange pair",
	//		args: args{
	//			ctx: context.Background(),
	//		},
	//		beforeTest: func(tradingPairRepo *MockTradingPairsRepository) {
	//			tradingPairRepo.EXPECT().GetExchangePair(constants.CoinbaseExchangePairURL).Return([]models.TradingPairs{}, nil)
	//		},
	//		want: []models.TradingPairs{},
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	tradingPairsRepo := NewMockTradingPairsRepository(ctrl)

	w := usecase.NewBinanceUsecase(tradingPairsRepo)
	//if tt.beforeTest != nil {
	//	tt.beforeTest(tradingPairsRepo)
	//}
	u := httptest.NewRecorder()
	ctx := GetTestGinContext(u)

	got, err := w.GetTradingPairs(ctx)
	fmt.Println("got=", got)
	if (err != nil) != true {
		t.Errorf("registerUserUseCase.GetExchangePair() error = %v, wantErr %v", err, "hello")
		return
	}
	if !reflect.DeepEqual(got, "hello") {
		t.Errorf("registerUserUseCase.GetExchangePair() = %v, want %v", got, "hello")
	}
	//	})
	//}
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
