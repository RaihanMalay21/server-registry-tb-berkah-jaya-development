package config

import (
	"github.com/RaihanMalay21/web-manage-gudang/models"

	"gorm.io/gorm"
)

func DB_RemoveRow(ID_Row uint) error {
	if err := DB.Where("id = ?", ID_Row).Delete(&models.Row{}).Error; err != nil {
		return err
	}
	return nil
}

// cek apakah barang masih tersida di dalam shelf rows block
func DB_CheckItemsExistOrNotRow(ID_Row uint) (bool, error) {
	var count int64

	if err := DB.Model(&models.Barang{}).
		Joins("JOIN block_barangs ON block_barangs.barang_id = barang.id").
		Joins("JOIN block ON block.id = block_barangs.block_id").
		Joins("JOIN row_blocks ON row_blocks.block_id = block.id").
		Where("row_blocks.row_id = ?", ID_Row).
		Count(&count).Error; err != nil {
		return false, err
	}

	// Jika ada satu atau lebih record yang ditemukan
	if count > 0 {
		return false, nil
	}

	// Jika tidak ada record yang ditemukan
	return true, nil
}