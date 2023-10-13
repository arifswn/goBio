package model

import (
	"time"
)

type Users struct {
	Email     string     `json:"email" form:"email" gorm:"primaryKey;not null;size:50"`
	Nama      string     `json:"nama" form:"nama" gorm:"not null;size:50"`
	NoHp      string     `json:"no_hp" form:"no_hp" gorm:"not null;size:15"`
	Alamat    string     `json:"alamat" form:"alamat" gorm:"not null;size:100"` //size 100 berarti maksimal 100 karakter
	Ktp       string     `json:"ktp" form:"ktp" gorm:"not null;size:100"`       // KTP berupa nama file yang akan disimpan di folder uploads
	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"null"`
	DeletedAt *time.Time `json:"deleted_at"`
}
