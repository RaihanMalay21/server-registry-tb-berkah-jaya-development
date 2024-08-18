package models 

import (
	"gorm.io/gorm"
	"time"
)

type Barang struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	Kode string `gorm:"type:varchar(20); not null; unique" json:"kode" validate:"required,max=20"`
	NameBarang string `gorm:"type:varchar(100); not null" json:"nama_barang" validate:"required,max=100"`
	Material string `gorm:"type:varchar(100); not null" json:"material" validate:"required,max=100"`
	Desain string `gorm:"type:varchar(200); not null" json:"desain" validate:"required,max=200"`
	Fitur string `gorm:"type:varchar(200); not null" json:"Fitur" validate:"required,max=200"`
	Image string `gorm:"type:varchar(200); not null" json:"image" validate:"required,max=200"`
	Block []Block `gorm:"many2many:block_barangs"`    
} 