package model

import "time"

type Media struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	URL       string    `gorm:"not null" json:"url" validate:"required,url"`
	CreatedAt time.Time `json:"createdAt"`
}

type PhotoType string

const (
	PhotoTypeFront  PhotoType = "FRONT"
	PhotoTypeBack   PhotoType = "BACK"
	PhotoTypeSelfie PhotoType = "SELFIE"
)

type Photo struct {
	Media
	PhotoType        PhotoType `gorm:"type:varchar(20);not null" json:"photo_type" validate:"required"`
	OriginalFileNAme string    `gorm:"column:original_filename" json:"original_filename,omitempty"`
	ContentType      string    `gorm:"column:content_type" json:"content_type,omitempty"`
	FileSize         int64     `gorm:"column:file_size" json:"file_size,omitempty"`

	UserID     *string `gorm:"column:client_id;type:uuid;index" json:"user_id,omitempty"`
	DocumentID *string `gorm:"column:document_id;type:uuid;index" json:"document_id"`

	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Photo) TableName() string {
	return "photos"
}
