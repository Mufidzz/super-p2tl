package presentation

import "time"

const USER_ROLE_PETUGAS = 3

type FilterParamUserData struct {
	Name string `json:"name"`
	Role int    `json:"role"`
}

type GetDataPetugasResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

type GetUserPasswordResponse struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     int    `json:"role" db:"role"`
}

type GetUserWorkloadResponse struct {
	CreatedAt time.Time `json:"created_at"`
}

type UserWorkloadCore struct {
	ID     *int    `json:"id"`
	Idpel  *string `json:"idpel"`
	Nama   *string `json:"nama"`
	Tarif  *string `json:"tarif"`
	Daya   *string `json:"daya"`
	Alamat *string `json:"alamat"`
}
