// Code generated by mockery. DO NOT EDIT.

package route

import (
	context "context"

	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"

	net "net"

	nethttp "net/http"

	route "github.com/goravel/framework/contracts/route"
)

// Route is an autogenerated mock type for the Route type
type Route struct {
	mock.Mock
}

type Route_Expecter struct {
	mock *mock.Mock
}

func (_m *Route) EXPECT() *Route_Expecter {
	return &Route_Expecter{mock: &_m.Mock}
}

// Any provides a mock function with given fields: relativePath, handler
func (_m *Route) Any(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Any_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Any'
type Route_Any_Call struct {
	*mock.Call
}

// Any is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Any(relativePath interface{}, handler interface{}) *Route_Any_Call {
	return &Route_Any_Call{Call: _e.mock.On("Any", relativePath, handler)}
}

func (_c *Route_Any_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Any_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Any_Call) Return() *Route_Any_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Any_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Any_Call {
	_c.Run(run)
	return _c
}

// Delete provides a mock function with given fields: relativePath, handler
func (_m *Route) Delete(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Route_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Delete(relativePath interface{}, handler interface{}) *Route_Delete_Call {
	return &Route_Delete_Call{Call: _e.mock.On("Delete", relativePath, handler)}
}

func (_c *Route_Delete_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Delete_Call) Return() *Route_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Delete_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Delete_Call {
	_c.Run(run)
	return _c
}

// Fallback provides a mock function with given fields: handler
func (_m *Route) Fallback(handler http.HandlerFunc) {
	_m.Called(handler)
}

// Route_Fallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fallback'
type Route_Fallback_Call struct {
	*mock.Call
}

// Fallback is a helper method to define mock.On call
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Fallback(handler interface{}) *Route_Fallback_Call {
	return &Route_Fallback_Call{Call: _e.mock.On("Fallback", handler)}
}

func (_c *Route_Fallback_Call) Run(run func(handler http.HandlerFunc)) *Route_Fallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Fallback_Call) Return() *Route_Fallback_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Fallback_Call) RunAndReturn(run func(http.HandlerFunc)) *Route_Fallback_Call {
	_c.Run(run)
	return _c
}

// Get provides a mock function with given fields: relativePath, handler
func (_m *Route) Get(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Route_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Get(relativePath interface{}, handler interface{}) *Route_Get_Call {
	return &Route_Get_Call{Call: _e.mock.On("Get", relativePath, handler)}
}

func (_c *Route_Get_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Get_Call) Return() *Route_Get_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Get_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Get_Call {
	_c.Run(run)
	return _c
}

// GlobalMiddleware provides a mock function with given fields: middlewares
func (_m *Route) GlobalMiddleware(middlewares ...http.Handler) {
	_va := make([]interface{}, len(middlewares))
	for _i := range middlewares {
		_va[_i] = middlewares[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Route_GlobalMiddleware_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GlobalMiddleware'
type Route_GlobalMiddleware_Call struct {
	*mock.Call
}

// GlobalMiddleware is a helper method to define mock.On call
//   - middlewares ...http.Handler
func (_e *Route_Expecter) GlobalMiddleware(middlewares ...interface{}) *Route_GlobalMiddleware_Call {
	return &Route_GlobalMiddleware_Call{Call: _e.mock.On("GlobalMiddleware",
		append([]interface{}{}, middlewares...)...)}
}

func (_c *Route_GlobalMiddleware_Call) Run(run func(middlewares ...http.Handler)) *Route_GlobalMiddleware_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]http.Handler, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(http.Handler)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Route_GlobalMiddleware_Call) Return() *Route_GlobalMiddleware_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_GlobalMiddleware_Call) RunAndReturn(run func(...http.Handler)) *Route_GlobalMiddleware_Call {
	_c.Run(run)
	return _c
}

// Group provides a mock function with given fields: handler
func (_m *Route) Group(handler route.GroupFunc) {
	_m.Called(handler)
}

// Route_Group_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Group'
type Route_Group_Call struct {
	*mock.Call
}

// Group is a helper method to define mock.On call
//   - handler route.GroupFunc
func (_e *Route_Expecter) Group(handler interface{}) *Route_Group_Call {
	return &Route_Group_Call{Call: _e.mock.On("Group", handler)}
}

func (_c *Route_Group_Call) Run(run func(handler route.GroupFunc)) *Route_Group_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(route.GroupFunc))
	})
	return _c
}

func (_c *Route_Group_Call) Return() *Route_Group_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Group_Call) RunAndReturn(run func(route.GroupFunc)) *Route_Group_Call {
	_c.Run(run)
	return _c
}

