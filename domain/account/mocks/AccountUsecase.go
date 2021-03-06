// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	account "github.com/madeindra/devoria-workshop-to-challenge/domain/account"

	mock "github.com/stretchr/testify/mock"

	response "github.com/madeindra/devoria-workshop-to-challenge/internal/response"
)

// AccountUsecase is an autogenerated mock type for the AccountUsecase type
type AccountUsecase struct {
	mock.Mock
}

// GetAccount provides a mock function with given fields: ctx, ID
func (_m *AccountUsecase) GetAccount(ctx context.Context, ID int64) response.Response {
	ret := _m.Called(ctx, ID)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64) response.Response); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

// Login provides a mock function with given fields: ctx, params
func (_m *AccountUsecase) Login(ctx context.Context, params account.AccountLoginRequest) response.Response {
	ret := _m.Called(ctx, params)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, account.AccountLoginRequest) response.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

// Register provides a mock function with given fields: ctx, params
func (_m *AccountUsecase) Register(ctx context.Context, params account.AccountRegisterRequest) response.Response {
	ret := _m.Called(ctx, params)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, account.AccountRegisterRequest) response.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}
