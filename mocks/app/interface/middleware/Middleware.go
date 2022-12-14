// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// HitExternalAPI provides a mock function with given fields: ctx
func (_m *Middleware) HitExternalAPI(ctx *gin.Context) {
	_m.Called(ctx)
}

// Join provides a mock function with given fields: ctx
func (_m *Middleware) Join(ctx *gin.Context) {
	_m.Called(ctx)
}

// M32 provides a mock function with given fields: _a0
func (_m *Middleware) M32(_a0 *gin.Context) {
	_m.Called(_a0)
}

// M37 provides a mock function with given fields: _a0
func (_m *Middleware) M37(_a0 *gin.Context) {
	_m.Called(_a0)
}

// SampleReadValidation provides a mock function with given fields: _a0
func (_m *Middleware) SampleReadValidation(_a0 *gin.Context) {
	_m.Called(_a0)
}
