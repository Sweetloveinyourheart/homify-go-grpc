package models

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Country    string
	City       string
	Latitude   float32
	Longitude  float32
	PropertyId uint
}
