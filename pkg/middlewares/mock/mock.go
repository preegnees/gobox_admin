// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock_middlewares is a generated GoMock package.
package mock_middlewares

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockIMiddleware is a mock of IMiddleware interface.
type MockIMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockIMiddlewareMockRecorder
}

// MockIMiddlewareMockRecorder is the mock recorder for MockIMiddleware.
type MockIMiddlewareMockRecorder struct {
	mock *MockIMiddleware
}

// NewMockIMiddleware creates a new mock instance.
func NewMockIMiddleware(ctrl *gomock.Controller) *MockIMiddleware {
	mock := &MockIMiddleware{ctrl: ctrl}
	mock.recorder = &MockIMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMiddleware) EXPECT() *MockIMiddlewareMockRecorder {
	return m.recorder
}

// CheckJWT mocks base method.
func (m *MockIMiddleware) CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckJWT", next)
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// CheckJWT indicates an expected call of CheckJWT.
func (mr *MockIMiddlewareMockRecorder) CheckJWT(next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckJWT", reflect.TypeOf((*MockIMiddleware)(nil).CheckJWT), next)
}
