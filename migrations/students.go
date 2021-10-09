package migrations

import (
	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	ID   int
	Name string `gorm:"size:255" `
	Age  int
}
