// Code generated by mockery. DO NOT EDIT.

package orm

import mock "github.com/stretchr/testify/mock"

// Association is an autogenerated mock type for the Association type
type Association struct {
	mock.Mock
}

type Association_Expecter struct {
	mock *mock.Mock
}

func (_m *Association) EXPECT() *Association_Expecter {
	return &Association_Expecter{mock: &_m.Mock}
}

// Append provides a mock function with given fields: values
func (_m *Association) Append(values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Append")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Association_Append_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Append'
type Association_Append_Call struct {
	*mock.Call
}

// Append is a helper method to define mock.On call
//   - values ...interface{}
func (_e *Association_Expecter) Append(values ...interface{}) *Association_Append_Call {
	return &Association_Append_Call{Call: _e.mock.On("Append",
		append([]interface{}{}, values...)...)}
}

func (_c *Association_Append_Call) Run(run func(values ...interface{})) *Association_Append_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Association_Append_Call) Return(_a0 error) *Association_Append_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Association_Append_Call) RunAndReturn(run func(...interface{}) error) *Association_Append_Call {
	_c.Call.Return(run)
	return _c
}

// Clear provides a mock function with no fields
func (_m *Association) Clear() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Association_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type Association_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
func (_e *Association_Expecter) Clear() *Association_Clear_Call {
	return &Association_Clear_Call{Call: _e.mock.On("Clear")}
}

func (_c *Association_Clear_Call) Run(run func()) *Association_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Association_Clear_Call) Return(_a0 error) *Association_Clear_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Association_Clear_Call) RunAndReturn(run func() error) *Association_Clear_Call {
	_c.Call.Return(run)
	return _c
}

// Count provides a mock function with no fields
func (_m *Association) Count() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Association_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type Association_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
func (_e *Association_Expecter) Count() *Association_Count_Call {
	return &Association_Count_Call{Call: _e.mock.On("Count")}
}

func (_c *Association_Count_Call) Run(run func()) *Association_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Association_Count_Call) Return(_a0 int64) *Association_Count_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Association_Count_Call) RunAndReturn(run func() int64) *Association_Count_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: values
func (_m *Association) Delete(values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Association_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Association_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - values ...interface{}
func (_e *Association_Expecter) Delete(values ...interface{}) *Association_Delete_Call {
	return &Association_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{}, values...)...)}
}

func (_c *Association_Delete_Call) Run(run func(values ...interface{})) *Association_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Association_Delete_Call) Return(_a0 error) *Association_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Association_Delete_Call) RunAndReturn(run func(...interface{}) error) *Association_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: out, conds
func (_m *Association) Find(out interface{}, conds ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, out)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(out, conds...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Association_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type Association_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - out interface{}
//   - conds ...interface{}
func (_e *Association_Expecter) Find(out interface{}, conds ...interface{}) *Association_Find_Call {
	return &Association_Find_Call{Call: _e.mock.On("Find",
		append([]interface{}{out}, conds...)...)}
}

func (_c *Association_Find_Call) Run(run func(out interface{}, conds ...interface{})) *Association_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Association_Find_Call) Return(_a0 error) *Association_Find_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Association_Find_Call) RunAndReturn(run func(interface{}, ...interface{}) error) *Association_Find_Call {
	_c.Call.Return(run)
	return _c
}

// Replace provides a mock function with given fields: values
func (_m *Association) Replace(values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Replace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Association_Replace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Replace'
type Association_Replace_Call struct {
	*mock.Call
}

// Replace is a helper method to define mock.On call
//   - values ...interface{}
func (_e *Association_Expecter) Replace(values ...interface{}) *Association_Replace_Call {
	return &Association_Replace_Call{Call: _e.mock.On("Replace",
		append([]interface{}{}, values...)...)}
}

func (_c *Association_Replace_Call) Run(run func(values ...interface{})) *Association_Replace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Association_Replace_Call) Return(_a0 error) *Association_Replace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Association_Replace_Call) RunAndReturn(run func(...interface{}) error) *Association_Replace_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssociation creates a new instance of Association. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssociation(t interface {
	mock.TestingT
	Cleanup(func())
}) *Association {
	mock := &Association{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
