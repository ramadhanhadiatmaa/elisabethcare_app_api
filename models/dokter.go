package models

type Dokter struct {
	KodeDokter string `gorm:"type:varchar(50); primaryKey" json:"kode_dokter"`
	Nama       string `gorm:"type:varchar(100)" json:"nama"`
	Spesialis  string `gorm:"type:varchar(100)" json:"spesialis"`
	Jenis      string `gorm:"type:varchar(50)" json:"jenis"`
	Foto       string `gorm:"type:varchar(150)" json:"foto"`
	Tentang    string `gorm:"type:varchar(500)" json:"tentang"`
	Senin      string `gorm:"type:varchar(150)" json:"senin"`
	Selasa     string `gorm:"type:varchar(150)" json:"selasa"`
	Rabu       string `gorm:"type:varchar(150)" json:"rabu"`
	Kamis      string `gorm:"type:varchar(150)" json:"kamis"`
	Jumat      string `gorm:"type:varchar(150)" json:"jumat"`
	Sabtu      string `gorm:"type:varchar(150)" json:"sabtu"`
	Minggu     string `gorm:"type:varchar(150)" json:"minggu"`
}