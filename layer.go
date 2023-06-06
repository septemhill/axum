package axum

import (
	"net/http"
)

type Layer interface {
	Layer(Service) Service
}

type LayerService interface {
	Handle(*http.Request, Service) ResponsePacker
}

type Layer2Func func(*Request, Service) ResponsePacker

func (fn Layer2Func) Handle(r *http.Request, srv Service) ResponsePacker {
	return fn(&Request{Value: r}, srv)
}

type Layer3Func[T1 RequestExtractor] func(*T1, *Request, Service) (ResponsePacker, error)

func (fn Layer3Func[T1]) Handle(r *http.Request, srv Service) (ResponsePacker, error) {
	var ext1 T1
	t1, err := ext1.Extract(r)
	if err != nil {
		return nil, err
	}
	s1 := t1.(T1)

	return fn(&s1, &Request{Value: r}, srv)
}

func LayerArgs3Func[T1 RequestExtractor](fn Layer3Func[T1]) Layer3Func[T1] {
	return fn
}

type Layer4Func[T1, T2 RequestExtractor] func(*T1, *T2, *Request, Service) (ResponsePacker, error)

func (fn Layer4Func[T1, T2]) Handle(r *http.Request, srv Service) (ResponsePacker, error) {
	var ext1 T1
	t1, err := ext1.Extract(r)
	if err != nil {
		return nil, err
	}
	s1 := t1.(T1)

	var ext2 T2
	t2, err := ext2.Extract(r)
	if err != nil {
		return nil, err
	}
	s2 := t2.(T2)

	return fn(&s1, &s2, &Request{Value: r}, srv)
}

type Layer5Func[T1, T2, T3 RequestExtractor] func(*T1, *T2, *T3, *Request, Service) (ResponsePacker, error)

func (fn Layer5Func[T1, T2, T3]) Handle(r *http.Request, srv Service) (ResponsePacker, error) {
	var ext1 T1
	t1, err := ext1.Extract(r)
	if err != nil {
		return nil, err
	}
	s1 := t1.(T1)

	var ext2 T2
	t2, err := ext2.Extract(r)
	if err != nil {
		return nil, err
	}
	s2 := t2.(T2)

	var ext3 T3
	t3, err := ext3.Extract(r)
	if err != nil {
		return nil, err
	}
	s3 := t3.(T3)

	return fn(&s1, &s2, &s3, &Request{Value: r}, srv)
}

type Layer6Func[T1, T2, T3, T4 RequestExtractor] func(*T1, *T2, *T3, *T4, *Request, Service) (ResponsePacker, error)

func (fn Layer6Func[T1, T2, T3, T4]) Handle(r *http.Request, srv Service) (ResponsePacker, error) {
	var ext1 T1
	t1, err := ext1.Extract(r)
	if err != nil {
		return nil, err
	}
	s1 := t1.(T1)

	var ext2 T2
	t2, err := ext2.Extract(r)
	if err != nil {
		return nil, err
	}
	s2 := t2.(T2)

	var ext3 T3
	t3, err := ext3.Extract(r)
	if err != nil {
		return nil, err
	}
	s3 := t3.(T3)

	var ext4 T4
	t4, err := ext4.Extract(r)
	if err != nil {
		return nil, err
	}
	s4 := t4.(T4)

	return fn(&s1, &s2, &s3, &s4, &Request{Value: r}, srv)
}

type service struct {
	inner Service
	f     LayerService
}

func (s *service) Handle(r *http.Request) ResponsePacker {
	return s.f.Handle(r, s.inner)
}

func (s *service) Layer(inner Service) Service {
	return ServiceFunc(func(r *http.Request) ResponsePacker {
		return s.f.Handle(r, inner)
	})
}

func NewService(f LayerService) *service {
	return &service{
		f: f,
	}
}
