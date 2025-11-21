package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type DocumentType string

const (
	DocumentTypeIdCard   DocumentType = "ID_CARD"
	DocumentTypePassport DocumentType = "PASSPORT"
)

type Document struct {
	ID             string       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID         string       `gorm:"type:uuid;not null;index" json:"user_id"`
	User           User         `gorm:"foreignKey:UserID"`
	DocumentType   DocumentType `gorm:"type:varchar(30);not null" json:"document_type" validate:"required"`
	DocumentNumber string       `gorm:"uniqueIndex;not null" json:"document_number" validate:"required"`
	IssuingCountry string       `gorm:"not null" json:"issuing_country" validate:"required"`
	IssueDate      time.Time    `gorm:"not null" json:"issue_date" validate:"required"`
	ExpiryDate     time.Time    `gorm:"not null" json:"expiry_date" validate:"required"`

	DocumentPhotos []Photo `gorm:"foreignKey:DocumentID;constraints:OnDelete:CASCADE" json:"document_photos,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Document) TableName() string {
	return "documents"
}
func (d *Document) BeforeCreate(tx *gorm.DB) error {
	return d.Validate()
}

func (d *Document) BeforeUpdate(tx *gorm.DB) error {
	return d.Validate()
}

func (d *Document) Validate() error {
	minDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

	if d.IssueDate.Before(minDate) {
		return errors.New("issue date cannot be earlier than 1900-01-01")
	}

	if !d.ExpiryDate.After(d.IssueDate) {
		return errors.New("expiry date must be after issue date")
	}
	return nil
}
