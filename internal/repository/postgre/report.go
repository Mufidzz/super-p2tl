package postgre

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/dbutils"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"time"
)

func (db *Postgre) CreateFindingReports(input []presentation.CreateFindingReportsRequest) (lastInsertedID int, err error) {
	q := `INSERT INTO dt_temuan (nomor_ba, redaksi_temuan, material_kwh_meter, material_mcb, material_tic_cable, idpel, created_at) VALUES`

	queryParamLen := 7

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range input {
		//seq := make([]interface{}, queryParamLen + 1)
		var seq []interface{}
		seq = append(seq, q)
		seq = append(seq, dbutils.GenerateParamSequence(paramCount, queryParamLen)...)

		fmt.Println(seq)

		q = fmt.Sprintf("%s ($%d::TEXT, $%d::TEXT, $%d::TEXT, $%d::TEXT, $%d::TEXT, $%d::TEXT, $%d),", seq...)
		paramArgs = append(paramArgs, v.NomorBA, v.RedaksiTemuan, v.MaterialKWHMeter, v.MaterialMCB, v.MaterialTICCable, v.IDPEL, time.Now())
		paramCount += queryParamLen
	}

	q = fmt.Sprintf("%s RETURNING id", q[:len(q)-1])

	fmt.Println(q)

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateLogRoomEvent",
			Description:  "failed exec",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err = rows.Scan(&lastInsertedID)
		if err != nil {
			return 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "CreateLogRoomEvent",
				Description:  "failed get rows affected",
				Trace:        err,
			}.Error()
		}
	}

	return lastInsertedID, nil
}

func (db *Postgre) CreatePenormalanReports(input []presentation.CreatePenormalanReportsRequest) (insertedID []int, err error) {
	q := `INSERT INTO dt_penormalan (merk_meter, idpel, type_meter, no_meter, tahun_meter, stand_cabut, stand_pasang, merk_pembatas, rating_pembatas, panjang_sr, no_segel) VALUES`

	queryParamLen := 11

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range input {
		var seq []interface{}
		seq = append(seq, q)
		seq = append(seq, dbutils.GenerateParamSequence(paramCount, queryParamLen)...)

		q = fmt.Sprintf("%s ($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),", seq...)
		//paramCount, paramC*ount+1, paramCount+2, paramCount+3, paramCount+4, paramCount+5, paramCount+6, paramCount+7, paramCount+8, paramCount+9, paramCount+10, paramCount+11)
		paramArgs = append(paramArgs, v.MerkMeter, v.IDPEL, v.TypeMeter, v.NoMeter, v.TahunMeter, v.StandCabut, v.StandPasang, v.MerkPembatas, v.RatingPembatas, v.PanjangSr, v.NoSegel)
		paramCount += queryParamLen
	}

	q = fmt.Sprintf("%s RETURNING id", q[:len(q)-1])

	fmt.Println(q)

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateLogRoomEvent",
			Description:  "failed exec",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t int

		err = rows.Scan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "CreateLogRoomEvent",
				Description:  "failed get rows affected",
				Trace:        err,
			}.Error()
		}

		insertedID = append(insertedID, _t)
	}

	return insertedID, nil
}

func (db *Postgre) GetTemuanReport(filter *presentation.FilterTemuanReport, pagination presentation.Pagination) (res []presentation.GetTemuanReportResponse, err error) {
	q := `SELECT dt_temuan.*, dd.nama
			FROM dt_temuan
			LEFT JOIN dt_dil dd on dt_temuan.idpel = dd.idpel`

	// Implement Filter
	q = filterTemuan(filter, q)

	// Implement Pagination
	q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

	fmt.Println(q)

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetTemuanReport",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetTemuanReportResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetTemuanReport",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) GetTemuanReportCount(filter *presentation.FilterTemuanReport) (totalData int64, err error) {
	q := `SELECT count(idpel) FROM dt_temuan`

	q = filterTemuan(filter, q)

	fmt.Println(q)

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetTemuanReportCount",
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
				FunctionName: "GetTemuanReportCount",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	return totalData, nil
}

func (db *Postgre) GetPenormalanReport(filter *presentation.FilterPenormalanReport, pagination presentation.Pagination) (res []presentation.GetPenormalanReportResponse, err error) {
	q := `SELECT * from dt_penormalan`

	// Implement Filter
	if filter != nil {
		if filter.ID != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "id", int64(filter.ID))
		}

		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "dt_temuan.idpel", filter.IDPEL)
		}
	}

	// Implement Pagination
	q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetPenormalanReport",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetPenormalanReportResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetPenormalanReport",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) GetPenormalanReportCount(filter *presentation.FilterPenormalanReport) (totalData int64, err error) {
	q := `SELECT count(idpel) FROM dt_penormalan`

	if filter != nil {
		if filter.ID != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "id", int64(filter.ID))
		}

		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "idpel", filter.IDPEL)
		}
	}

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetTemuanReportCount",
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
				FunctionName: "GetTemuanReportCount",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	return totalData, nil
}

func (db *Postgre) GetListTemuanMangkrak(pagination presentation.Pagination) (res []presentation.GetListTemuanMangkrakResponse, err error) {
	q := `SELECT created_at, nomor_ba, idpel, sm.nama as status_mangkrak, sb.nama as status_bayar
			FROM dt_temuan
			LEFT JOIN sdt_status_mangkrak sm on sm.id_mangkrak = dt_temuan.status_mangkrak
			LEFT JOIN sdt_status_bayar sb on sb.id_status = dt_temuan.status
			
			WHERE (status = 1 AND status_mangkrak IN (1,2,3) AND
				   EXTRACT(EPOCH FROM (now()::timestamp - COALESCE(last_status_mangkrak_update, created_at))) >= (3 * 24 * 3600))
			OR (status = 1 AND status_mangkrak = 4 AND
				   EXTRACT(EPOCH FROM (now()::timestamp - COALESCE(last_status_mangkrak_update, created_at))) >= (5 * 24 * 3600))
			OR (status = 1 AND status_mangkrak = 5 AND
				   EXTRACT(EPOCH FROM (now()::timestamp - COALESCE(last_status_mangkrak_update, created_at))) >= (6 * 24 * 3600))`

	q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetListTemuanMangkrak",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetListTemuanMangkrakResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetListTemuanMangkrak",
				Description:  "failed running struct scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) GetListTemuanMangkrakCount() (res int, err error) {
	q := `SELECT count(*)
			FROM dt_temuan
			LEFT JOIN sdt_status_mangkrak sm on sm.id_mangkrak = dt_temuan.status_mangkrak
			LEFT JOIN sdt_status_bayar sb on sb.id_status = dt_temuan.status
			
			WHERE (status = 1 AND status_mangkrak IN (1,2,3) AND
				   EXTRACT(EPOCH FROM (now()::timestamp - COALESCE(last_status_mangkrak_update, created_at))) >= (3 * 24 * 3600))
			OR (status = 1 AND status_mangkrak = 4 AND
				   EXTRACT(EPOCH FROM (now()::timestamp - COALESCE(last_status_mangkrak_update, created_at))) >= (5 * 24 * 3600))
			OR (status = 1 AND status_mangkrak = 5 AND
				   EXTRACT(EPOCH FROM (now()::timestamp - COALESCE(last_status_mangkrak_update, created_at))) >= (6 * 24 * 3600))
		   ORDER BY last_status_mangkrak_update DESC`

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetListTemuanMangkrak",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetListTemuanMangkrak",
				Description:  "failed running struct scan",
				Trace:        err,
			}.Error()
		}
	}

	return res, nil
}
