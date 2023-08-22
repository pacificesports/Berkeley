package model

import "time"

type School struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	IconURL     string    `json:"icon_url"`
	BannerURL   string    `json:"banner_url"`
	Type        string    `json:"type"`
	Address     string    `json:"address"`
	Verified    bool      `json:"verified"`
	Roles       []Tag     `gorm:"-" json:"tags"`
	Privacy     []Email   `gorm:"-" json:"emails"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (School) TableName() string {
	return "school"
}

func (school School) String() string {
	return "(" + school.ID + ")" + " " + school.Name + " - " + school.Type
}
