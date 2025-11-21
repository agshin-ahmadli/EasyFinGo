package model

import "time"

type Country struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"not null;uniqueIndex" json:"name" validate:"required"`
	CountryCode string `gorm:"not null;uniqueIndex;size:2" json:"country_code" validate:"required,len=2"`
	CallingCode string `gorm:"not null" json:"calling_code" validate:"required"`
	FlagUrl     string `gorm:"not null" json:"flag_url" validate:"required,url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Country) TableName() string {
	return "countries"
}
