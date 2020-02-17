package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type DefaultResponse struct {
	w http.ResponseWriter
}

func (r *DefaultResponse) SetHeader(key string, value string) {
	r.w.Header().Set(key, value)
}

func (r *DefaultResponse) Text(statuscode int, body string) {
	r.w.Header().Set("Content-Type", "text/plain")
	_, _ = io.WriteString(*(r.Writer()), body)
	r.w.WriteHeader(statuscode)
}

func (r *DefaultResponse) Json(statuscode int, content interface{}) {
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(statuscode)

	if body, err := json.Marshal(content); err == nil {
		_, _ = r.w.Write(body)
	} else {
		r.w.WriteHeader(http.StatusInternalServerError)
	}
}

func (r *DefaultResponse) Bytes(statuscode int, body []byte) {
	r.w.Header().Set("Content-Type", "application/octet-stream")
	r.w.WriteHeader(statuscode)
	_, _ = r.w.Write(body)
}

func (r *DefaultResponse) Writer() *http.ResponseWriter {
	return &(r.w)
}
