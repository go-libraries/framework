// Code generated by mockery. DO NOT EDIT.

package route

import (
	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	route "github.com/goravel/framework/contracts/route"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

type Router_Expecter struct {
	mock *mock.Mock
}

func (_m *Router) EXPECT() *Router_Expecter {
	return &Router_Expecter{mock: &_m.Mock}
}

// Any provides a mock function with given fields: relativePath, handler
func (_m *Router) Any(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Any_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Any'
type Router_Any_Call struct {
	*mock.Call
}

// Any is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Any(relativePath interface{}, handler interface{}) *Router_Any_Call {
	return &Router_Any_Call{Call: _e.mock.On("Any", relativePath, handler)}
}

func (_c *Router_Any_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Any_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Any_Call) Return() *Router_Any_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Any_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Any_Call {
	_c.Run(run)
	return _c
}

// Delete provides a mock function with given fields: relativePath, handler
func (_m *Router) Delete(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Router_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Delete(relativePath interface{}, handler interface{}) *Router_Delete_Call {
	return &Router_Delete_Call{Call: _e.mock.On("Delete", relativePath, handler)}
}

func (_c *Router_Delete_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Delete_Call) Return() *Router_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Delete_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Delete_Call {
	_c.Run(run)
	return _c
}

// Get provides a mock function with given fields: relativePath, handler
func (_m *Router) Get(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Router_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Get(relativePath interface{}, handler interface{}) *Router_Get_Call {
	return &Router_Get_Call{Call: _e.mock.On("Get", relativePath, handler)}
}

func (_c *Router_Get_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Get_Call) Return() *Router_Get_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Get_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Get_Call {
	_c.Run(run)
	return _c
}

// Group provides a mock function with given fields: handler
func (_m *Router) Group(handler route.GroupFunc) {
	_m.Called(handler)
}

// Router_Group_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Group'
type Router_Group_Call struct {
	*mock.Call
}

// Group is a helper method to define mock.On call
//   - handler route.GroupFunc
func (_e *Router_Expecter) Group(handler interface{}) *Router_Group_Call {
	return &Router_Group_Call{Call: _e.mock.On("Group", handler)}
}

func (_c *Router_Group_Call) Run(run func(handler route.GroupFunc)) *Router_Group_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(route.GroupFunc))
	})
	return _c
}

func (_c *Router_Group_Call) Return() *Router_Group_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Group_Call) RunAndReturn(run func(route.GroupFunc)) *Router_Group_Call {
	_c.Run(run)
	return _c
}

// Middleware provides a mock function with given fields: middlewares
func (_m *Router) Middleware(middlewares ...http.Middleware) route.Router {
	_va := make([]interface{}, len(middlewares))
	for _i := range middlewares {
		_va[_i] = middlewares[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Middleware")
	}

	var r0 route.Router
	if rf, ok := ret.Get(0).(func(...http.Middleware) route.Router); ok {
		r0 = rf(middlewares...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Router)
		}
	}

	return r0
}

// Router_Middleware_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Middleware'
type Router_Middleware_Call struct {
	*mock.Call
}

// Middleware is a helper method to define mock.On call
//   - middlewares ...http.Middleware
func (_e *Router_Expecter) Middleware(middlewares ...interface{}) *Router_Middleware_Call {
	return &Router_Middleware_Call{Call: _e.mock.On("Middleware",
		append([]interface{}{}, middlewares...)...)}
}

func (_c *Router_Middleware_Call) Run(run func(middlewares ...http.Middleware)) *Router_Middleware_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]http.Middleware, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(http.Middleware)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Router_Middleware_Call) Return(_a0 route.Router) *Router_Middleware_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Router_Middleware_Call) RunAndReturn(run func(...http.Middleware) route.Router) *Router_Middleware_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: relativePath, handler
func (_m *Router) Options(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type Router_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Options(relativePath interface{}, handler interface{}) *Router_Options_Call {
	return &Router_Options_Call{Call: _e.mock.On("Options", relativePath, handler)}
}

func (_c *Router_Options_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Options_Call) Return() *Router_Options_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Options_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Options_Call {
	_c.Run(run)
	return _c
}

// Patch provides a mock function with given fields: relativePath, handler
func (_m *Router) Patch(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type Router_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Patch(relativePath interface{}, handler interface{}) *Router_Patch_Call {
	return &Router_Patch_Call{Call: _e.mock.On("Patch", relativePath, handler)}
}

func (_c *Router_Patch_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Patch_Call) Return() *Router_Patch_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Patch_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Patch_Call {
	_c.Run(run)
	return _c
}

// Post provides a mock function with given fields: relativePath, handler
func (_m *Router) Post(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type Router_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Post(relativePath interface{}, handler interface{}) *Router_Post_Call {
	return &Router_Post_Call{Call: _e.mock.On("Post", relativePath, handler)}
}

func (_c *Router_Post_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Post_Call) Return() *Router_Post_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Post_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Post_Call {
	_c.Run(run)
	return _c
}

// Prefix provides a mock function with given fields: addr
func (_m *Router) Prefix(addr string) route.Router {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for Prefix")
	}

	var r0 route.Router
	if rf, ok := ret.Get(0).(func(string) route.Router); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Router)
		}
	}

	return r0
}