// Listen provides a mock function with given fields: l
func (_m *Route) Listen(l net.Listener) error {
	ret := _m.Called(l)

	if len(ret) == 0 {
		panic("no return value specified for Listen")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Listener) error); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_Listen_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Listen'
type Route_Listen_Call struct {
	*mock.Call
}

// Listen is a helper method to define mock.On call
//   - l net.Listener
func (_e *Route_Expecter) Listen(l interface{}) *Route_Listen_Call {
	return &Route_Listen_Call{Call: _e.mock.On("Listen", l)}
}

func (_c *Route_Listen_Call) Run(run func(l net.Listener)) *Route_Listen_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.Listener))
	})
	return _c
}

func (_c *Route_Listen_Call) Return(_a0 error) *Route_Listen_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_Listen_Call) RunAndReturn(run func(net.Listener) error) *Route_Listen_Call {
	_c.Call.Return(run)
	return _c
}

// ListenTLS provides a mock function with given fields: l
func (_m *Route) ListenTLS(l net.Listener) error {
	ret := _m.Called(l)

	if len(ret) == 0 {
		panic("no return value specified for ListenTLS")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Listener) error); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_ListenTLS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListenTLS'
type Route_ListenTLS_Call struct {
	*mock.Call
}

// ListenTLS is a helper method to define mock.On call
//   - l net.Listener
func (_e *Route_Expecter) ListenTLS(l interface{}) *Route_ListenTLS_Call {
	return &Route_ListenTLS_Call{Call: _e.mock.On("ListenTLS", l)}
}

func (_c *Route_ListenTLS_Call) Run(run func(l net.Listener)) *Route_ListenTLS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.Listener))
	})
	return _c
}

func (_c *Route_ListenTLS_Call) Return(_a0 error) *Route_ListenTLS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_ListenTLS_Call) RunAndReturn(run func(net.Listener) error) *Route_ListenTLS_Call {
	_c.Call.Return(run)
	return _c
}

// ListenTLSWithCert provides a mock function with given fields: l, certFile, keyFile
func (_m *Route) ListenTLSWithCert(l net.Listener, certFile string, keyFile string) error {
	ret := _m.Called(l, certFile, keyFile)

	if len(ret) == 0 {
		panic("no return value specified for ListenTLSWithCert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Listener, string, string) error); ok {
		r0 = rf(l, certFile, keyFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_ListenTLSWithCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListenTLSWithCert'
type Route_ListenTLSWithCert_Call struct {
	*mock.Call
}

// ListenTLSWithCert is a helper method to define mock.On call
//   - l net.Listener
//   - certFile string
//   - keyFile string
func (_e *Route_Expecter) ListenTLSWithCert(l interface{}, certFile interface{}, keyFile interface{}) *Route_ListenTLSWithCert_Call {
	return &Route_ListenTLSWithCert_Call{Call: _e.mock.On("ListenTLSWithCert", l, certFile, keyFile)}
}

func (_c *Route_ListenTLSWithCert_Call) Run(run func(l net.Listener, certFile string, keyFile string)) *Route_ListenTLSWithCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.Listener), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Route_ListenTLSWithCert_Call) Return(_a0 error) *Route_ListenTLSWithCert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_ListenTLSWithCert_Call) RunAndReturn(run func(net.Listener, string, string) error) *Route_ListenTLSWithCert_Call {
	_c.Call.Return(run)
	return _c
}

// Middleware provides a mock function with given fields: middlewares
func (_m *Route) Middleware(middlewares ...http.Handler) route.Router {
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
	if rf, ok := ret.Get(0).(func(...http.Handler) route.Router); ok {
		r0 = rf(middlewares...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Router)
		}
	}

	return r0
}

// Route_Middleware_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Middleware'
type Route_Middleware_Call struct {
	*mock.Call
}

// Middleware is a helper method to define mock.On call
//   - middlewares ...http.Handler
func (_e *Route_Expecter) Middleware(middlewares ...interface{}) *Route_Middleware_Call {
	return &Route_Middleware_Call{Call: _e.mock.On("Middleware",
		append([]interface{}{}, middlewares...)...)}
}

