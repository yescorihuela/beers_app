// Code generated by MockGen. DO NOT EDIT.
// Source: domain/beer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/yescorihuela/beers_app/domain"
	errs "github.com/yescorihuela/beers_app/errs"
	reflect "reflect"
)

// MockBeerRepository is a mock of BeerRepository interface
type MockBeerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBeerRepositoryMockRecorder
}

// MockBeerRepositoryMockRecorder is the mock recorder for MockBeerRepository
type MockBeerRepositoryMockRecorder struct {
	mock *MockBeerRepository
}

// NewMockBeerRepository creates a new mock instance
func NewMockBeerRepository(ctrl *gomock.Controller) *MockBeerRepository {
	mock := &MockBeerRepository{ctrl: ctrl}
	mock.recorder = &MockBeerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeerRepository) EXPECT() *MockBeerRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockBeerRepository) FindAll() ([]domain.Beer, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Beer)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockBeerRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockBeerRepository)(nil).FindAll))
}

// FindOne mocks base method
func (m *MockBeerRepository) FindOne(arg0 int) (*domain.Beer, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0)
	ret0, _ := ret[0].(*domain.Beer)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne
func (mr *MockBeerRepositoryMockRecorder) FindOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockBeerRepository)(nil).FindOne), arg0)
}

// FindIfExists mocks base method
func (m *MockBeerRepository) FindIfExists(arg0 uint) *errs.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIfExists", arg0)
	ret0, _ := ret[0].(*errs.AppError)
	return ret0
}

// FindIfExists indicates an expected call of FindIfExists
func (mr *MockBeerRepositoryMockRecorder) FindIfExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIfExists", reflect.TypeOf((*MockBeerRepository)(nil).FindIfExists), arg0)
}

// Create mocks base method
func (m *MockBeerRepository) Create(beer domain.Beer) (*domain.Beer, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", beer)
	ret0, _ := ret[0].(*domain.Beer)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBeerRepositoryMockRecorder) Create(beer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBeerRepository)(nil).Create), beer)
}
