package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/RaihanMalay21/api-gateway-tb-berkah-jaya/middlewares"
)

func main() {
	r := mux.NewRouter()

	r.Use(middlewares.CorsMiddlewares)

	api := r.PathPrefix("/berkahjaya").Subrouter()
	api.PathPrefix("/access").Handler(middlewares.ReverseProxy("https://server-registry-tb-berkah-jaya-750892348569.us-central1.run.app"))
	api.PathPrefix("/custemer").Handler(middlewares.ReverseProxy("https://server-customer-tb-berkah-jaya-750892348569.us-central1.run.app"))

	log.Fatal(http.ListenAndServe(":8080", r))
}