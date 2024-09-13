package main

import (
	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	"net/http"
	"log"

	"github.com/RaihanMalay21/server-registry-tb-berkah-jaya-development/controller"
	"github.com/RaihanMalay21/server-registry-tb-berkah-jaya-development/controller/template"
	config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
)

func main() {
	r := mux.NewRouter()

	config.DB_Connection()
	// r.Use(corsMiddlewares)
	api := r.PathPrefix("/berkahjaya").Subrouter()
	api.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
	api.HandleFunc("/signup", controller.SignUp).Methods("POST", "OPTIONS") 
	api.HandleFunc("/logout", controller.LogOut).Methods("GET", "OPTIONS")
	// r.HandleFunc("/get/hadiah", controller.Hadiah).Methods("GET")
	api.HandleFunc("/forgot/password", controller.ForgotPassword).Methods("POST", "OPTIONS")
	api.HandleFunc("/forgot/password/reset", template.PageResetPassword).Methods("GET", "OPTIONS")
	api.HandleFunc("/forgot/password/reset", controller.ForgotPasswordChangePassword).Methods("POST", "OPTIONS")
	
	// corsHandler := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"https://fe-tb-berkah-jaya-igcfjdj5fa-uc.a.run.app"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 	handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	handlers.AllowCredentials(),
	// )

	log.Fatal(http.ListenAndServe(":8082", r))
}

// func corsMiddlewares(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		origin := r.Header.Get("Origin")

// 		allowedOrigins := "https://fe-tb-berkah-jaya-750892348569.us-central1.run.app"

// 		if origin == allowedOrigins {
// 			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
// 			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Authorization")
// 			w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		}

// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }