package postgre

import (
	"encoding/json"
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/lib/pq"
	"time"
)

func (db *Postgre) GetDataTOSO(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error) {
	q := `SELECT dts.id, dts.created_at, dts.idpel, dd.nama, dd.alamat, dd.tarif, dd.daya, dts.keterangan, dd.nama_gardu, dd.no_tiang, dd.koordinat_x, dd.koordinat_y FROM dt_to_so dts
			LEFT JOIN dt_dil dd on dts.idpel = dd.idpel`

	// Implement Filter
	q = filterTOSO(filter, q)

	// Implement Pagination
	q = fmt.Sprintf("%s ORDER BY dts.id LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)

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

func (db *Postgre) GetDataTOSOCount(filter *presentation.FilterParamTOSOData) (todayCreation, totalData int64, err error) {
	q := `SELECT count(*) FROM dt_to_so dts`

	q = filterTOSO(filter, q)

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

	q = `SELECT count(*) FROM dt_to_so dts WHERE created_at = $1`

	q = filterTOSO(filter, q)

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

// Deprecated, use GetDataTOSO instead
//func (db *Postgre) GetDataTOSOCore(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error) {
//	q := `SELECT id, created_at, idpel, nama, alamat, kdgardu, namagardu, tarif, pemkwh, jamnyala, jamnyala400, jamnyala600, alasan_koreksi  FROM dt_to_so`
//
//	// Implement Filter
//	if filter != nil {
//		if filter.IDPEL != "" && filter.IDPEL != "All" {
//			q = dbutils.AddStringFilter(q, "AND", "idpel", filter.IDPEL)
//		}
//
//		if filter.Nyala400 != "" && filter.Nyala400 != "All" {
//			q = dbutils.AddStringFilter(q, "AND", "jamnyala400", filter.Nyala400)
//		}
//
//		if filter.Nyala600 != "" && filter.Nyala600 != "All" {
//			q = dbutils.AddStringFilter(q, "AND", "id", filter.Nyala600)
//		}
//
//		if filter.NotAssignedOnly != false {
//			q = dbutils.AddCustomFilter(q, "AND", "id", "NOT IN", "(SELECT DISTINCT to_so_id FROM dt_user_to_so)")
//		}
//	}
//
//	// Implement Pagination
//	q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, pagination.Count, pagination.Offset)
//
//	fmt.Println(q)
//
//	rows, err := db.chiefDatabase.Master.Queryx(q)
//	if err != nil {
//		return nil, response.InternalError{
//			Type:         "Repo",
//			Name:         "Postgre",
//			FunctionName: "GetDataTOSO",
//			Description:  "failed running queryx",
//			Trace:        err,
//		}.Error()
//	}
//
//	for rows.Next() {
//		var _t presentation.GetDataTOSOCoreResponse
//
//		err = rows.StructScan(&_t)
//		if err != nil {
//			return nil, response.InternalError{
//				Type:         "Repo",
//				Name:         "Postgre",
//				FunctionName: "GetDataTOSO",
//				Description:  "failed scan",
//				Trace:        err,
//			}.Error()
//		}
//
//		res = append(res, _t)
//	}
//
//	return res, nil
//}

func (db *Postgre) GetDataUserTOSO(userId int) (res []presentation.GetDataUserTOSOResponse, err error) {
	//q := `SELECT duts.created_at::DATE, json_agg(json_build_object(
	//			'id', duts.id,
	//			'idpel', dts.idpel,
	//			'nama', dd.nama,
	//			'tarif', dd.tarif,
	//			'daya', dd.daya,
	//			'alamat', dd.alamat
	//			)) as to_so_detail FROM dt_user_to_so duts
	//		LEFT JOIN dt_to_so dts on duts.to_so_id = dts.id
	//		LEFT JOIN dt_dil dd on dd.idpel = dts.idpel
	//		WHERE duts.user_id = $1
	//		GROUP BY duts.created_at::DATE
	//		ORDER BY duts.created_at::DATE DESC`

	q := `SELECT duts.created_at::DATE, json_agg(json_build_object(
				'id', duts.id,
				'idpel', dts.idpel,
				'nama', dd.nama,
				'tarif', dd.tarif,
				'daya', dd.daya,
				'alamat', dd.alamat
				)) as to_so_detail FROM dt_user_to_so duts
			LEFT JOIN dt_to_so dts on duts.to_so_id = dts.id
			LEFT JOIN dt_dil dd on dd.idpel = dts.idpel
			WHERE duts.user_id = $1 AND duts.finish_at IS NULL
			GROUP BY duts.created_at::DATE
            ORDER BY duts.created_at::DATE DESC`

	rows, err := db.chiefDatabase.Master.Queryx(q, userId)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetDataUserTOSO",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		sli, err := rows.SliceScan()
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataUserTOSO",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		var _t []presentation.UserWorkloadCore
		err = json.Unmarshal([]byte(fmt.Sprintf("%s", sli[1])), &_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataUserTOSO",
				Description:  "failed unmarshall json",
				Trace:        err,
			}.Error()
		}

		res = append(res, presentation.GetDataUserTOSOResponse{
			CreatedAt:  sli[0].(time.Time),
			ToSoDetail: _t,
		})
	}

	return res, nil
}

func (db *Postgre) UpdateUserTOSO(newUserTOSO presentation.UpdateUserTOSORequest) error {
	//q := `UPDATE dt_user_to_so SET to_so_id = $1, user_id = $2, finish_at = $3 WHERE id = $4`
	q := `UPDATE dt_user_to_so SET`

	paramCount := 1
	paramArgs := []interface{}{}

	if newUserTOSO.ToSoID != 0 {
		q = fmt.Sprintf("%s %s = $%d::INTEGER,", q, "to_so_id", paramCount)
		paramCount++
		paramArgs = append(paramArgs, newUserTOSO.ToSoID)
	}

	if newUserTOSO.UserID != 0 {
		q = fmt.Sprintf("%s %s = $%d::INTEGER,", q, "user_id", paramCount)
		paramCount++
		paramArgs = append(paramArgs, newUserTOSO.UserID)
	}

	if !newUserTOSO.FinishAt.IsZero() {
		q = fmt.Sprintf("%s %s = $%d::TIMESTAMP,", q, "finish_at", paramCount)
		paramCount++
		paramArgs = append(paramArgs, newUserTOSO.FinishAt)
	}

	q = fmt.Sprintf("%s WHERE id = $%d", q[:len(q)-1], paramCount)
	paramCount++
	paramArgs = append(paramArgs, newUserTOSO.ID)

	fmt.Println(q)

	_, err := db.chiefDatabase.Master.Exec(q, paramArgs...)

	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "UpdateUserTOSO",
			Description:  "failed running Exec",
			Trace:        err,
		}.Error()
	}

	return nil
}

