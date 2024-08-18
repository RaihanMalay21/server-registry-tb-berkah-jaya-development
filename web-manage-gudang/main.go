package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"github.com/RaihanMalay21/web-gudang/controller"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/warehouse").Subrouter()
	
	api.HandleFunc("/shelfs", controller.Shelfs).Methods("GET")
	api.HandleFunc("/add/shelf", controller.AddShelf).Methods("post")
	api.HandleFunc("/remove/shelf", controller.RemoveShelf).Methods("POST")
	api.HandleFunc("/shelf/row/add", controller.AddRow).Methods("POST")
	api.HandleFunc("/shelf/row/remove", controller.RemoveRow).Methods("POST")
	// api.HandleFunc("/shelf/row/block", ).Methods("GET")
	api.HandleFunc("/shelf/row/block/add", controller.AddBlock).Methods("POST")
	api.HandleFunc("/shelf/row/block/remove", controller.RemoveBlock).Methods("POPST")
	// api.HandleFunc("/shelf/row/block/barang", ).Methods("GET")
	api.HandleFunc("/shelf/row/block/barang/add", controller.AddBarang).Methods("POST")
	// api.HandleFunc("/shelf/row/block/barang/out", ).Methods("POST")
	// api.HandleFunc("/shelf/row/block/barang/remove", ).Methods("POST")
	// api.HandleFunc("/search/barang", ).Methods("POST")

	fmt.Println("App running on port http://localhost:8080/warehouse")
	log.Fatal(http.ListenAndServe(":8080", 
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowCredentials(),
		)(r)))
}