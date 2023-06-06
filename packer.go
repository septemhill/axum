package axum

import "net/http"

type Response struct {
	Body           map[string]interface{}
	HTTPStatusCode int
	Headers        http.Header
	Cookies        http.Cookie
}

type ResponsePacker interface {
	Pack() (*Response, error)
}
