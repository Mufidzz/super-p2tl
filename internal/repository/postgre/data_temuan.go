package postgre

import (
	"encoding/json"
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"time"
)

func (db *Postgre) GetDataUserTemuan(userId int) (res []presentation.GetDataUserTemuanResponse, err error) {
	q := `SELECT
			dut.created_at::DATE, json_agg(
			    json_build_object(
			        'id', dt.id,
			        'created_at', dt.created_at,
			        'nomor_ba' , COALESCE(dt.nomor_ba, 'N/A'),
			        'jenis_temuan' , COALESCE(dt.jenis_temuan, 'N/A'),
			        'material_kwh_meter' , COALESCE(dt.material_kwh_meter, 'N/A'),
			        'material_mcb' , COALESCE(dt.material_mcb, 'N/A'),
			        'material_tic_cable' , COALESCE(dt.material_tic_cable, 'N/A'),
			        'idpel' , COALESCE(dt.idpel, 'N/A'),
			        'pemakaian_kwh' , COALESCE(dt.pemakaian_kwh, 0)
			        )
			    ) as temuan_detail
		FROM
			 dt_user_temuan dut
        LEFT JOIN dt_temuan dt on dut.temuan_id = dt.id
		LEFT JOIN dt_dil dd on dt.idpel = dd.idpel
		WHERE user_id = $1 AND dt.finish_at IS NULL 
		GROUP BY dut.created_at::DATE
		ORDER BY dut.created_at::DATE DESC;`

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

		var _t []presentation.UserTemuanWorkload
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

		res = append(res, presentation.GetDataUserTemuanResponse{
			CreatedAt:    sli[0].(time.Time),
			TemuanDetail: _t,
		})
	}

	return res, nil
}

func (db *Postgre) UpdateDataTemuanReport(in []presentation.DataTemuan) (updatedID []int, err error) {
	q := `UPDATE dt_temuan SET
		  	nomor_ba = new_values.nomor_ba,jenis_temuan = new_values.jenis_temuan,material_kwh_meter = new_values.material_kwh_meter,material_mcb = new_values.material_mcb,material_tic_cable = new_values.material_tic_cable,idpel = new_values.idpel,redaksi_temuan = new_values.redaksi_temuan,pemakaian_kwh = new_values.pemakaian_kwh,status = new_values.status,status_mangkrak = new_values.status_mangkrak
			FROM (VALUES %s) as new_values (id, nomor_ba, jenis_temuan, material_kwh_meter, material_mcb, material_tic_cable, idpel, redaksi_temuan, pemakaian_kwh, status, status_mangkrak) WHERE dt_temuan.id = new_values.id RETURNING dt_temuan.id`

	queryParamLen := 11

	queryValues := ""
	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in {
		queryValues = fmt.Sprintf("%s ($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,),", queryValues, paramCount, paramCount+1, paramCount+2, paramCount+3, paramCount+4, paramCount+5, paramCount+6, paramCount+7, paramCount+8, paramCount+9, paramCount+10)
		paramArgs = append(paramArgs, v.ID, v.NomorBa, v.JenisTemuan, v.MaterialKwhMeter, v.MaterialMcb, v.MaterialTicCable,
			v.Idpel, v.RedaksiTemuan, v.PemakaianKwh, v.Status, v.StatusMangkrak)
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

func (db *Postgre) UpdateBulkJenisTemuanOnTemuanReport(in presentation.UpdateBulkJenisTemuanOnTemuanReportRequest) (updatedID []int, err error) {
	q := `UPDATE dt_temuan SET
		  	jenis_temuan = new_values.jenis_temuan
			FROM (VALUES %s) as new_values (id, jenis_temuan) WHERE dt_temuan.id = new_values.id RETURNING dt_temuan.id`

	queryParamLen := 2

	queryValues := ""
	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range in.IDs {
		queryValues = fmt.Sprintf("%s ($%d::INTEGER, $%d),", queryValues, paramCount, paramCount+1)
		paramArgs = append(paramArgs, v, in.JenisTemuan)
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

func (db *Postgre) GetPerformanceKwhReport() (res []presentation.GetPerformanceKwhReportResponse, err error) {
	q := `SELECT DATE_TRUNC('month',created_at::DATE) AS created_at, SUM(pemakaian_kwh::bigint) as pemakaian_kwh from dt_temuan
			GROUP BY DATE_TRUNC('month',created_at::DATE)`

	rows, err := db.chiefDatabase.Master.Queryx(q)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetPerformanceKwhReport",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetPerformanceKwhReportResponse
		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetPerformanceKwhReport",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}
