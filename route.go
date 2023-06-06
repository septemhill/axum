package axum

import (
	"net/http"
	"strings"
)

type RouteEntry struct {
	Method       string
	Path         string
	PathSegments pathSegments
	Handler      Service
}

func (re *RouteEntry) Match(r *http.Request) bool {
	if re.Method != r.Method {
		return false
	}

	if !re.PathSegments.Match(r.URL.Path) {
		return false
	}

	re.setUrlParams(r)

	return true
}

func (re *RouteEntry) setUrlParams(r *http.Request) {
	args := make([]string, 0)
	for i := 0; i < len(re.PathSegments); i++ {
		if re.PathSegments[i].arg != "" {
			args = append(args, re.PathSegments[i].arg)
		}
	}

	path := NewPath(strings.Join(args, "&"))
	re.Handler = path.Layer(re.Handler)
}

func NewRouteEntry(method, path string, handler Service) *RouteEntry {
	re := &RouteEntry{
		Method:  method,
		Handler: handler,
	}

	path = strings.Trim(path, "/")
	segs := strings.Split(path, "/")

	ps := make([]*pathSegment, 0)
	for i := 0; i < len(segs); i++ {
		ps = append(ps, newPathSegment(segs[i]))
	}

	re.PathSegments = pathSegments(ps)

	return re
}
