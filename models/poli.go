package models

type Poli struct {
	KodePoli string `gorm:"type:varchar(50); primaryKey" json:"kode_poli"`
	Dokter   string `gorm:"type:varchar(250)" json:"dokter"`
	Nama     string `gorm:"type:varchar(150)" json:"nama"`
	Foto     string `gorm:"type:varchar(250)" json:"foto"`
	Status   string `gorm:"type:varchar(50)" json:"status"`
}