package presentation

import "time"

type GetDataUserTemuanResponse struct {
	CreatedAt    time.Time            `db:"created_at"`
	TemuanDetail []UserTemuanWorkload `db:"temuan_detail"`
}

type UserTemuanWorkload struct {
	ID               int    `json:"id"`
	CreatedAt        string `json:"created_at"`
	NomorBa          string `json:"nomor_ba"`
	JenisTemuan      string `json:"jenis_temuan"`
	PemakaianKWH     int64  `json:"pemakaian_kwh"`
	MaterialKwhMeter string `json:"material_kwh_meter"`
	MaterialMcb      string `json:"material_mcb"`
	MaterialTicCable string `json:"material_tic_cable"`
	IDPEL            string `json:"idpel"`
}

type DataTemuan struct {
	ID               int         `json:"id"`
	CreatedAt        string      `json:"created_at"`
	NomorBa          string      `json:"nomor_ba"`
	JenisTemuan      interface{} `json:"jenis_temuan"`
	MaterialKwhMeter string      `json:"material_kwh_meter"`
	MaterialMcb      string      `json:"material_mcb"`
	MaterialTicCable string      `json:"material_tic_cable"`
	Idpel            string      `json:"idpel"`
	RedaksiTemuan    string      `json:"redaksi_temuan"`
	PemakaianKwh     interface{} `json:"pemakaian_kwh"`
	Status           int         `json:"status"`
	StatusMangkrak   int         `json:"status_mangkrak"`
}

type UpdateBulkJenisTemuanOnTemuanReportRequest struct {
	IDs         []int  `json:"ids"`
	JenisTemuan string `json:"jenis_temuan"`
}

type GetPerformanceKwhReportResponse struct {
	CreatedAt    string `json:"created_at" db:"created_at"`
	PemakaianKwh int64  `json:"pemakaian_kwh" db:"pemakaian_kwh"`
}