func (_c *Route_Middleware_Call) Run(run func(middlewares ...http.Handler)) *Route_Middleware_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]http.Handler, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(http.Handler)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Route_Middleware_Call) Return(_a0 route.Router) *Route_Middleware_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_Middleware_Call) RunAndReturn(run func(...http.Handler) route.Router) *Route_Middleware_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: relativePath, handler
func (_m *Route) Options(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type Route_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Options(relativePath interface{}, handler interface{}) *Route_Options_Call {
	return &Route_Options_Call{Call: _e.mock.On("Options", relativePath, handler)}
}

func (_c *Route_Options_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Options_Call) Return() *Route_Options_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Options_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Options_Call {
	_c.Run(run)
	return _c
}

// Patch provides a mock function with given fields: relativePath, handler
func (_m *Route) Patch(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Patch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Patch'
type Route_Patch_Call struct {
	*mock.Call
}

// Patch is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Patch(relativePath interface{}, handler interface{}) *Route_Patch_Call {
	return &Route_Patch_Call{Call: _e.mock.On("Patch", relativePath, handler)}
}

func (_c *Route_Patch_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Patch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Patch_Call) Return() *Route_Patch_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Patch_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Patch_Call {
	_c.Run(run)
	return _c
}

// Post provides a mock function with given fields: relativePath, handler
func (_m *Route) Post(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type Route_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Post(relativePath interface{}, handler interface{}) *Route_Post_Call {
	return &Route_Post_Call{Call: _e.mock.On("Post", relativePath, handler)}
}

func (_c *Route_Post_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Post_Call) Return() *Route_Post_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Post_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Post_Call {
	_c.Run(run)
	return _c
}

// Prefix provides a mock function with given fields: addr
func (_m *Route) Prefix(addr string) route.Router {
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

// Route_Prefix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prefix'
type Route_Prefix_Call struct {
	*mock.Call
}

// Prefix is a helper method to define mock.On call
//   - addr string
func (_e *Route_Expecter) Prefix(addr interface{}) *Route_Prefix_Call {
	return &Route_Prefix_Call{Call: _e.mock.On("Prefix", addr)}
}

func (_c *Route_Prefix_Call) Run(run func(addr string)) *Route_Prefix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Route_Prefix_Call) Return(_a0 route.Router) *Route_Prefix_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_Prefix_Call) RunAndReturn(run func(string) route.Router) *Route_Prefix_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: relativePath, handler
func (_m *Route) Put(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Route_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type Route_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - relativePath string
//   - handler http.HandlerFunc
func (_e *Route_Expecter) Put(relativePath interface{}, handler interface{}) *Route_Put_Call {
	return &Route_Put_Call{Call: _e.mock.On("Put", relativePath, handler)}
}

func (_c *Route_Put_Call) Run(run func(relativePath string, handler http.HandlerFunc)) *Route_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.HandlerFunc))
	})
	return _c
}

func (_c *Route_Put_Call) Return() *Route_Put_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Put_Call) RunAndReturn(run func(string, http.HandlerFunc)) *Route_Put_Call {
	_c.Run(run)
	return _c
}

// Recover provides a mock function with given fields: recoverFunc
func (_m *Route) Recover(recoverFunc func(http.Context, interface{})) {
	_m.Called(recoverFunc)
}

// Route_Recover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recover'
type Route_Recover_Call struct {
	*mock.Call
}

// Recover is a helper method to define mock.On call
//   - recoverFunc func(http.Context , interface{})
func (_e *Route_Expecter) Recover(recoverFunc interface{}) *Route_Recover_Call {
	return &Route_Recover_Call{Call: _e.mock.On("Recover", recoverFunc)}
}

func (_c *Route_Recover_Call) Run(run func(recoverFunc func(http.Context, interface{}))) *Route_Recover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(http.Context, interface{})))
	})
	return _c
}

func (_c *Route_Recover_Call) Return() *Route_Recover_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Recover_Call) RunAndReturn(run func(func(http.Context, interface{}))) *Route_Recover_Call {
	_c.Run(run)
	return _c
}

// Resource provides a mock function with given fields: relativePath, controller
func (_m *Route) Resource(relativePath string, controller http.ResourceController) {
	_m.Called(relativePath, controller)
}

// Route_Resource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resource'
type Route_Resource_Call struct {
	*mock.Call
}

// Resource is a helper method to define mock.On call
//   - relativePath string
//   - controller http.ResourceController
func (_e *Route_Expecter) Resource(relativePath interface{}, controller interface{}) *Route_Resource_Call {
	return &Route_Resource_Call{Call: _e.mock.On("Resource", relativePath, controller)}
}

func (_c *Route_Resource_Call) Run(run func(relativePath string, controller http.ResourceController)) *Route_Resource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(http.ResourceController))
	})
	return _c
}

func (_c *Route_Resource_Call) Return() *Route_Resource_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Resource_Call) RunAndReturn(run func(string, http.ResourceController)) *Route_Resource_Call {
	_c.Run(run)
	return _c
}

