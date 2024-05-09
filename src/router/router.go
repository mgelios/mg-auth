package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RunServer() {
	router := chi.NewRouter()
	Handler(router)
	println("Started serving requests")
	err := http.ListenAndServe("localhost:8800", router)
	if err != nil {
		panic(err)
	}
}

func Handler(router *chi.Mux) {
	router.Route("/api/v1", func(router chi.Router) {
		router.Post("/login", HandleLogin)
	})
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
}
