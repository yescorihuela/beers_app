// Code generated by MockGen. DO NOT EDIT.
// Source: services/currencyService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	errs "github.com/yescorihuela/beers_app/errs"
	reflect "reflect"
)

// MockCurrencyService is a mock of CurrencyService interface
type MockCurrencyService struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyServiceMockRecorder
}

// MockCurrencyServiceMockRecorder is the mock recorder for MockCurrencyService
type MockCurrencyServiceMockRecorder struct {
	mock *MockCurrencyService
}

// NewMockCurrencyService creates a new mock instance
func NewMockCurrencyService(ctrl *gomock.Controller) *MockCurrencyService {
	mock := &MockCurrencyService{ctrl: ctrl}
	mock.recorder = &MockCurrencyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCurrencyService) EXPECT() *MockCurrencyServiceMockRecorder {
	return m.recorder
}

// ConvertPrice mocks base method
func (m *MockCurrencyService) ConvertPrice(fromCurrency, toCurrency string, value float32) (*float32, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertPrice", fromCurrency, toCurrency, value)
	ret0, _ := ret[0].(*float32)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// ConvertPrice indicates an expected call of ConvertPrice
func (mr *MockCurrencyServiceMockRecorder) ConvertPrice(fromCurrency, toCurrency, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertPrice", reflect.TypeOf((*MockCurrencyService)(nil).ConvertPrice), fromCurrency, toCurrency, value)
}
