package main

import "encoding/json"

type RequestRecorder struct {
	requests []string
}

func NewRequestRecorder() *RequestRecorder{
	r := &RequestRecorder{}
return r
}

func (r *RequestRecorder) add (request string){
	r.requests = append(r.requests, request)
}

func (r *RequestRecorder) toJson () ([]byte, error){
	js, err := json.Marshal(r.requests)
	return js,err
}


