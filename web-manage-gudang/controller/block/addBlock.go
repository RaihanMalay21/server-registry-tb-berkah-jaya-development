package controller

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func AddBlock(w http.ResponseWriter, r *http.Request) {
	var field map[string]string

	JSON := json.NewDecoder(r.Body)
	if err :=  JSON.Decode(&field); err != nil {
		log.Println("Error cant decode json:", err)
		msg := map[string]string{"message": "Error Tidak Dapat Mendecode data json"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}

	idRow, err := strconv.ParseUint(field["id_row"], 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	nomorblock, err := strconv.ParseUint(field["nomor_block"], 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	capacitybarang, err := strconv.ParseFloat(field["capacity_barang"], 64)
	if err != nil {
		fmt.Println("Error parsing string to float64:", err)
		return
	}

	capacityblock, err := strconv.ParseFloat(field["capacity_block"], 64)
	if err != nil {
		fmt.Println("Error parsing string to float64:", err)
		return
	}

	if capacityblock == nomorblock {
		msg := map[string]string{"message": "Kuota Block Sudah Habis"}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	Blocks := models.Block{
		NomorBlock: uint(nomorblock),
		CapacityBarang: capacitybarang,
	}

	if err := config.DB_AddBlock(Block, idRow); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := map[string]string{"message": "berhasil Membuat Block"}
	helper.Response(w, msg, http.StatusOK)
}