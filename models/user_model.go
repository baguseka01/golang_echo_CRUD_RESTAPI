package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:name"`
	Location string `gorm:"column:location"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
}
