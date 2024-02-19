// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/cart.go

// Package mockUser is a generated GoMock package.
package mockUser

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	response "main.go/pkg/common/response"
)

// MockCartUsecase is a mock of CartUsecase interface.
type MockCartUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCartUsecaseMockRecorder
}

// MockCartUsecaseMockRecorder is the mock recorder for MockCartUsecase.
type MockCartUsecaseMockRecorder struct {
	mock *MockCartUsecase
}

// NewMockCartUsecase creates a new mock instance.
func NewMockCartUsecase(ctrl *gomock.Controller) *MockCartUsecase {
	mock := &MockCartUsecase{ctrl: ctrl}
	mock.recorder = &MockCartUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartUsecase) EXPECT() *MockCartUsecaseMockRecorder {
	return m.recorder
}

// AddToCart mocks base method.
func (m *MockCartUsecase) AddToCart(productId, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", productId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCart indicates an expected call of AddToCart.
func (mr *MockCartUsecaseMockRecorder) AddToCart(productId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockCartUsecase)(nil).AddToCart), productId, userId)
}

// CreateCart mocks base method.
func (m *MockCartUsecase) CreateCart(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockCartUsecaseMockRecorder) CreateCart(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockCartUsecase)(nil).CreateCart), id)
}

// ListCart mocks base method.
func (m *MockCartUsecase) ListCart(userId int) (response.ViewCart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCart", userId)
	ret0, _ := ret[0].(response.ViewCart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCart indicates an expected call of ListCart.
func (mr *MockCartUsecaseMockRecorder) ListCart(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCart", reflect.TypeOf((*MockCartUsecase)(nil).ListCart), userId)
}

// RemoveFromCart mocks base method.
func (m *MockCartUsecase) RemoveFromCart(userId, productId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", userId, productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockCartUsecaseMockRecorder) RemoveFromCart(userId, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockCartUsecase)(nil).RemoveFromCart), userId, productId)
}