package models

import "time"

type Kamar struct {
	Id          string    `gorm:"type:varchar(20); primaryKey" json:"id"`
	MarkusNicu  string    `gorm:"type:varchar(20)" json:"markus_nicu"`
	MarkusVvip  string    `gorm:"type:varchar(20)" json:"markus_vvip"`
	MarkusVip   string    `gorm:"type:varchar(20)" json:"markus_vip"`
	Lukas       string    `gorm:"type:varchar(20)" json:"maria"`
	Fransiskus  string    `gorm:"type:varchar(20)" json:"fransiskus"`
	Matius      string    `gorm:"type:varchar(20)" json:"matius"`
	Teresia     string    `gorm:"type:varchar(20)" json:"teresia"`
	TeresiaTiga string    `gorm:"type:varchar(20)" json:"teresia_tiga"`
	Yosef       string    `gorm:"type:varchar(20)" json:"yosef"`
	Klara       string    `gorm:"type:varchar(20)" json:"klara"`
	Egidio      string    `gorm:"type:varchar(20)" json:"egidio"`
	Yohanes     string    `gorm:"type:varchar(20)" json:"yohanes"`
	UpdatedAt   time.Time `json:"updated_at"`
}