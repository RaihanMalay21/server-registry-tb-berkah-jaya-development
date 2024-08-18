package helper

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Content-Type", "application/json") // Perbaikan typo di sini
	w.WriteHeader(code)
	w.Write(response)
}