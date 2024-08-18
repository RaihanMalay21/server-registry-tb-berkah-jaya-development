package config

import (
	"github.com/RaihanMalay21/web-manage-gudang/models"

	"gorm.io/gorm"
)

func DB_AddRows(idShelf uint, row models.Row) error {
	// Temukan Shelf berdasarkan ID
	var shelf models.Shelf
	if err := DB.Where("ID = ?", idShelf).First(&shelf).Error(); err != nil {
		return err
	}

	// Tambahkan Row ke Shelf
	if err := DB.Model(&shelf).Association("Rows").Append(&row); err != nil {
		return err
	}

	return nil
}