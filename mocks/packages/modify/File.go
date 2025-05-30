// Code generated by mockery. DO NOT EDIT.

package modify

import mock "github.com/stretchr/testify/mock"

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

type File_Expecter struct {
	mock *mock.Mock
}

func (_m *File) EXPECT() *File_Expecter {
	return &File_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with no fields
func (_m *File) Apply() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Apply")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type File_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
func (_e *File_Expecter) Apply() *File_Apply_Call {
	return &File_Apply_Call{Call: _e.mock.On("Apply")}
}

func (_c *File_Apply_Call) Run(run func()) *File_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *File_Apply_Call) Return(_a0 error) *File_Apply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *File_Apply_Call) RunAndReturn(run func() error) *File_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// NewFile creates a new instance of File. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *File {
	mock := &File{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
