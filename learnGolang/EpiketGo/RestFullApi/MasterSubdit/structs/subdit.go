package structs

import "github.com/jinzhu/gorm"

type Subdit struct {
	gorm.Model
	Kode_Subdirektorat string
	Nama_Subdirektorat string
}

func (Subdit) Subdirektorat() string {
	return "subdit"
}