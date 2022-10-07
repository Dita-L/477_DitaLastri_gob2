package models

import "time"

type Order struct {
	Id            uint   `gorm:"primaryKey;autoIncrement"`
	Costumer_Name string `gorm:"not null"`
	Items         []Item
	Ordered_At    time.Time
}
