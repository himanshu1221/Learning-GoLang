package router

import (
	"github.com/gorilla/mux"
	"github.com/himanshu1221/Learning-GoLang/25mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkedAsWatch).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
