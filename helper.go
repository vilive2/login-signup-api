package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ErrorJSON(w http.ResponseWriter, err error, status int) {
	resp := Response{
		Status:  "error",
		Message: err.Error(),
		Data:    nil,
	}

	body, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

func WriteJSON(w http.ResponseWriter, data any) {
	resp := Response{
		Status:  "success",
		Message: "",
		Data:    data,
	}

	body, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
