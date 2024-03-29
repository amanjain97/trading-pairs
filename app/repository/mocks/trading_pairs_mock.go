// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/trading_pairs.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	models "tradingpairs/domain/models"

	gomock "github.com/golang/mock/gomock"
)

// MockTradingPairsRepository is a mock of TradingPairsRepository interface.
type MockTradingPairsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTradingPairsRepositoryMockRecorder
}

// MockTradingPairsRepositoryMockRecorder is the mock recorder for MockTradingPairsRepository.
type MockTradingPairsRepositoryMockRecorder struct {
	mock *MockTradingPairsRepository
}

// NewMockTradingPairsRepository creates a new mock instance.
func NewMockTradingPairsRepository(ctrl *gomock.Controller) *MockTradingPairsRepository {
	mock := &MockTradingPairsRepository{ctrl: ctrl}
	mock.recorder = &MockTradingPairsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTradingPairsRepository) EXPECT() *MockTradingPairsRepositoryMockRecorder {
	return m.recorder
}

// GetExchangePair mocks base method.
func (m *MockTradingPairsRepository) GetExchangePair(url string) ([]models.TradingPairs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangePair", url)
	ret0, _ := ret[0].([]models.TradingPairs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangePair indicates an expected call of GetExchangePair.
func (mr *MockTradingPairsRepositoryMockRecorder) GetExchangePair(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangePair", reflect.TypeOf((*MockTradingPairsRepository)(nil).GetExchangePair), url)
}

// WriteToFile mocks base method.
func (m *MockTradingPairsRepository) WriteToFile(pairs []models.TradingPairs, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToFile", pairs, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteToFile indicates an expected call of WriteToFile.
func (mr *MockTradingPairsRepositoryMockRecorder) WriteToFile(pairs, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToFile", reflect.TypeOf((*MockTradingPairsRepository)(nil).WriteToFile), pairs, filename)
}
