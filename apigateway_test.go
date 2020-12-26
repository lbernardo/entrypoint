package entrypoint

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestNewRequestByApiGateway(t *testing.T) {
	b := NewRequestByApiGateway(RequestApiGateway{
		Body: "{\"status\":\"OK\"}",
	})
	if b.Body != "{\"status\":\"OK\"}" {
		t.Errorf("Was expected %v and return %v", "{\"status\":\"OK\"}", b.Body)
	}
}

func TestNewResponseToApiGateway(t *testing.T) {
	re := Response{Success: true}
	bo,_ := json.Marshal(&re)
	r := NewResponseToApiGateway(http.StatusOK, re)
	if r.StatusCode != http.StatusOK {
		t.Errorf("Was expected %v and return %v\n", http.StatusOK, r.StatusCode)
	}
	if r.Body != string(bo) {
		t.Errorf("Was expected %v and return %v", string(bo), r.Body)
	}
}
