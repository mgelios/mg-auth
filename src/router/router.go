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
	var user model.User = model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

}
