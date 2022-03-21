package presentation

import (
	"time"
)

type FilterParamTOSOData struct {
	IDPEL           string `json:"idpel,omitempty"`
	Nyala600        string `json:"nyala600,omitempty"`
	Nyala400        string `json:"nyala400,omitempty"`
	NotAssignedOnly bool   `json:"not_assigned_only,omitempty"`
}

type GetDataTOSOResponse struct {
	ID             int        `db:"id"`
	CreatedAt      *time.Time `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
	THBLREK        *int64     `db:"thblrek"`
	IDPEL          *int64     `db:"idpel"`
	NAMA           *string    `db:"nama"`
	ALAMAT         *string    `db:"alamat"`
	NOBANG         *string    `db:"nobang"`
	KETNOBANG      *string    `db:"ketnobang"`
	RT             *string    `db:"rt"`
	NODLMRT        *string    `db:"nodlmrt"`
	KETNODLMRT     *string    `db:"ketnodlmrt"`
	RW             *string    `db:"rw"`
	KODEPOS        *string    `db:"kodepos"`
	KDGARDU        *string    `db:"kdgardu"`
	NAMAGARDU      *string    `db:"namagardu"`
	KDDK           *string    `db:"kddk"`
	UNITAP         *string    `db:"unitap"`
	UNITUP         *int64     `db:"unitup"`
	TARIF          *string    `db:"tarif"`
	KDPT           *int64     `db:"kdpt"`
	Kdpt2          *int64     `db:"kdpt_2"`
	DAYA           *float64   `db:"daya"`
	KDPROSESKLP    *string    `db:"kdprosesklp"`
	POSTINGBILLING *int64     `db:"postingbilling"`
	MSG            *int64     `db:"msg"`
	RPPTL          *float64   `db:"rpptl"`
	RPTB           *int64     `db:"rptb"`
	RPPPN          *float64   `db:"rpppn"`
	RPBPJU         *float64   `db:"rpbpju"`
	RPBPTRAFO      *float64   `db:"rpbptrafo"`
	RPSEWATRAFO    *float64   `db:"rpsewatrafo"`
	RPSEWAKAP      *float64   `db:"rpsewakap"`
	RPANGSA        *float64   `db:"rpangsa"`
	RPANGSB        *float64   `db:"rpangsb"`
	RPANGSC        *float64   `db:"rpangsc"`
	RPMAT          *float64   `db:"rpmat"`
	RPPLN          *float64   `db:"rppln"`
	RPTAG          *float64   `db:"rptag"`
	RPBK1          *float64   `db:"rpbk1"`
	RPBK2          *float64   `db:"rpbk2"`
	RPBK3          *float64   `db:"rpbk3"`
	SLALWBP        *float64   `db:"slalwbp"`
	SahlwbpCabut   *float64   `db:"sahlwbp_cabut"`
	SlalwbpPasang  *float64   `db:"slalwbp_pasang"`
	SAHLWBP        *float64   `db:"sahlwbp"`
	SLAWBP         *float64   `db:"slawbp"`
	SahwbpCabut    *float64   `db:"sahwbp_cabut"`
	SlawbpPasang   *float64   `db:"slawbp_pasang"`
	SAHWBP         *float64   `db:"sahwbp"`
	SLAKVARH       *float64   `db:"slakvarh"`
	SahkvarhCabut  *float64   `db:"sahkvarh_cabut"`
	SlakvarhPasang *float64   `db:"slakvarh_pasang"`
	SAHKVARH       *float64   `db:"sahkvarh"`
	PEMKWH         *float64   `db:"pemkwh"`
	JAMNYALA       *float64   `db:"jamnyala"`
	PEMKVARH       *float64   `db:"pemkvarh"`
	KELBKVARH      *float64   `db:"kelbkvarh"`
	DAYAMAKS       *float64   `db:"dayamaks"`
	DayamaxWbp     *float64   `db:"dayamax_wbp"`
	PEMDA          *int64     `db:"pemda"`
	KOGOL          *int64     `db:"kogol"`
	SUBKOGOL       *string    `db:"subkogol"`
	FAKM           *float64   `db:"fakm"`
	FAKMKVARH      *float64   `db:"fakmkvarh"`
	TGLCABUTPASANG *int64     `db:"tglcabutpasang"`
	DLPD           *string    `db:"dlpd"`
	DlpdLm         *string    `db:"dlpd_lm"`
	DlpdFkm        *string    `db:"dlpd_fkm"`
	DlpdKvarh      *string    `db:"dlpd_kvarh"`
	Dlpd3bln       *string    `db:"dlpd_3bln"`
	DlpdJnsmutasi  *string    `db:"dlpd_jnsmutasi"`
	DlpdTglbaca    *string    `db:"dlpd_tglbaca"`
	AlasanKoreksi  *int64     `db:"alasan_koreksi"`
	JAMNYALA600    *string    `db:"jamnyala600"`
	JAMNYALA400    *string    `db:"jamnyala400"`
}

type GetDataTOSOCoreResponse struct {
	ID            int        `db:"id"`
	CreatedAt     *time.Time `db:"created_at"`
	IDPEL         *int64     `db:"idpel"`
	NAMA          *string    `db:"nama"`
	ALAMAT        *string    `db:"alamat"`
	KDGARDU       *string    `db:"kdgardu"`
	NAMAGARDU     *string    `db:"namagardu"`
	TARIF         *string    `db:"tarif"`
	PEMKWH        *float64   `db:"pemkwh"`
	JAMNYALA      *float64   `db:"jamnyala"`
	JAMNYALA600   *string    `db:"jamnyala600"`
	JAMNYALA400   *string    `db:"jamnyala400"`
	AlasanKoreksi *int64     `db:"alasan_koreksi"`
}

type AssignUserTOSOWorkloadRequest struct {
	UserID int64   `json:"user_id"`
	ToSoID []int64 `json:"to_so_id"`
}
