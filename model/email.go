package model

import "time"

type Email struct {
	SchoolID  string    `gorm:"primaryKey" json:"school_id"`
	Email     string    `gorm:"primaryKey" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Email) TableName() string {
	return "school_email"
}
