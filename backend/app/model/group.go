package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name string `gorm:"type:varchar(64);not null" json:"name"`
	Type string `gorm:"type:varchar(16);unique;not null" json:"type"`
}