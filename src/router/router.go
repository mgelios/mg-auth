package router

import (
	"encoding/json"
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
		router.Post("/login", GetWeatherForCity)
	})
}

func GetWeatherForCity(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "city")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode()
}
