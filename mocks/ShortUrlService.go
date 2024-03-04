// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ShortUrlService is an autogenerated mock type for the ShortUrlService type
type ShortUrlService struct {
	mock.Mock
}

type ShortUrlService_Expecter struct {
	mock *mock.Mock
}

func (_m *ShortUrlService) EXPECT() *ShortUrlService_Expecter {
	return &ShortUrlService_Expecter{mock: &_m.Mock}
}

// Reverse provides a mock function with given fields: ctx, shortLink
func (_m *ShortUrlService) Reverse(ctx context.Context, shortLink string) (string, error) {
	ret := _m.Called(ctx, shortLink)

	if len(ret) == 0 {
		panic("no return value specified for Reverse")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, shortLink)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortLink)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortLink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShortUrlService_Reverse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reverse'
type ShortUrlService_Reverse_Call struct {
	*mock.Call
}

// Reverse is a helper method to define mock.On call
//   - ctx context.Context
//   - shortLink string
func (_e *ShortUrlService_Expecter) Reverse(ctx interface{}, shortLink interface{}) *ShortUrlService_Reverse_Call {
	return &ShortUrlService_Reverse_Call{Call: _e.mock.On("Reverse", ctx, shortLink)}
}

func (_c *ShortUrlService_Reverse_Call) Run(run func(ctx context.Context, shortLink string)) *ShortUrlService_Reverse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ShortUrlService_Reverse_Call) Return(_a0 string, _a1 error) *ShortUrlService_Reverse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ShortUrlService_Reverse_Call) RunAndReturn(run func(context.Context, string) (string, error)) *ShortUrlService_Reverse_Call {
	_c.Call.Return(run)
	return _c
}

// Shorten provides a mock function with given fields: ctx, someString
func (_m *ShortUrlService) Shorten(ctx context.Context, someString string) (string, error) {
	ret := _m.Called(ctx, someString)

	if len(ret) == 0 {
		panic("no return value specified for Shorten")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, someString)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, someString)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, someString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShortUrlService_Shorten_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shorten'
type ShortUrlService_Shorten_Call struct {
	*mock.Call
}

// Shorten is a helper method to define mock.On call
//   - ctx context.Context
//   - someString string
func (_e *ShortUrlService_Expecter) Shorten(ctx interface{}, someString interface{}) *ShortUrlService_Shorten_Call {
	return &ShortUrlService_Shorten_Call{Call: _e.mock.On("Shorten", ctx, someString)}
}

func (_c *ShortUrlService_Shorten_Call) Run(run func(ctx context.Context, someString string)) *ShortUrlService_Shorten_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ShortUrlService_Shorten_Call) Return(_a0 string, _a1 error) *ShortUrlService_Shorten_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ShortUrlService_Shorten_Call) RunAndReturn(run func(context.Context, string) (string, error)) *ShortUrlService_Shorten_Call {
	_c.Call.Return(run)
	return _c
}

// NewShortUrlService creates a new instance of ShortUrlService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewShortUrlService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ShortUrlService {
	mock := &ShortUrlService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
