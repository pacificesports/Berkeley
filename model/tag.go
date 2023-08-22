package model

import "time"

type Tag struct {
	SchoolID  string    `gorm:"primaryKey" json:"school_id"`
	Tag       string    `gorm:"primaryKey" json:"tag"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Tag) TableName() string {
	return "school_tag"
}
