package models

import (
	"gorm.io/gorm"
	"time"
)

type Row struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID:"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	NumberRow uint `gorm:"not null" json:"nama_block" validate:"required,number,max=9"`
	CapacityBlock  float64 `gorm:"type:DECIMAL(10, 0);not null" json:"kapasitas" validate:"required,number,max=9"`
	Shelf []Shelf `gorm:"many2many:shelf_rows"`
	Blocks []Block `gorm:"many2many:row_blocks;constraint:OnDelete:CASCADE;"`
}