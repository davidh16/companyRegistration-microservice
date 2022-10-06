package model

import "time"

type Company struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	OIB       int64     `json:"oib" gorm:"unique"`
	Address   string    `json:"address"`
	Deleted   bool      `json:"-" gorm:"default=false"`
	DeletedAt time.Time `json:"-" gorm:"default:"`
}
