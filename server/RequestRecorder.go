package main

import (
	"encoding/json"
	"time"
)

type RequestRecord struct {
	Request string
	Timing  time.Duration
}

type RequestRecorder struct {
	requests []RequestRecord
}

func NewRequestRecorder() *RequestRecorder {
	r := &RequestRecorder{}
	return r
}

func (r *RequestRecorder) add(request RequestRecord) {
	r.requests = append(r.requests, request)
}

func (r *RequestRecorder) toJson() ([]byte, error) {
	js, err := json.Marshal(r.requests)
	return js, err
}
