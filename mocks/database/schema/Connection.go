// Code generated by mockery. DO NOT EDIT.

package schema

import mock "github.com/stretchr/testify/mock"

// Connection is an autogenerated mock type for the Connection type
type Connection struct {
	mock.Mock
}

type Connection_Expecter struct {
	mock *mock.Mock
}

func (_m *Connection) EXPECT() *Connection_Expecter {
	return &Connection_Expecter{mock: &_m.Mock}
}

// Connection provides a mock function with given fields:
func (_m *Connection) Connection() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connection")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Connection_Connection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connection'
type Connection_Connection_Call struct {
	*mock.Call
}

// Connection is a helper method to define mock.On call
func (_e *Connection_Expecter) Connection() *Connection_Connection_Call {
	return &Connection_Connection_Call{Call: _e.mock.On("Connection")}
}

func (_c *Connection_Connection_Call) Run(run func()) *Connection_Connection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Connection_Connection_Call) Return(_a0 string) *Connection_Connection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Connection_Connection_Call) RunAndReturn(run func() string) *Connection_Connection_Call {
	_c.Call.Return(run)
	return _c
}

// NewConnection creates a new instance of Connection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnection(t interface {
	mock.TestingT
	Cleanup(func())
}) *Connection {
	mock := &Connection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