// Run provides a mock function with given fields: host
func (_m *Route) Run(host ...string) error {
	_va := make([]interface{}, len(host))
	for _i := range host {
		_va[_i] = host[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(host...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type Route_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - host ...string
func (_e *Route_Expecter) Run(host ...interface{}) *Route_Run_Call {
	return &Route_Run_Call{Call: _e.mock.On("Run",
		append([]interface{}{}, host...)...)}
}

func (_c *Route_Run_Call) Run(run func(host ...string)) *Route_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Route_Run_Call) Return(_a0 error) *Route_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_Run_Call) RunAndReturn(run func(...string) error) *Route_Run_Call {
	_c.Call.Return(run)
	return _c
}

// RunTLS provides a mock function with given fields: host
func (_m *Route) RunTLS(host ...string) error {
	_va := make([]interface{}, len(host))
	for _i := range host {
		_va[_i] = host[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RunTLS")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(host...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_RunTLS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunTLS'
type Route_RunTLS_Call struct {
	*mock.Call
}

// RunTLS is a helper method to define mock.On call
//   - host ...string
func (_e *Route_Expecter) RunTLS(host ...interface{}) *Route_RunTLS_Call {
	return &Route_RunTLS_Call{Call: _e.mock.On("RunTLS",
		append([]interface{}{}, host...)...)}
}

func (_c *Route_RunTLS_Call) Run(run func(host ...string)) *Route_RunTLS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Route_RunTLS_Call) Return(_a0 error) *Route_RunTLS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_RunTLS_Call) RunAndReturn(run func(...string) error) *Route_RunTLS_Call {
	_c.Call.Return(run)
	return _c
}

// RunTLSWithCert provides a mock function with given fields: host, certFile, keyFile
func (_m *Route) RunTLSWithCert(host string, certFile string, keyFile string) error {
	ret := _m.Called(host, certFile, keyFile)

	if len(ret) == 0 {
		panic("no return value specified for RunTLSWithCert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(host, certFile, keyFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_RunTLSWithCert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunTLSWithCert'
type Route_RunTLSWithCert_Call struct {
	*mock.Call
}

// RunTLSWithCert is a helper method to define mock.On call
//   - host string
//   - certFile string
//   - keyFile string
func (_e *Route_Expecter) RunTLSWithCert(host interface{}, certFile interface{}, keyFile interface{}) *Route_RunTLSWithCert_Call {
	return &Route_RunTLSWithCert_Call{Call: _e.mock.On("RunTLSWithCert", host, certFile, keyFile)}
}

func (_c *Route_RunTLSWithCert_Call) Run(run func(host string, certFile string, keyFile string)) *Route_RunTLSWithCert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Route_RunTLSWithCert_Call) Return(_a0 error) *Route_RunTLSWithCert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_RunTLSWithCert_Call) RunAndReturn(run func(string, string, string) error) *Route_RunTLSWithCert_Call {
	_c.Call.Return(run)
	return _c
}

// ServeHTTP provides a mock function with given fields: writer, request
func (_m *Route) ServeHTTP(writer nethttp.ResponseWriter, request *nethttp.Request) {
	_m.Called(writer, request)
}

// Route_ServeHTTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServeHTTP'
type Route_ServeHTTP_Call struct {
	*mock.Call
}

// ServeHTTP is a helper method to define mock.On call
//   - writer nethttp.ResponseWriter
//   - request *nethttp.Request
func (_e *Route_Expecter) ServeHTTP(writer interface{}, request interface{}) *Route_ServeHTTP_Call {
	return &Route_ServeHTTP_Call{Call: _e.mock.On("ServeHTTP", writer, request)}
}

func (_c *Route_ServeHTTP_Call) Run(run func(writer nethttp.ResponseWriter, request *nethttp.Request)) *Route_ServeHTTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(nethttp.ResponseWriter), args[1].(*nethttp.Request))
	})
	return _c
}

func (_c *Route_ServeHTTP_Call) Return() *Route_ServeHTTP_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_ServeHTTP_Call) RunAndReturn(run func(nethttp.ResponseWriter, *nethttp.Request)) *Route_ServeHTTP_Call {
	_c.Run(run)
	return _c
}

