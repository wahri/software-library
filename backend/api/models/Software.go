package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Software struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name           string    `gorm:"size:255;not null" json:"name"`
	ZipFile        string    `gorm:"size:255;not null" json:"ZipFile"`
	LinkSource     string    `gorm:"size:255;not null" json:"LinkSource"`
	LinkPreview    string    `gorm:"size:255;not null" json:"LinkPreview"`
	LinkTutorial   string    `gorm:"size:255;not null" json:"LinkTutorial"`
	License        string    `gorm:"size:255;not null" json:"License"`
	Description    string    `gorm:"size:255;not null" json:"Description"`
	PreviewImage   string    `gorm:"size:255;not null" json:"PreviewImage"`
	ProductVersion float64   `gorm:"not null" json:"ProductVersion"`
	ReleaseDate    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ReleaseDate"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Software) SaveSoftware(db *gorm.DB) (*Software, error) {
	err := db.Debug().Model(&Software{}).Create(&p).Error
	if err != nil {
		return &Software{}, err
	}
	return p, nil
}
