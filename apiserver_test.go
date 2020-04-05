package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/api?key=hello&value=gopher", nil)

	if err != nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiServer)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want % v", status, http.StatusOK)
	}
	expected := `{"Key":"hello","Value":"gopher"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

//func TestApiServerWithoutQueryParams(t *testing.T) {
//	req, err := http.NewRequest("GET", "/api", nil)
//
//	if err != nil{
//		t.Fatal(err)
//	}
//
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(ApiServer)
//
//	handler.ServeHTTP(rr, req)
//
//	if status := rr.Code; status != http.StatusInternalServerError {
//		t.Errorf("handler returned wrong status code: got %v want % v", status, http.StatusInternalServerError)
//	}
//	expected := 0
//
//	if rr.Body.Len() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
//	}
//}
