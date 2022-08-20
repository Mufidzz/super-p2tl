package postgre

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/lib/pq"
)

func (db *Postgre) GetDataDIL(filter *presentation.FilterParamDIL, pagination presentation.Pagination) (res []presentation.GetDataDILResponse, err error) {
	q := `SELECT id, created_at, idpel, nama, alamat, merk_meter, daya, no_tiang, nama_gardu, tarif, no_kwh, jenis_mk, th_tera_kwh, koordinat_y, koordinat_x from dt_dil`

	// Implement Filter
	q = filterDIL(filter, q)

	// Implement Pagination
	q = fmt.Sprintf("%s ORDER BY id LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetDataTOSO",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetDataDILResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataTOSO",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) GetDataDILCount(filter *presentation.FilterParamDIL) (totalData int64, err error) {
	q := `SELECT count(idpel) FROM dt_dil`

	q = filterDIL(filter, q)

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetDataDILCount",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err = rows.Scan(&totalData)
		if err != nil {
			return 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataTOSOCount",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	return totalData, nil
}

func (db *Postgre) CreateBulkDataDIL(in []presentation.DataDIL) (insertedID []int, err error) {
	q := `INSERT INTO public.dt_dil (idpel, nama, alamat, merk_meter, daya, no_tiang, nama_gardu, tarif, no_kwh,
                           jenis_mk, th_tera_kwh, koordinat_x, koordinat_y) VALUES`
	queryParamLen := 13

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		q = fmt.Sprintf("%s ($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", q, paramCount, paramCount+1, paramCount+2, paramCount+3, paramCount+4, paramCount+5, paramCount+6, paramCount+7, paramCount+8, paramCount+9, paramCount+10, paramCount+11, paramCount+12)
		paramArgs = append(paramArgs, v.Idpel, v.Nama, v.Alamat, v.MerkMeter, v.Daya, v.NoTiang, v.NamaGardu, v.Tarif, v.NoKwh, v.JenisMk, v.ThTeraKwh, v.KoordinatY, v.KoordinatY)
		paramCount += queryParamLen
	}

	// Remove Comma From end of line and Fetch ID after creation
	q = fmt.Sprintf("%s RETURNING id", q[:len(q)-1])

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateBulkDataDIL",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "CreateBulkNewsTags",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		insertedID = append(insertedID, id)
	}

	return insertedID, nil
}

func (db *Postgre) UpdateBulkDataDIL(in []presentation.DataDIL) (updatedID []int, err error) {
	q := `UPDATE dt_dil SET
		  	idpel = new_values.idpel,nama = new_values.nama,alamat = new_values.alamat,merk_meter = new_values.merk_meter,
                  daya = new_values.daya,no_tiang = new_values.no_tiang,nama_gardu = new_values.nama_gardu,tarif = new_values.tarif,
                  no_kwh = new_values.no_kwh,jenis_mk = new_values.jenis_mk,th_tera_kwh = new_values.th_tera_kwh,koordinat_x = new_values.koordinat_x,
                  koordinat_y = new_values.koordinat_y
			FROM (VALUES %s) as new_values (id, idpel, nama, alamat, merk_meter, daya, no_tiang, nama_gardu, tarif, no_kwh,
		   jenis_mk, th_tera_kwh, koordinat_x, koordinat_y) WHERE dt_dil.id = new_values.id RETURNING dt_dil.id`

	queryParamLen := 14

	queryValues := ""
	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		queryValues = fmt.Sprintf("%s ($%d::INTEGER,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),", queryValues, paramCount, paramCount+1, paramCount+2, paramCount+3, paramCount+4, paramCount+5, paramCount+6, paramCount+7, paramCount+8, paramCount+9, paramCount+10, paramCount+11, paramCount+12, paramCount+13)
		paramArgs = append(paramArgs, v.ID, v.Idpel, v.Nama, v.Alamat, v.MerkMeter, v.Daya, v.NoTiang, v.NamaGardu, v.Tarif, v.NoKwh, v.JenisMk, v.ThTeraKwh, v.KoordinatY, v.KoordinatY)
		paramCount += queryParamLen
	}

	q = fmt.Sprintf(q, queryValues[:len(queryValues)-1])

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "UpdateBulkDataDIL",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "UpdateBulkDataDIL",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		updatedID = append(updatedID, id)
	}

	return updatedID, nil
}

func (db *Postgre) DeleteBulkDataDIL(dataIDs []int) error {
	q := `DELETE FROM dt_dil WHERE id = ANY($1)`

	_, err := db.chiefDatabase.Master.Exec(q, pq.Array(dataIDs))
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "DeleteBulkDataDIL",
			Description:  "failed exec queryx",
			Trace:        err,
		}.Error()
	}

	return nil
}
