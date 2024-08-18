package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func RemoveBlock(w http.ResponseWriter, r *http.Request) {
	idBlock = r.FormValue("id_block")

	IDBlock, err := strconv.ParseUint(idBlock, 10, 32)
	if err != nil {
		log.Println("Error parsing shelf ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	exists, err := DB_CheckItemsExistOrNotBlock(IDBlock)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		if err := DB.DB_RemoveBlock(IDBlock); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		msg := map[string]string{"message": "Berhasil Menghapus Block"}
		helper.Response(w, msg, http.StatusOK)
	} else {
		msg := map[string]string{"message": "Tidak Dapat Menghapus Block, Barang Masih Tersedia"}
		helper.Response(w, msg, http.StatusBadRequest)
	}
} 