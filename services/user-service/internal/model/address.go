package model

import "time"

type AddressType string

const (
	AddressTypeHome      AddressType = "HOME"
	AddressTypeWork      AddressType = "WORK"
	AddressTypeMailing   AddressType = "MAILING"
	AddressTypeTemporary AddressType = "TEMPORARY"
)

type Address struct {
	ID              string      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID          string      `gorm:"type:uuid;not null;index"`
	AddressType     AddressType `gorm:"type:varchar(20);not null" json:"address_type" validate:"required"`
	Postcode        string      `json:"postcode,omitempty"`
	StreetName      string      `gorm:"not null" json:"street_name" validate:"required"`
	BuildingNumber  string      `gorm:"not null" json:"building_number" validate:"required"`
	ApartmentNumber string      `json:"apartment_number,omitempty"`
	City            string      `gorm:"not null" json:"city" validate:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Address) TableName() string {
	return "addresses"
}
