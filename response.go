package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}

func successResponse(w http.ResponseWriter, payload interface{}) {
	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"` // Use interface{} type to accept any data
	}
	respondWithJson(w, http.StatusOK, Response{
		Status:  "success",
		Message: "success",
		Data:    payload, // Include the payload in the response
	})
}

func errorResponse(w http.ResponseWriter, message string) {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	respondWithJson(w, http.StatusBadRequest, Response{
		Status:  "error",
		Message: message,
	})
}
