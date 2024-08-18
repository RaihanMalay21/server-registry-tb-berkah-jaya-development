package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func RemoveShelf(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("ID_Shelf")

	// Konversi id menjadi uint
	shelfID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println("Error parsing shelf ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataShelf := models.Shelf{
		ID: uint(shelfID),
	}

	// memastikan jika barang sudah tidak tersedia lagi di dalam shelf, rows, block
	exists, err := config.DB_CheckItemsExistOrNot(dataShelf.ID)
	if err != nil {
		log.Println("Error fails to scan item exist or not in database:", err)
		msg := map[string]string{"message": "Scan barang in database exist or not fails"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}

	if exist {
		if err := config.DB_RemoveShelf(dataShelf); err != nil {
			log.Println("Error Cant delete shelf:", err)
			msg := map[string]string{"message": "Error Tidak berhasil Menghapus data Shelf"}
			helper.Response(w, msg, err)
			return
		}
		msg := map[string]string{"message": "Berhasil Menghapus Shelf"}
		helper.Response(w, msg, http.StatusOK)
	} else {
		msg := map[string]string{"message": "Tidak Dapat Menghapus Shelf, Barang Masih Tersida"}
		helper.Response(w, msg, http.StatusBadRequest)
	}

}