package presentation

const USER_ROLE_PETUGAS = 3

type FilterParamUserData struct {
	Name string `json:"name"`
	Role int    `json:"role"`
}

type GetDataPetugasResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}
