package controller

import (
	"net/http"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func Shelfs(w http.ResponseWriter, r *http.Request) {
	shelfs, err := config.GetShelfs()
	if err != nil {
		log.Println("error catn retreaving data shelfs from database:", err)
		helper.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.Response(w, shelfs, http.StatusOK)
}