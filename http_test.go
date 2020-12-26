package entrypoint

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRequestByHttp(t *testing.T) {
	r := &http.Request{
		Body: ioutil.NopCloser(bytes.NewBufferString("{\"status\": \"ok\"}")),
	}
	byHttp := NewRequestByHttp(r)
	if byHttp.Body != "{\"status\": \"ok\"}" {
		t.Errorf("Was expected %v and return %v\n", "{\"status\": \"ok\"}", byHttp.Body)
	}
}

func TestNewResponseToHttp(t *testing.T) {
	w := httptest.NewRecorder()
	NewResponseToHttp(w, http.StatusOK, Response{Success: true})
	if w.Header().Get("Content-type") != "application/json" {
		t.Errorf("Was expected application/json and return %v\n", w.Header().Get("Content-type"))
	}
	if w.Code != http.StatusOK {
		t.Errorf("Was expected %v and return %v", http.StatusOK, w.Code)
	}
}
