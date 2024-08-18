package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/go-playground/validator/v10"
)

func AddShelf(w http.ResponseWriter, r *http.Request) {
	var shelf models.Shelf
	JSON := json.NewDecoder(r.Body)
	if err :=  JSON.Decode(&shelf); err != nil {
		log.Println("Error cant decode json:", err)
		msg := map[string]string{"message": "Error Tidak Dapat Mendecode data json"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}

	// inialisasi validator
	validate := validator.New(validator.WithRequiredStructEnabled())
	Trans := helper.TranslatorIDN()

	if err := validate.Struct(&shelf); err != nil {
		// map untuk menyimpan pesan error
		errors := make(map[string]string)

		// menyimpan errors kedalam map error berupa field dan pesannya
		for _, err := range err.(validator.ValidationErrors) {
			NameField := err.StructField()
			errTranlate := err.Translate(Trans)
			errors[NameField] = errTranlate
		}

		helper.Response(w, errors, http.StatusInternalServerError)
		return
	} 

	if err := config.DB_AddShelf(shelf); err != nil {
		log.Println("Error cant insert data shelf to database:", err)
		msg := map[string]string{"message": "Error cant insert data to database"}
		helper.Response(w, msg, http.StatusInternalServerError)
	}

	msg := map[string]string{"message": "successfuly add shelf"}
	helper.Response(w, msg, http.StatusOK)
}