// Shutdown provides a mock function with given fields: ctx
func (_m *Route) Shutdown(ctx ...context.Context) error {
	_va := make([]interface{}, len(ctx))
	for _i := range ctx {
		_va[_i] = ctx[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Shutdown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...context.Context) error); ok {
		r0 = rf(ctx...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Route_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type Route_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
//   - ctx ...context.Context
func (_e *Route_Expecter) Shutdown(ctx ...interface{}) *Route_Shutdown_Call {
	return &Route_Shutdown_Call{Call: _e.mock.On("Shutdown",
		append([]interface{}{}, ctx...)...)}
}

func (_c *Route_Shutdown_Call) Run(run func(ctx ...context.Context)) *Route_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]context.Context, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(context.Context)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Route_Shutdown_Call) Return(_a0 error) *Route_Shutdown_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Route_Shutdown_Call) RunAndReturn(run func(...context.Context) error) *Route_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// Static provides a mock function with given fields: relativePath, root
func (_m *Route) Static(relativePath string, root string) {
	_m.Called(relativePath, root)
}

// Route_Static_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Static'
type Route_Static_Call struct {
	*mock.Call
}

// Static is a helper method to define mock.On call
//   - relativePath string
//   - root string
func (_e *Route_Expecter) Static(relativePath interface{}, root interface{}) *Route_Static_Call {
	return &Route_Static_Call{Call: _e.mock.On("Static", relativePath, root)}
}

func (_c *Route_Static_Call) Run(run func(relativePath string, root string)) *Route_Static_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Route_Static_Call) Return() *Route_Static_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_Static_Call) RunAndReturn(run func(string, string)) *Route_Static_Call {
	_c.Run(run)
	return _c
}

// StaticFS provides a mock function with given fields: relativePath, fs
func (_m *Route) StaticFS(relativePath string, fs nethttp.FileSystem) {
	_m.Called(relativePath, fs)
}

// Route_StaticFS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StaticFS'
type Route_StaticFS_Call struct {
	*mock.Call
}

// StaticFS is a helper method to define mock.On call
//   - relativePath string
//   - fs nethttp.FileSystem
func (_e *Route_Expecter) StaticFS(relativePath interface{}, fs interface{}) *Route_StaticFS_Call {
	return &Route_StaticFS_Call{Call: _e.mock.On("StaticFS", relativePath, fs)}
}

func (_c *Route_StaticFS_Call) Run(run func(relativePath string, fs nethttp.FileSystem)) *Route_StaticFS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(nethttp.FileSystem))
	})
	return _c
}

func (_c *Route_StaticFS_Call) Return() *Route_StaticFS_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_StaticFS_Call) RunAndReturn(run func(string, nethttp.FileSystem)) *Route_StaticFS_Call {
	_c.Run(run)
	return _c
}

// StaticFile provides a mock function with given fields: relativePath, filepath
func (_m *Route) StaticFile(relativePath string, filepath string) {
	_m.Called(relativePath, filepath)
}

// Route_StaticFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StaticFile'
type Route_StaticFile_Call struct {
	*mock.Call
}

// StaticFile is a helper method to define mock.On call
//   - relativePath string
//   - filepath string
func (_e *Route_Expecter) StaticFile(relativePath interface{}, filepath interface{}) *Route_StaticFile_Call {
	return &Route_StaticFile_Call{Call: _e.mock.On("StaticFile", relativePath, filepath)}
}

func (_c *Route_StaticFile_Call) Run(run func(relativePath string, filepath string)) *Route_StaticFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Route_StaticFile_Call) Return() *Route_StaticFile_Call {
	_c.Call.Return()
	return _c
}

func (_c *Route_StaticFile_Call) RunAndReturn(run func(string, string)) *Route_StaticFile_Call {
	_c.Run(run)
	return _c
}

// Test provides a mock function with given fields: request
func (_m *Route) Test(request *nethttp.Request) (*nethttp.Response, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for Test")
	}

	var r0 *nethttp.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*nethttp.Request) (*nethttp.Response, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*nethttp.Request) *nethttp.Response); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nethttp.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*nethttp.Request) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Route_Test_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Test'
type Route_Test_Call struct {
	*mock.Call
}

// Test is a helper method to define mock.On call
//   - request *nethttp.Request
func (_e *Route_Expecter) Test(request interface{}) *Route_Test_Call {
	return &Route_Test_Call{Call: _e.mock.On("Test", request)}
}

func (_c *Route_Test_Call) Run(run func(request *nethttp.Request)) *Route_Test_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*nethttp.Request))
	})
	return _c
}

func (_c *Route_Test_Call) Return(_a0 *nethttp.Response, _a1 error) *Route_Test_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Route_Test_Call) RunAndReturn(run func(*nethttp.Request) (*nethttp.Response, error)) *Route_Test_Call {
	_c.Call.Return(run)
	return _c
}

// NewRoute creates a new instance of Route. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRoute(t interface {
	mock.TestingT
	Cleanup(func())
}) *Route {
	mock := &Route{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
