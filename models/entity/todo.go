package entity

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Task      string    `gorm:"type:varchar(255);not null" json:"task"`
	Status    bool      `gorm:"default:false" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}