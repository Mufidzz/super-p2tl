package presentation

import (
	"time"
)

type GetDataUserTOSOResponse struct {
	CreatedAt  time.Time          `db:"created_at"`
	ToSoDetail []UserWorkloadCore `db:"to_so_detail"`
}

type FilterParamDIL struct {
	ID    int    `json:"id"`
	IDPEL string `json:"idpel,omitempty"`
	NoKwh string `json:"no_kwh"`
	Nama  string `json:"nama"`
}

type FilterParamTOSOData struct {
	ID              int    `json:"id,omitempty"`
	IDPEL           string `json:"idpel,omitempty"`
	Tarif           string `json:"tarif"`
	Daya            string `json:"daya"`
	Keterangan      string `json:"keterangan"`
	NotAssignedOnly bool   `json:"not_assigned_only,omitempty"`
}

type FilterTemuanReport struct {
	ID              int    `json:"id,omitempty"`
	IDPEL           string `json:"idpel,omitempty"`
	StatusBayar     int    `json:"status_bayar"`
	DateFrom        string `json:"date_from"`
	DateTo          string `json:"date_to"`
	NotAssignedOnly bool   `json:"not_assigned_only,omitempty"`
}

type FilterPenormalanReport struct {
	ID    int    `json:"id,omitempty"`
	IDPEL string `json:"idpel,omitempty"`
}

type GetDataTOSOCoreResponse struct {
	ID         *int       `json:"id" db:"id"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
	IDPEL      *string    `json:"idpel" db:"idpel"`
	Nama       *string    `json:"nama" db:"nama"`
	Alamat     *string    `json:"alamat" db:"alamat"`
	Tarif      *string    `json:"tarif" db:"tarif"`
	Daya       *string    `json:"daya" db:"daya"`
	Keterangan *string    `json:"keterangan" db:"keterangan"`
	NamaGardu  *string    `json:"nama_gardu" db:"nama_gardu"`
	NoTiang    *string    `json:"no_tiang" db:"no_tiang"`
	KoordinatX *string    `json:"koordinat_x" db:"koordinat_x"`
	KoordinatY *string    `json:"koordinat_y" db:"koordinat_y"`
}

type GetDataDILResponse struct {
	ID         *int       `json:"id" db:"id"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
	Idpel      *string    `json:"idpel" db:"idpel"`
	Nama       *string    `json:"nama" db:"nama"`
	Alamat     *string    `json:"alamat" db:"alamat"`
	MerkMeter  *string    `json:"merk_meter" db:"merk_meter"`
	Daya       *string    `json:"daya" db:"daya"`
	NoTiang    *string    `json:"no_tiang" db:"no_tiang"`
	NamaGardu  *string    `json:"nama_gardu" db:"nama_gardu"`
	Tarif      *string    `json:"tarif" db:"tarif"`
	NoKwh      *string    `json:"no_kwh" db:"no_kwh"`
	JenisMk    *string    `json:"jenis_mk" db:"jenis_mk"`
	ThTeraKwh  *string    `json:"th_tera_kwh" db:"th_tera_kwh"`
	KoordinatX *string    `json:"koordinat_x" db:"koordinat_x"`
	KoordinatY *string    `json:"koordinat_y" db:"koordinat_y"`
}

type AssignUserTOSOWorkloadRequest struct {
	UserID int64   `json:"user_id"`
	ToSoID []int64 `json:"to_so_id"`
}

type AssignUserTemuanWorkloadRequest struct {
	UserID   int64   `json:"user_id"`
	TemuanID []int64 `json:"temuan_id"`
}

type UpdateUserTOSORequest struct {
	ID       int       `json:"id"`
	ToSoID   int       `json:"to_so_id"`
	UserID   int       `json:"user_id"`
	FinishAt time.Time `json:"finish_at"`
}

type DataTOSO struct {
	ID         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	Idpel      string `json:"idpel"`
	Keterangan string `json:"keterangan"`
}
