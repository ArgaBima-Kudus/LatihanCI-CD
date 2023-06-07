// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

type Logger_Expecter struct {
	mock *mock.Mock
}

func (_m *Logger) EXPECT() *Logger_Expecter {
	return &Logger_Expecter{mock: &_m.Mock}
}

// Log provides a mock function with given fields: msg
func (_m *Logger) Log(msg string) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Logger_Log_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Log'
type Logger_Log_Call struct {
	*mock.Call
}

// Log is a helper method to define mock.On call
//   - msg string
func (_e *Logger_Expecter) Log(msg interface{}) *Logger_Log_Call {
	return &Logger_Log_Call{Call: _e.mock.On("Log", msg)}
}

func (_c *Logger_Log_Call) Run(run func(msg string)) *Logger_Log_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Logger_Log_Call) Return(_a0 error) *Logger_Log_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Logger_Log_Call) RunAndReturn(run func(string) error) *Logger_Log_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogger(t mockConstructorTestingTNewLogger) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
