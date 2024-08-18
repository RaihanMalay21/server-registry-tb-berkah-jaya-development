package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_AddShelf(model models.Shelf) error {
	if err := DB.Create(&model).Error(); err != nil {
		return err
	}
	return nil
}