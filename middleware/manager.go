package middleware

import "net/http"

// Middleware is just a function type, it says if anyone uses me it have to provide
// me func and it's param will be http.Handler and it'll also return http.Handler
type Middleware func(http.Handler) http.Handler

// `Manager` is a struct that contains `globalMiddlewares`, which holds a slice of `Middleware`.
// Since `Middleware` is of type `func(http.Handler) http.Handler`,
// `globalMiddlewares` can store many functions that accept an `http.Handler` and
// return an `http.Handler`, and use them whenever needed.
type Manager struct {
	globalMiddlewares []Middleware
}

// `NewManager` creates a new `Manager` with an empty list of `globalMiddlewares`, ready to store middleware functions later.
// It returns a pointer (*Manager) so we can use and modify the same instance without making a copy.
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
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	// `mngr` is a pointer to the Manager instance, so it can directly access and modify `globalMiddlewares` inside the Manager.
	// `globalMiddlewares` is a slice of `Middleware`. `Middleware` itself is a function type: `func(http.Handler) http.Handler`.
	// `middlewares` is a slice that holds all Middleware functions passed to `Use` using the variadic parameter (`...Middleware`).
	// `append` adds all Middleware functions from `middlewares` into `mngr.globalMiddlewares`.
	// The `...` expands the slice, so every middleware function is appended one by one, not the slice itself.

}

// This is a local middleware helper that takes a handler (route) and wraps it with the given middleware.
// With takes *Manager for consistency and future flexibility, even though it doesn’t use the Manager state right now.
// With takes an http.Handler and any number of Middleware functions.
// It starts by storing the original handler in `h`.
// Then it loops through each middleware in `middlewares`.
// Each middleware wraps the current handler (`h`) and returns a new handler, which is assigned back to `h`.
// Finally, it returns the fully wrapped handler.
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
