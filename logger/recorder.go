package logger

import (
	"encoding/json"
	"time"
)

type Record struct {
	Request string
	Timing  time.Duration
}

type RequestRecorder struct {
	requests []Record
}

func NewRequestRecorder() *RequestRecorder {
	r := &RequestRecorder{}
	return r
}

func (r *RequestRecorder) add(request Record) {
	r.requests = append(r.requests, request)
}

func (r *RequestRecorder) ToJson() ([]byte, error) {
	js, err := json.Marshal(r.requests)
	return js, err
}
