package entrypoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewRequestByHttp(r *http.Request) Request {
	b, _ := ioutil.ReadAll(r.Body)
	headers := map[string]string{}
	for k, _ := range r.Header {
		headers[k] = r.Header.Get(k)
	}
	return Request{
		Body:    string(b),
		Headers: headers,
	}
}

func NewResponseToHttp(w http.ResponseWriter, status int,  response Response) {
	b,_  := json.Marshal(&response)
	w.Header().Add("Content-type","application/json")
	w.WriteHeader(status)
	if !response.Success {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(b)
}