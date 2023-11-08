package models

import "time"

type Booking struct {
	NoBooking string  `gorm:"type:varchar(50); primaryKey" json:"no_booking"`
	NoKtp     string  `gorm:"type:varchar(50)" json:"no_ktp"`
	Nama      string  `gorm:"type:varchar(250)" json:"nama"`
	Tempat    string  `gorm:"type:varchar(150)" json:"tempat"`
	Tanggal   string  `gorm:"type:varchar(50)" json:"tanggal"`
	Telepon   string  `gorm:"type:varchar(50)" json:"telepon"`
	Alamat    string  `gorm:"type:varchar(250)" json:"alamat"`
	Kelamin   string  `gorm:"type:varchar(250)" json:"kelamin"`
	Agama     string  `gorm:"type:varchar(250)" json:"agama"`
	Ibu       string  `gorm:"type:varchar(250)" json:"ibu"`
	Keluarga  string  `gorm:"type:varchar(250)" json:"keluarga"`
	Status    string  `gorm:"type:varchar(250)" json:"status"`
	Suku      string  `gorm:"type:varchar(250)" json:"suku"`
	Pekerjaan string  `gorm:"type:varchar(250)" json:"pekerjaan"`
	Propinsi  string  `gorm:"type:varchar(250)" json:"propinsi"`
	Kabupaten string  `gorm:"type:varchar(250)" json:"kabupaten"`
	Kecamatan string  `gorm:"type:varchar(250)" json:"kecamatan"`
	Desa      string  `gorm:"type:varchar(250)" json:"desa"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}