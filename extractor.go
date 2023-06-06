package axum

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
)

type RequestExtractor interface {
	Extract(*http.Request) (RequestExtractor, error)
}

// Extension extractor
type Extension[T any] struct {
	Value T
}

func (e Extension[T]) Extract(req *http.Request) (RequestExtractor, error) {
	return req.Context().Value(Extension[T]{}).(RequestExtractor), nil
}

func (e *Extension[T]) Layer(inner Service) Service {
	return ServiceFunc(func(r *http.Request) ResponsePacker {
		ctx := context.WithValue(r.Context(), Extension[T]{}, *e)
		r = r.WithContext(ctx)
		return inner.Handle(r)
	})
}

func NewExtension[T any](v T) *Extension[T] {
	return &Extension[T]{Value: v}
}

// HeaderMap extractor
type HeaderMap struct {
	Value map[string]string
}

func (h HeaderMap) Extract(req *http.Request) (RequestExtractor, error) {
	h.Value = make(map[string]string)

	for k, vs := range req.Header {
		h.Value[k] = strings.Join(vs, ",")
	}

	return h, nil
}

// Json extractor
type Json[T any] struct {
	Value T
}

func (j Json[T]) Extract(req *http.Request) (RequestExtractor, error) {
	if err := json.NewDecoder(req.Body).Decode(&j.Value); err != nil {
		return nil, err
	}
	return j, nil
}

// Path extractor
type Path[T any] struct {
	Value T
}

func (p Path[T]) Extract(req *http.Request) (RequestExtractor, error) {
	params := req.Context().Value(Path[string]{}).(string)
	u, err := url.ParseQuery(params)

	if err != nil {
		return nil, err
	}

	if err := schema.NewDecoder().Decode(&p.Value, u); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Path[T]) Layer(inner Service) Service {
	return ServiceFunc(func(r *http.Request) ResponsePacker {
		ctx := context.WithValue(r.Context(), Path[T]{}, p.Value)
		r = r.WithContext(ctx)
		return inner.Handle(r)
	})
}

func NewPath[T any](v T) *Path[T] {
	return &Path[T]{Value: v}
}

// Query extractor
type Query[T any] struct {
	Value T
}

func (j Query[T]) Extract(req *http.Request) (RequestExtractor, error) {
	if err := req.ParseForm(); err != nil {
		return nil, err
	}

	if err := schema.NewDecoder().Decode(&j.Value, req.Form); err != nil {
		return nil, err
	}

	return j, nil
}

// Request extractor
type Request struct {
	Value *http.Request
}

func (r Request) Extract(req *http.Request) (RequestExtractor, error) {
	r.Value = req
	return r, nil
}
