package model

import "time"

type Weather struct {
	ID          uint   `gorm:"primaryKey"`
	City        string `gorm:"index"`
	Description string
	Temperature float32
	Timestamp   time.Time
}
