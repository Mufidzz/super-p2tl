package presentation

import "time"

type DataDIL struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Idpel      string    `json:"idpel"`
	Nama       string    `json:"nama"`
	Alamat     string    `json:"alamat"`
	MerkMeter  string    `json:"merk_meter"`
	Daya       string    `json:"daya"`
	NoTiang    string    `json:"no_tiang"`
	NamaGardu  string    `json:"nama_gardu"`
	Tarif      string    `json:"tarif"`
	NoKwh      string    `json:"no_kwh"`
	JenisMk    string    `json:"jenis_mk"`
	ThTeraKwh  string    `json:"th_tera_kwh"`
	KoordinatX string    `json:"koordinat_x"`
	KoordinatY string    `json:"koordinat_y"`
}
