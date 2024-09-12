package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/RaihanMalay21/api-gateway-tb-berkah-jaya-development/middlewares"
)

func main() {
	r := mux.NewRouter()
	r.Use(middlewares.CorsMiddlewares)
	r.PathPrefix("/customer").Handler(middlewares.ReverseProxy("http://localhost:8081"))
	r.PathPrefix("/access").Handler(middlewares.ReverseProxy("http://localhost:8081"))
	r.PathPrefix("/admin").Handler(middlewares.ReverseProxy("http://localhost:8083"))
	log.Fatal(http.ListenAndServe(":8080", r))
}