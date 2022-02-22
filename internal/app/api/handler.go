package api

import (
	"net/http"
)


type Message struct{
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	IsError bool `json:"is_error"`
}

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type","application/json")
}
