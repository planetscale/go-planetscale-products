package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string          `gorm:"size:255;not null;" json:"name" binding:"required"`
	SerialNumber string          `gorm:"size:255;not null;unique" json:"serialNumber" binding:"required"`
	Quantity     uint            `json:"quantity" binding:"required"`
	Price        decimal.Decimal `gorm:"type:numeric" json:"price" binding:"required"`
	Description  string          `gorm:"type:text" json:"description" binding:"required"`
}