// Router_Prefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prefix'
type Router_Prefix_Call struct {
	*mock.Call
}

// Prefix is a helper method to define mock.On call
//   - addr string
func (_e *Router_Expecter) Prefix(addr interface{}) *Router_Prefix_Call {
	return &Router_Prefix_Call{Call: _e.mock.On("Prefix", addr)}
}

func (_c *Router_Prefix_Call) Run(run func(addr string)) *Router_Prefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Router_Prefix_Call) Return(_a0 route.Router) *Router_Prefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Router_Prefix_Call) RunAndReturn(run func(string) route.Router) *Router_Prefix_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: relativePath, handler
func (_m *Router) Put(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Router_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type Router_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Router_Expecter) Put(relativePath interface{}, handler interface{}) *Router_Put_Call {
	return &Router_Put_Call{Call: _e.mock.On("Put", relativePath, handler)}
}

func (_c *Router_Put_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Router_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Router_Put_Call) Return() *Router_Put_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Put_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Router_Put_Call {
	_c.Run(run)
	return _c
}

// Resource provides a mock function with given fields: relativePath, controller
func (_m *Router) Resource(relativePath string, controller http.ResourceController) {
	_m.Called(relativePath, controller)
}

// Router_Resource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resource'
type Router_Resource_Call struct {
	*mock.Call
}

// Resource is a helper method to define mock.On call
//   - relativePath string
//   - controller http.ResourceController
func (_e *Router_Expecter) Resource(relativePath interface{}, controller interface{}) *Router_Resource_Call {
	return &Router_Resource_Call{Call: _e.mock.On("Resource", relativePath, controller)}
}

func (_c *Router_Resource_Call) Run(run func(relativePath string, controller http.ResourceController)) *Router_Resource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.ResourceController))
	})
	return _c
}

func (_c *Router_Resource_Call) Return() *Router_Resource_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Resource_Call) RunAndReturn(run func(string, http.ResourceController)) *Router_Resource_Call {
	_c.Run(run)
	return _c
}

// Static provides a mock function with given fields: relativePath, root
func (_m *Router) Static(relativePath string, root string) {
	_m.Called(relativePath, root)
}

// Router_Static_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Static'
type Router_Static_Call struct {
	*mock.Call
}

// Static is a helper method to define mock.On call
//   - relativePath string
//   - root string
func (_e *Router_Expecter) Static(relativePath interface{}, root interface{}) *Router_Static_Call {
	return &Router_Static_Call{Call: _e.mock.On("Static", relativePath, root)}
}

func (_c *Router_Static_Call) Run(run func(relativePath string, root string)) *Router_Static_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Router_Static_Call) Return() *Router_Static_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Static_Call) RunAndReturn(run func(string, string)) *Router_Static_Call {
	_c.Run(run)
	return _c
}

// StaticFS provides a mock function with given fields: relativePath, fs
func (_m *Router) StaticFS(relativePath string, fs nethttp.FileSystem) {
	_m.Called(relativePath, fs)
}

// Router_StaticFS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StaticFS'
type Router_StaticFS_Call struct {
	*mock.Call
}

// StaticFS is a helper method to define mock.On call
//   - relativePath string
//   - fs nethttp.FileSystem
func (_e *Router_Expecter) StaticFS(relativePath interface{}, fs interface{}) *Router_StaticFS_Call {
	return &Router_StaticFS_Call{Call: _e.mock.On("StaticFS", relativePath, fs)}
}

func (_c *Router_StaticFS_Call) Run(run func(relativePath string, fs nethttp.FileSystem)) *Router_StaticFS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(nethttp.FileSystem))
	})
	return _c
}

func (_c *Router_StaticFS_Call) Return() *Router_StaticFS_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_StaticFS_Call) RunAndReturn(run func(string, nethttp.FileSystem)) *Router_StaticFS_Call {
	_c.Run(run)
	return _c
}

// StaticFile provides a mock function with given fields: relativePath, filepath
func (_m *Router) StaticFile(relativePath string, filepath string) {
	_m.Called(relativePath, filepath)
}

// Router_StaticFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StaticFile'
type Router_StaticFile_Call struct {
	*mock.Call
}

// StaticFile is a helper method to define mock.On call
//   - relativePath string
//   - filepath string
func (_e *Router_Expecter) StaticFile(relativePath interface{}, filepath interface{}) *Router_StaticFile_Call {
	return &Router_StaticFile_Call{Call: _e.mock.On("StaticFile", relativePath, filepath)}
}

func (_c *Router_StaticFile_Call) Run(run func(relativePath string, filepath string)) *Router_StaticFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Router_StaticFile_Call) Return() *Router_StaticFile_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_StaticFile_Call) RunAndReturn(run func(string, string)) *Router_StaticFile_Call {
	_c.Run(run)
	return _c
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRouter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
