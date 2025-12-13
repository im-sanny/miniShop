package middleware

import "net/http"

// Middleware is just a function type, it says if anyone uses me it have to provide
// me func and it's param will be http.Handler and it'll also return http.Handler
type Middleware func(http.Handler) http.Handler

// `Manager` is a struct that contains `globalMiddlewares`, which holds a slice of
// `Middleware`. Since `Middleware` is of type `func(http.Handler) http.Handler`,
// `globalMiddlewares` can store many functions that accept an `http.Handler` and
// return an `http.Handler`, and use them whenever needed.
type Manager struct {
	globalMiddlewares []Middleware
}

// NewManager creates a new Manager with an empty list of globalMiddlewares,
// ready to store middleware functions later.
// It returns a pointer (*Manager) so we can use and modify the same instance
// without making a copy.
// &Manager gets the address of the new instance so the pointer points to it.
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0), // empty slice to hold middleware functions
	}
}

// mngr parameter takes the instance of Manager
// middlewares parameter takes any number of Middleware functions (because of ...Middleware)
// inside Use, middlewares behaves like a slice holding all passed middleware functions, ready to be used later
func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...) // mngr called globalMiddlewares from Manager since it's using instance of Manager by pointer, and it appends with Middleware which is a function type, so i can say that globalMiddlewares append with slice of Middleware with Middleware function type, but if we go deep then we can see that globalMiddlewares is slice of func(http.Handler) http.Handler under the hood and Middleware is func(http.Handler) http.Handler itself
}

// this local middleware will take a routes and wrap it with middleware
func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

// global middleware, this one will be used everywhere
func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
