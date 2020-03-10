package structs

import "github.com/jinzhu/gorm"

// deklarasi Nama Fielad harus huruf kapital di depan
type Login struct {
	gorm.Model
	Nip string
	Password string
	Name string
	No_hp string
	Kode_subdirekotrat string
	Kode_seksi string
	Aktif int
	Level_pengguna string
	Token string
}

func (Login) Login() string {
	return "login"
}