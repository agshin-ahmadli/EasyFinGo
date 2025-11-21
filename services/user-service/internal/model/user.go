package model

import (
	"gorm.io/gorm"
	"time"
)

type RegistrationStatus string

const (
	StatusStarted       RegistrationStatus = "STARTED"
	StatusEmailVerified RegistrationStatus = "EMAIL_VERIFIED"
	StatusBasicInfo     RegistrationStatus = "BASIC_INFO"
	StatusDocuments     RegistrationStatus = "DOCUMENTS_UPLOADED"
	StatusPasscodeSet   RegistrationStatus = "PASSCODE_SET"
	StatusCompleted     RegistrationStatus = "COMPLETED"
)

type User struct {
	ID                 string             `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FirstName          string             `gorm:"not null" json:"first_name" validate:"required"`
	MiddleName         string             `json:"middle_name,omitempty"`
	LastName           string             `gorm:"not null" json:"last_name" validate:"required"`
	DateOfBirth        time.Time          `gorm:"not null" json:"date_of_birth" validate:"required"`
	Pesel              string             `gorm:"uniqueIndex;not null" json:"pesel" validate:"required"`
	PhoneNumber        string             `gorm:"uniqueIndex;not null" json:"phone_number" validate:"required"`
	Email              string             `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Password           string             `gorm:"not null" json:"-"`
	Passcode           string             `json:"-"`
	Sub                string             `gorm:"type:uuid" json:"sub"`
	RegistrationStatus RegistrationStatus `gorm:"type:varchar(50);default:'STARTED'" json:"registration_status"`
	Version            int                `gorm:"default:0" json:"version"`

	Addresses   []Address  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"addresses,omitempty"`
	Documents   []Document `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"documents,omitempty"`
	PhotoSelfie *Photo     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"photo_selfie,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "clients"
}
