package structs

import "github.com/jinzhu/gorm"

type Kasi struct {
	gorm.Model
	Kode_Kasi string
	Nama_Kasi string
}

func (Kasi) Kasi()string{
	return "kasi"
}