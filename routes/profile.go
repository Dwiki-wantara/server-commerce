package routes

import (
	"blendpedia/handlers"
	"blendpedia/pkg/middleware"
	"blendpedia/pkg/mysql"
	"blendpedia/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profile", middleware.Auth(h.GetProfile)).Methods("GET")
}
