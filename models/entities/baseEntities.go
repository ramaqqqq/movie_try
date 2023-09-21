package entities

import "time"

type BaseTime struct {
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated"`
}
