package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"github.com/RaihanMalay21/web-gudang/models"
)

func AddBarang(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile{
			log.Println(err.Error())
			msg := map[string]string{"message": "Tidak ada file yang di unggah"}
			helper.Response(w, msg, http.StatusBadRequest)
			return
		}
		log.Println(err)
		msg := map[string]string{"Message": "Image Gagal Di Input"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// mengambil ext dari nama file
	ext := filepath.Ext(handler.Filename)
	if ext == "" || (ext != ".jpg" && ext != ".png" && ext != ".gif") {
		log.Println("Tipe Gambar harus jpg, png, dan gift")
		msg := map[string]string{"message": "Tipe Gambar harus jpg, png, dan gift"}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	// size image 
	fileSize := handler.Size

	// authentikasi ukuran file 
	if fileSize > 2000000 {
		log.Println("error on line 61 function input barang : Ukuran FIle terlalu besar")
		message := map[string]string{"message":"Ukuran Image Terlalu Besar, max 2MB"}
		helper.Response(w, message, http.StatusBadRequest)
		return
	}

	// mengambil nama filenya 
	nameOnly := filepath.Base(handler.Filename[:len(handler.Filename) - len(ext)])
	
	// menkonversi nama file menggunakan sha256 menjadi byte dan ubah menjadi string
	hasher := sha256.Sum256([]byte(nameOnly))
	namaFileStringByte := hex.EncodeToString(hasher[:])

	idBlock, err := strconv.ParseUint(r.FormValue("id_block"), 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	barang := models.Barang{
		Kode: r.FormValue("kode"),
		NameBarang: r.FormValue("nama_barang"),
		Material: r.FormValue("material"),
		Desain: r.FormValue("desain"),
		Fitur: r.FormValue("fitur"),
		Image: namaFileStringByte + ext,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	Trans := helper.TranslatorIDN()

	if err := validate.Struct(&barang); err != nil {
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

	// cek apakah nama image sudah tersedia di database
	var count uint
	if err := config.DB.Model(&models.Barang{}).Where("image = ?", barang.Image).Count(&count).Error; err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count > 0 {
		barang.Image = namaFileStringByte + count + ext
	} 

	tx := config.DB.Transaction()

	if err := config.DB_AddBarang(tx, idBlock, barang); err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileDir := helper.DestinationFolder("C:\\Users\\raiha\\Documents\\web-manage-gudang\\image", barang.Image)

	outFile, err := os.Create(fileDir)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if _, err := io.Copy(outFile, file); err != nil {
		tx.Rollback()
		// menghapus file yang baru saja di buat
		if err := os.Remove(fileDir); err != nil {
			log.Println(err.Error())
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := map[string]string{"message": "Berhasil Menambahkan Barang"}
	helper.Response(w, msg, http.StatusOK)
}