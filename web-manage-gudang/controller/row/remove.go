package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func RemoveRow(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("ID_Row")

	rowID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Println("Error parsing shelf ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	exist, err := DB_CheckItemsExistOrNotRow(rowID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exist {
		if err := DB_RemoveRow(rowID); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		msg := map[string]string{"message" : "Berhasil Menghapus Row"}
		helper.Response(w, msg, http.StatusOK)
	} else {
		msg := map[string]string{"message" : "Tidak Dapat Menghapus Row, Barang Masih Tersida"}
		helper.Response(w, msg, http.StatusInternalServerError)
	}
}