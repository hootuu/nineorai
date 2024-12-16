package io

import (
	"encoding/json"
	"github.com/hootuu/gelato/crtpto/rand"
	"github.com/hootuu/gelato/errors"
	"time"
)

type Response[T any] struct {
	RequestID string        `json:"request_id" bson:"request_id"`
	Success   bool          `json:"success" bson:"success"`
	Data      *T            `json:"data,omitempty" bson:"data,omitempty"`
	Error     *errors.Error `json:"error,omitempty" bson:"error,omitempty"`
	Timestamp int64         `json:"timestamp" bson:"timestamp"`
	Nonce     int64         `json:"nonce" bson:"nonce"`
}

func NewResponse[T any](requestID string, data *T) *Response[T] {
	randInt64, _ := rand.Int64()
	return &Response[T]{
		RequestID: requestID,
		Success:   true,
		Data:      data,
		Error:     nil,
		Timestamp: time.Now().UnixMicro(),
		Nonce:     randInt64,
	}
}

func FailResponse[T any](requestID string, err *errors.Error) *Response[T] {
	randInt64, _ := rand.Int64()
	return &Response[T]{
		RequestID: requestID,
		Success:   false,
		Data:      nil,
		Error:     err,
		Timestamp: time.Now().UnixMilli(),
		Nonce:     randInt64,
	}
}

func (r *Response[T]) Marshal() ([]byte, *errors.Error) {
	byteArr, err := json.Marshal(r)
	if err != nil {
		return nil, errors.System("json_marshal_failed", err)
	}
	return byteArr, nil
}

func (r *Response[T]) Json() string {
	byteArr, err := json.Marshal(r)
	if err != nil {
		return "{\"result\":false}"
	}
	return string(byteArr)
}

func UnmarshalResponse[T any](data []byte) (*Response[T], *errors.Error) {
	var r Response[T]
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, errors.E("json_unmarshal_failed", "unmarshal response failed: %w", err)
	}
	return &r, nil
}
