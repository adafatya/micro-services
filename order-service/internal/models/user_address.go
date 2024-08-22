package models

type UserAddress struct {
	ID        int
	UserID    int
	Alamat    string
	KodePos   string
	Kelurahan string
	Kecamatan string
	Kabupaten string
	Provinsi  string
}
