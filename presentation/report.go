package presentation

import "time"

type CreateFindingReportsRequest struct {
	NomorBA          string `json:"nomor_ba" db:"nomor_ba"`
	RedaksiTemuan    string `json:"redaksi_temuan" db:"redaksi_temuan"`
	MaterialKWHMeter string `json:"material_kwh_meter" db:"material_kwh_meter"`
	MaterialMCB      string `json:"material_mcb" db:"material_mcb"`
	MaterialTICCable string `json:"material_tic_cable" db:"material_tic_cable"`
	IDPEL            string `json:"idpel" db:"idpel"`
	TOSOId           int    `json:"toso_id"`
}

type CreatePenormalanReportsRequest struct {
	MerkMeter      string `json:"merk_meter" db:"merk_meter"`
	IDPEL          string `json:"idpel" db:"idpel"`
	TypeMeter      string `json:"type_meter" db:"type_meter"`
	NoMeter        string `json:"no_meter" db:"no_meter"`
	TahunMeter     string `json:"tahun_meter" db:"tahun_meter"`
	StandCabut     string `json:"stand_cabut" db:"stand_cabut"`
	StandPasang    string `json:"stand_pasang" db:"stand_pasang"`
	MerkPembatas   string `json:"merk_pembatas" db:"merk_pembatas"`
	RatingPembatas string `json:"rating_pembatas" db:"rating_pembatas"`
	PanjangSr      string `json:"panjang_sr" db:"panjang_sr"`
	NoSegel        string `json:"no_segel" db:"no_segel"`
}

type GetTemuanReportResponse struct {
	ID                       *int       `json:"id" db:"id"`
	CreatedAt                *string    `json:"created_at" db:"created_at"`
	NomorBa                  *string    `json:"nomor_ba" db:"nomor_ba"`
	JenisTemuan              *string    `json:"jenis_temuan" db:"jenis_temuan"`
	MaterialKwhMeter         *string    `json:"material_kwh_meter" db:"material_kwh_meter"`
	MaterialMcb              *string    `json:"material_mcb" db:"material_mcb"`
	MaterialTicCable         *string    `json:"material_tic_cable" db:"material_tic_cable"`
	Idpel                    *string    `json:"idpel" db:"idpel"`
	RedaksiTemuan            *string    `json:"redaksi_temuan" db:"redaksi_temuan"`
	PemakaianKwh             *string    `json:"pemakaian_kwh" db:"pemakaian_kwh"`
	Status                   *int       `json:"status" db:"status"`
	StatusMangkrak           *int       `json:"status_mangkrak" db:"status_mangkrak"`
	Nama                     *string    `json:"nama" db:"nama"`
	LastStatusMangkrakUpdate *time.Time `json:"last_status_mangkrak_update" db:"last_status_mangkrak_update"`
}

type GetPenormalanReportResponse struct {
	ID             *int    `json:"id" db:"id"`
	CreatedAt      *string `json:"created_at" db:"created_at"`
	MerkMeter      *string `json:"merk_meter" db:"merk_meter"`
	Idpel          *string `json:"idpel" db:"idpel"`
	TypeMeter      *string `json:"type_meter" db:"type_meter"`
	NoMeter        *string `json:"no_meter" db:"no_meter"`
	TahunMeter     *string `json:"tahun_meter" db:"tahun_meter"`
	StandCabut     *string `json:"stand_cabut" db:"stand_cabut"`
	StandPasang    *string `json:"stand_pasang" db:"stand_pasang"`
	MerkPembatas   *string `json:"merk_pembatas" db:"merk_pembatas"`
	RatingPembatas *string `json:"rating_pembatas" db:"rating_pembatas"`
	PanjangSr      *string `json:"panjang_sr" db:"panjang_sr"`
	NoSegel        *string `json:"no_segel" db:"no_segel"`
}

type GetListTemuanMangkrakResponse struct {
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	NomorBA        string    `json:"nomor_ba" db:"nomor_ba"`
	IDPEL          string    `json:"idpel" db:"idpel"`
	StatusMangkrak string    `json:"status_mangkrak" db:"status_mangkrak"`
	StatusBayar    string    `json:"status_bayar" db:"status_bayar"`
}
