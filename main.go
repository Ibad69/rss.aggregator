package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	os.Getenv("PORT")

	// r.route is a path prefix opnly it does not have it's own router it will work on the same router itself
	// use this when you very less apis
	r := chi.NewRouter()
	r.Route("/opt", func(ro chi.Router) {
		ro.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("HI from go222"))
		})
		ro.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("HI from go222"))
		})
		ro.Post("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("HI from go222"))
		})
	})
	// r.Mount is a path prefix and it has it's own router, seperate middleware & will route seperately mounting
	adminRouter := chi.NewRouter()
	registerAdminRoutes(adminRouter)
	r.Mount("/admin", adminRouter)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
