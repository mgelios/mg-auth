package router

import (
	"encoding/json"
	"mg-auth/src/model"
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
	var loginRequest model.LoginRequest = model.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		panic(err)
	}

	res, _ := json.Marshal(loginRequest)
	print(string(res))
}
