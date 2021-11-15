package echohelper

import (
	"github.com/labstack/echo"
)

var (
	defaultRouters = &Routers{m: make(map[string]*Router)}
)

type RouterMethod struct {
	Path       string
	Method     string
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

type Router struct {
	Path       string
	methods    []*RouterMethod
	Middleware []echo.MiddlewareFunc
}

func (router *Router) Route(r *echo.Group) {
	g := r.Group(router.Path, router.Middleware...)
	for _, method := range router.methods {
		g.Add(method.Method, method.Path, method.Handler, method.Middleware...)
	}
}

func (router *Router) AddRouterMethod(method *RouterMethod) {
	router.methods = append(router.methods, method)
}

type Routers struct {
	m map[string]*Router
}

func (routers *Routers) AddRouter(a Action, path string, method string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	router, ok := routers.m[a.Name()]
	if !ok {
		router = &Router{
			Path:    a.Path(),
			methods: make([]*RouterMethod, 0),
		}
	}

	router.AddRouterMethod(&RouterMethod{
		Path:       path,
		Method:     method,
		Handler:    handler,
		Middleware: middleware,
	})
	routers.m[a.Name()] = router
}

func (routers *Routers) Route(g *echo.Group) {
	for _, r := range routers.m {
		r.Route(g)
	}
}

func (routers *Routers) Set(name string, router *Router) {
	routers.m[name] = router
}

func (routers *Routers) Get(name string) *Router {
	return routers.m[name]
}

func (routers *Routers) Del(name string) *Router {
	router, has := routers.m[name]
	if has {
		delete(routers.m, name)
	}
	return router
}

func AddRouter(a Action, path string, method string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	defaultRouters.AddRouter(a, path, method, handler, middleware...)
}

func Route(g *echo.Group) {
	defaultRouters.Route(g)
}

func DefaultRouters() *Routers {
	return defaultRouters
}
