package models

import (
	"gorm.io/gorm"
	"time"
)

type Shelf struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	NameShelf string `gorm:"type:varchar(10); unique; not null" json:"nama_shelf" validate:"required,max=9"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	CapacityRow float64 `gorm:"type:DECIMAL(10, 0);not null" json:"kapasitas_row" validate:"required,number,max=9"`
	JumlahRows uint `gorm:"-" json:"jumlah_rows"`
	JumlahBlocks uint `gorm:"-" json:"jumlah_block"`
	JumlahBarangs uint `gorm:"-" json:"jumlah_barangs"`
	Rows []Row `gorm:"many2many:shelf_rows;constraint:OnDelete:CASCADE;"`
}