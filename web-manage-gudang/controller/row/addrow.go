package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"github.com/RaihanMalay21/web-gudang/models"
)

func AddRow(w http.ResponseWriter, r *http.Request) {
	var field map[string]string

	JSON := json.NewDecoder(r.Body)
	if err :=  JSON.Decode(&field); err != nil {
		log.Println("Error cant decode json:", err)
		msg := map[string]string{"message": "Error Tidak Dapat Mendecode data json"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}

	idShelf, err := strconv.ParseUint(field["id_shelf"], 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	numberrow, err := strconv.ParseUint(field["number_row"], 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	capacityblock, err := strconv.ParseFloat(field["capacity_block"], 64)
	if err != nil {
		fmt.Println("Error parsing string to float64:", err)
		return
	}

	capacityrow, err := strconv.ParseFloat(field["capacity_row"], 64)
	if err != nil {
		fmt.Println("Error parsing string to float64:", err)
		return
	}

	if capacityrow == numberrow {
		msg := map[string]string{"message": "Kouta Row Sudah habis"}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	rows := models.Row{
		// data yang dibutuhkan 
		NumberRow: uint(numberrow),
		CapacityBlock: field[capacityblock],
	}

	if err := config.DB_AddRows(uint(idShelf), rows); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := map[string]string{"message": "berhasil Membuat Row"}
	helper.Response(w, msg, http.StatusOK)
}