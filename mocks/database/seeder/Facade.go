// Code generated by mockery. DO NOT EDIT.

package seeder

import (
	seeder "github.com/goravel/framework/contracts/database/seeder"
	mock "github.com/stretchr/testify/mock"
)

// Facade is an autogenerated mock type for the Facade type
type Facade struct {
	mock.Mock
}

type Facade_Expecter struct {
	mock *mock.Mock
}

func (_m *Facade) EXPECT() *Facade_Expecter {
	return &Facade_Expecter{mock: &_m.Mock}
}

// Call provides a mock function with given fields: seeders
func (_m *Facade) Call(seeders []seeder.Seeder) error {
	ret := _m.Called(seeders)

	if len(ret) == 0 {
		panic("no return value specified for Call")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]seeder.Seeder) error); ok {
		r0 = rf(seeders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Facade_Call_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Call'
type Facade_Call_Call struct {
	*mock.Call
}

// Call is a helper method to define mock.On call
//   - seeders []seeder.Seeder
func (_e *Facade_Expecter) Call(seeders interface{}) *Facade_Call_Call {
	return &Facade_Call_Call{Call: _e.mock.On("Call", seeders)}
}

func (_c *Facade_Call_Call) Run(run func(seeders []seeder.Seeder)) *Facade_Call_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]seeder.Seeder))
	})
	return _c
}

func (_c *Facade_Call_Call) Return(_a0 error) *Facade_Call_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Facade_Call_Call) RunAndReturn(run func([]seeder.Seeder) error) *Facade_Call_Call {
	_c.Call.Return(run)
	return _c
}

// CallOnce provides a mock function with given fields: seeders
func (_m *Facade) CallOnce(seeders []seeder.Seeder) error {
	ret := _m.Called(seeders)

	if len(ret) == 0 {
		panic("no return value specified for CallOnce")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]seeder.Seeder) error); ok {
		r0 = rf(seeders)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Facade_CallOnce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CallOnce'
type Facade_CallOnce_Call struct {
	*mock.Call
}

// CallOnce is a helper method to define mock.On call
//   - seeders []seeder.Seeder
func (_e *Facade_Expecter) CallOnce(seeders interface{}) *Facade_CallOnce_Call {
	return &Facade_CallOnce_Call{Call: _e.mock.On("CallOnce", seeders)}
}

func (_c *Facade_CallOnce_Call) Run(run func(seeders []seeder.Seeder)) *Facade_CallOnce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]seeder.Seeder))
	})
	return _c
}

func (_c *Facade_CallOnce_Call) Return(_a0 error) *Facade_CallOnce_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Facade_CallOnce_Call) RunAndReturn(run func([]seeder.Seeder) error) *Facade_CallOnce_Call {
	_c.Call.Return(run)
	return _c
}

// GetSeeder provides a mock function with given fields: name
func (_m *Facade) GetSeeder(name string) seeder.Seeder {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetSeeder")
	}

	var r0 seeder.Seeder
	if rf, ok := ret.Get(0).(func(string) seeder.Seeder); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(seeder.Seeder)
		}
	}

	return r0
}

// Facade_GetSeeder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSeeder'
type Facade_GetSeeder_Call struct {
	*mock.Call
}

// GetSeeder is a helper method to define mock.On call
//   - name string
func (_e *Facade_Expecter) GetSeeder(name interface{}) *Facade_GetSeeder_Call {
	return &Facade_GetSeeder_Call{Call: _e.mock.On("GetSeeder", name)}
}

func (_c *Facade_GetSeeder_Call) Run(run func(name string)) *Facade_GetSeeder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Facade_GetSeeder_Call) Return(_a0 seeder.Seeder) *Facade_GetSeeder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Facade_GetSeeder_Call) RunAndReturn(run func(string) seeder.Seeder) *Facade_GetSeeder_Call {
	_c.Call.Return(run)
	return _c
}

// GetSeeders provides a mock function with no fields
func (_m *Facade) GetSeeders() []seeder.Seeder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSeeders")
	}

	var r0 []seeder.Seeder
	if rf, ok := ret.Get(0).(func() []seeder.Seeder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]seeder.Seeder)
		}
	}

	return r0
}

// Facade_GetSeeders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSeeders'
type Facade_GetSeeders_Call struct {
	*mock.Call
}

// GetSeeders is a helper method to define mock.On call
func (_e *Facade_Expecter) GetSeeders() *Facade_GetSeeders_Call {
	return &Facade_GetSeeders_Call{Call: _e.mock.On("GetSeeders")}
}

func (_c *Facade_GetSeeders_Call) Run(run func()) *Facade_GetSeeders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Facade_GetSeeders_Call) Return(_a0 []seeder.Seeder) *Facade_GetSeeders_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Facade_GetSeeders_Call) RunAndReturn(run func() []seeder.Seeder) *Facade_GetSeeders_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: seeders
func (_m *Facade) Register(seeders []seeder.Seeder) {
	_m.Called(seeders)
}

// Facade_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Facade_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - seeders []seeder.Seeder
func (_e *Facade_Expecter) Register(seeders interface{}) *Facade_Register_Call {
	return &Facade_Register_Call{Call: _e.mock.On("Register", seeders)}
}

func (_c *Facade_Register_Call) Run(run func(seeders []seeder.Seeder)) *Facade_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]seeder.Seeder))
	})
	return _c
}

func (_c *Facade_Register_Call) Return() *Facade_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *Facade_Register_Call) RunAndReturn(run func([]seeder.Seeder)) *Facade_Register_Call {
	_c.Run(run)
	return _c
}

// NewFacade creates a new instance of Facade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *Facade {
	mock := &Facade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
