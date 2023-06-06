package axum

import (
	"net/http"
	"path/filepath"
)

type Router struct {
	routes []*RouteEntry
}

func NewRouter() *Router {
	return &Router{
		routes: make([]*RouteEntry, 0),
	}
}

func (rt *Router) addRoute(method, path string, handler Service) *Router {
	rt.routes = append(rt.routes, NewRouteEntry(method, path, handler))
	return rt
}

func (rt *Router) Get(path string, handler Service) *Router {
	return rt.addRoute(http.MethodGet, path, handler)
}

func (rt *Router) Post(path string, handler Service) *Router {
	return rt.addRoute(http.MethodPost, path, handler)
}

func (rt *Router) Put(path string, handler Service) *Router {
	return rt.addRoute(http.MethodPut, path, handler)
}

func (rt *Router) Delete(path string, handler Service) *Router {
	return rt.addRoute(http.MethodDelete, path, handler)
}

func (rt *Router) Head(path string, handler Service) *Router {
	return rt.addRoute(http.MethodHead, path, handler)
}

func (rt *Router) Patch(path string, handler Service) *Router {
	return rt.addRoute(http.MethodPatch, path, handler)
}

func (rt *Router) SubRouter(path string, sub *Router) *Router {
	for _, re := range sub.routes {
		rt.routes = append(rt.routes, &RouteEntry{
			Method:  re.Method,
			Path:    filepath.Join(path, re.Path),
			Handler: re.Handler,
		})
	}
	return rt
}

func (rt *Router) Layer(layer Layer) *Router {
	for _, re := range rt.routes {
		re.Handler = layer.Layer(re.Handler)
	}
	return rt
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, re := range rt.routes {
		if !re.Match(r) {
			continue
		}

		packer := re.Handler.Handle(r)
		rsp, err := packer.Pack()
		if err != nil {
			// TODO: packer failed error handling
		}

		_ = rsp // TODO: response handling
		return
	}

	http.NotFound(w, r)
}
