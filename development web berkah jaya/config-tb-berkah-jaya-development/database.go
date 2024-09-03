package config 

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
)

var (
	DB *gorm.DB
	// custemer models.Custemer
	// barang  models.Barang
	// hadiah models.Hadiah
)

func DB_Connection() {
	db, err := gorm.Open(mysql.Open("root:0987@tcp(localhost:3306)/TB_Berkah_Jaya?parseTime=true"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Barang{})
	db.AutoMigrate(&models.Hadiah{})
	db.AutoMigrate(&models.Pembelian{})
	db.AutoMigrate(&models.Pembelian_Per_Item{})
	db.AutoMigrate(&models.HadiahUser{})

	DB = db
}