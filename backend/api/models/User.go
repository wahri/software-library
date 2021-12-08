package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "change_password":
		if u.Password == "" {
			return errors.New("required Password")
		}

		return nil
	case "update":
		if u.Username == "" {
			return errors.New("required Username")
		}
		if u.Password == "" {
			return errors.New("required Password")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("required Password")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("required Username")
		}
		if u.Password == "" {
			return errors.New("required Password")
		}
		return nil
	}
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (p *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Model(&User{}).Create(&p).Error
	if err != nil {
		return &User{}, err
	}
	return p, nil
}

func (p *User) GetAllJadwals(db *gorm.DB) (*[]User, error) {
	Users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Preload("Laporan").Find(&Users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &Users, nil
}
