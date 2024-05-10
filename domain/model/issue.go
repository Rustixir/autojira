package model

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	Summary     string
	Description string
	Type        string // Bug || Feature
	Project     string // PIN || SHOP
}
