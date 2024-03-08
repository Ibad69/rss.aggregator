package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerUserRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		successResponse(w, "HI from go222")
	})
	r.Get("/go", func(w http.ResponseWriter, r *http.Request) {
		successResponse(w, "HI from go222")
	})
}

func registerAdminRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		successResponse(w, "HI from go222")
	})
	r.Get("/go", func(w http.ResponseWriter, r *http.Request) {
		successResponse(w, "Hi json object")
	})
}
