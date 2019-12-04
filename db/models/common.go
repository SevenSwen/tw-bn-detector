package models

import "github.com/jinzhu/gorm"

type Currency struct {
	gorm.Model
	Name   string  `json:"name" gorm:"primary_key"`
//	Price  float64 `json:"price"`
//	Weight float64 `json:"weight"` // Вес валюты - общий процент роста/спада валюты за какой-либо промежуток времени
//	Type   int     `json:"type"`   // Тип, соответствует тому сколько дней учавствует в анализе (например 3, 2, 1)
}

