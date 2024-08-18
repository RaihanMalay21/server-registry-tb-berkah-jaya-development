package config

import (
	"github.com/RaihanMalay21/web-gudang/models"

	"gorm.io/gorm"
)

func DB_AddBlock(block models.Block, idRow uint) error {
	Row := models.Row{
		ID: idRow,
	}

	if err := DB.Model(&Row).Association("Blocks").Append(&block).Error(); err != nil {
		return err
	}

	return nil
}