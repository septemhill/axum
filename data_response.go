package axum

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DataResponse[T any] struct {
	Value T
}

func (rsp *DataResponse[T]) Pack() (*Response, error) {
	m := make(map[string]interface{})

	buffs := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffs).Encode(rsp.Value); err != nil {
		return nil, err
	}
	if err := json.NewDecoder(buffs).Decode(&m); err != nil {
		return nil, err
	}

	return &Response{
		Body:           m,
		HTTPStatusCode: http.StatusOK,
	}, nil
}

func NewDataResponse[T any](v T) *DataResponse[T] {
	return &DataResponse[T]{Value: v}
}

var _ ResponsePacker = (*DataResponse[int])(nil)
