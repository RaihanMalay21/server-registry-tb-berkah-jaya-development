package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
	"gorm.io/gorm"
)

func GetShelfs() (*[]models.Shelf, error) {
	var shelves []models.Shelf
	if err := DB.Preload("Rows").Preload("Rows.Blocks").Preload("Rows.Blocks.Barangs").Find(&shelves).Error; err != nil {
		return nil, err
	}

	// count number of row block and barang
	for index := range shelves {
		var rowsCount uint
		var blocksCount uint
		var barangCount uint

		// count rows
		if err := DB.Model(&models.Row{}).
			Joins("JOIN shelf_rows ON shelf_rows.row_id = rows.id").
			Where("shelf_rows.shelf_id = ?", shelves[index].ID).
			Count(&rowsCount).Error; err != nil {
			return nil, err
		}

		// count block
		if err := DB.Model(&models.Block{}).
			Joins("JOIN row_blocks ON row_blocks.block_id = blocks.id").
			Joins("JOIN rows ON rows.id = row_blocks.row_id").
			Joins("JOIN shelf_rows ON shelf_rows.row_id = rows.id").
			Where("shelf_rows.shelf_id = ?", shelves[index].ID).
			Count(&blocksCount).Error; err != nil {
			return nil, err
		}

		// count barangs
		if err := DB.Model(&models.Barang{}).
			Joins("JOIN block_barangs ON block_barangs.barang_id = barangs.id").
			Joins("JOIN blocks ON blocks.id = block_barangs.block_id").
			Joins("JOIN row_blocks ON row_blocks.block_id = blocks.id").
			Joins("JOIN rows ON rows.id = row_blocks.row_id").
			Joins("JOIN shelf_rows ON shelf_rows.row_id = rows.id").
			Where("shelf_rows.shelf_id = ?", shelves[index].ID).
			Count(&barangCount).Error; err != nil {
			return nil, err
		}

		shelves[index].JumlahRows = rowsCount
		shelves[index].JumlahBlocks = blocksCount
		shelves[index].JumlahBarangs = barangCount
	}

	return &shelves, nil
}