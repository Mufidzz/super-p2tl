package postgre

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/dbutils"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"time"
)

func (db *Postgre) GetDataTOSO(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOResponse, err error) {
	q := `SELECT * FROM dt_to_so`

	// Implement Filter
	if filter != nil {
		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "idpel", filter.IDPEL)
		}

		if filter.Nyala400 != "" && filter.Nyala400 != "All" {
			q = dbutils.AddStringFilter(q, "AND", "jamnyala400", filter.Nyala400)
		}

		if filter.Nyala600 != "" && filter.Nyala600 != "All" {
			q = dbutils.AddStringFilter(q, "AND", "jamnyala600", filter.Nyala600)
		}
	}

	// Implement Pagination
	q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

	fmt.Println(q)

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
		var _t presentation.GetDataTOSOResponse

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

func (db *Postgre) GetDataTOSOCount(filter *presentation.FilterParamTOSOData) (todayCreation, totalData int64, err error) {
	q := `SELECT count(*) FROM dt_to_so`

	if filter != nil {
		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "idpel", filter.IDPEL)
		}

		if filter.Nyala400 != "" && filter.Nyala400 != "All" {
			q = dbutils.AddStringFilter(q, "AND", "jamnyala400", filter.Nyala400)
		}

		if filter.Nyala600 != "" && filter.Nyala600 != "All" {
			q = dbutils.AddStringFilter(q, "AND", "jamnyala600", filter.Nyala600)
		}
	}

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return 0, 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetDataTOSOCount",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err = rows.Scan(&totalData)
		if err != nil {
			return 0, 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataTOSOCount",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	q = `SELECT count(*) FROM dt_to_so WHERE created_at = $1`

	rows, err = db.chiefDatabase.Master.Queryx(q, time.Now())
	if err != nil {
		return 0, 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetDataTOSOCount",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err = rows.Scan(&todayCreation)
		if err != nil {
			return 0, 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataTOSOCount",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	return todayCreation, totalData, nil
}

func (db *Postgre) GetDataTOSOCore(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error) {
	q := `SELECT id, created_at, idpel, nama, alamat, kdgardu, namagardu, tarif, pemkwh, jamnyala, jamnyala400, jamnyala600, alasan_koreksi  FROM dt_to_so`

	// Implement Filter
	if filter != nil {
		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "idpel", filter.IDPEL)
		}

		if filter.Nyala400 != "" && filter.Nyala400 != "All" {
			q = dbutils.AddStringFilter(q, "AND", "jamnyala400", filter.Nyala400)
		}

		if filter.Nyala600 != "" && filter.Nyala600 != "All" {
			q = dbutils.AddStringFilter(q, "AND", "id", filter.Nyala600)
		}

		if filter.NotAssignedOnly != false {
			q = dbutils.AddCustomFilter(q, "AND", "id", "NOT IN", "(SELECT DISTINCT to_so_id FROM dt_user_to_so)")
		}
	}

	// Implement Pagination
	q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

	fmt.Println(q)

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
		var _t presentation.GetDataTOSOCoreResponse

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