func (db *Postgre) CreateBulkDataTOSO(in []presentation.DataTOSO) (insertedID []int, err error) {
	q := `INSERT INTO public.dt_to_so (idpel, keterangan) VALUES`
	queryParamLen := 2

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		q = fmt.Sprintf("%s ($%d,$%d),", q, paramCount, paramCount+1)
		paramArgs = append(paramArgs, v.Idpel, v.Keterangan)
		paramCount += queryParamLen
	}

	// Remove Comma From end of line and Fetch ID after creation
	q = fmt.Sprintf("%s RETURNING id", q[:len(q)-1])

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "CreateBulkDataTOSO",
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

func (db *Postgre) UpdateBulkDataTOSO(in []presentation.DataTOSO) (updatedID []int, err error) {
	q := `UPDATE dt_to_so SET
		  	idpel = new_values.idpel, keterangan = new_values.keterangan
			FROM (VALUES %s) as new_values (id, idpel, keterangan) WHERE dt_to_so.id = new_values.id RETURNING dt_to_so.id`

	queryParamLen := 3

	queryValues := ""
	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		queryValues = fmt.Sprintf("%s ($%d::INTEGER,$%d,$%d),", queryValues, paramCount, paramCount+1, paramCount+2)
		paramArgs = append(paramArgs, v.ID, v.Idpel, v.Keterangan)
		paramCount += queryParamLen
	}

	q = fmt.Sprintf(q, queryValues[:len(queryValues)-1])

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "UpdateBulkDataTOSO",
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
				FunctionName: "UpdateBulkDataTOSO",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		updatedID = append(updatedID, id)
	}

	return updatedID, nil
}

func (db *Postgre) DeleteBulkDataTOSO(dataIDs []int) error {
	q := `DELETE FROM dt_to_so WHERE id = ANY($1)`

	_, err := db.chiefDatabase.Master.Exec(q, pq.Array(dataIDs))
	if err != nil {
		return response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "DeleteBulkDataTOSO",
			Description:  "failed exec queryx",
			Trace:        err,
		}.Error()
	}

	return nil
}
