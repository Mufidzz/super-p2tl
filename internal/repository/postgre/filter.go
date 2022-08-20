package postgre

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/dbutils"
	"github.com/SuperP2TL/Backend/presentation"
)

func filterDIL(filter *presentation.FilterParamDIL, q string) string {
	if filter != nil {
		if filter.ID != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "id", int64(filter.ID))
		}

		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "idpel", filter.IDPEL)
		}

		if filter.NoKwh != "" && filter.NoKwh != "All" {
			q = dbutils.AddStringFilter(q, "AND", "no_kwh", filter.NoKwh)
		}

		if filter.Nama != "" && filter.Nama != "All" {
			q = dbutils.AddStringFilter(q, "AND", "nama", filter.Nama)
		}
	}

	return q
}

func filterTOSO(filter *presentation.FilterParamTOSOData, q string) string {
	if filter != nil {
		if filter.ID != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "dts.id", int64(filter.ID))
		}

		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "dts.idpel", filter.IDPEL)
		}

		if filter.Tarif != "" && filter.Tarif != "All" {
			q = dbutils.AddStringFilter(q, "AND", "dd.tarif", filter.Tarif)
		}

		if filter.Daya != "" && filter.Daya != "All" {
			q = dbutils.AddStringFilter(q, "AND", "dd.daya", filter.Daya)
		}

		if filter.Keterangan != "" && filter.Keterangan != "All" {
			q = dbutils.AddStringFilter(q, "AND", "dd.keterangan", filter.Keterangan)
		}

		if filter.NotAssignedOnly {
			q = dbutils.AddCustomFilter(q, "AND", "NOT (dts.id", "=", fmt.Sprintf("ANY(%s))", "(SELECT DISTINCT to_so_id FROM dt_user_to_so)"))
		}
	}

	return q
}

func filterTemuan(filter *presentation.FilterTemuanReport, q string) string {
	if filter != nil {
		if filter.ID != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "id", int64(filter.ID))
		}

		if filter.IDPEL != "" && filter.IDPEL != "All" {
			q = dbutils.AddStringFilter(q, "AND", "dt_temuan.idpel", filter.IDPEL)
		}

		if filter.NotAssignedOnly {
			q = dbutils.AddCustomFilter(q, "AND", "NOT (dt_temuan.id", "=", fmt.Sprintf("ANY(%s))", "(SELECT DISTINCT temuan_id FROM dt_user_temuan)"))
		}

		if filter.StatusBayar != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "dt_temuan.status", int64(filter.StatusBayar))
		}

		if filter.DateFrom != "" {
			q = dbutils.AddCustomFilter(q, "AND", "date(dt_temuan.created_at)", ">=", "'"+filter.DateFrom+"'")
		}

		if filter.DateTo != "" {
			q = dbutils.AddCustomFilter(q, "AND", "date(dt_temuan.created_at)", "<", "'"+filter.DateTo+"'")
		}
	}

	return q

}
