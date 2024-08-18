package config 

import (
	"github.com/RaihanMalay21/web-gudang/models"

	"gorm.io/gorm"
)

func DB_AddBarang(tx *gorm.DB, id_Block uint, barang models.Barang) error {
	block := models.Block{
		ID: id_Block,
	}

	if err := tx.Model(&block).Association("barangs").Append(&barang); err != nil {
		return err
	}

	return nil
}