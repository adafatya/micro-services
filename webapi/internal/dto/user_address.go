package dto

type AddUserAddressRequest struct {
	UserID    int    `json:"-"`
	Alamat    string `json:"alamat"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`
	KodePos   string `json:"kode_pos"`
}

type AddUserAddressResponse struct {
	Message       string `json:"message"`
	ID            int    `json:"id"`
	AlamatLengkap string `json:"alamat_lengkap"`
}

type UserAddress struct {
	ID            int    `json:"id"`
	AlamatLengkap string `json:"alamat_lengkap"`
}

type GetUserAddressesResponse struct {
	Message       string        `json:"message"`
	UserAddresses []UserAddress `json:"user_addresses"`
}
