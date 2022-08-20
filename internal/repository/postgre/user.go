package postgre

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/dbutils"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
)

func (db *Postgre) GetUserPassword(username string) (res *presentation.GetUserPasswordResponse, err error) {
	q := `SELECT id, username, password, role from dt_user WHERE username = $1 LIMIT 1 OFFSET 0`

	rows, err := db.chiefDatabase.Master.Queryx(q, username)
	if err != nil {
		return nil, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "GetUserPassword",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetUserPasswordResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetUserPassword",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = &_t
	}

	return res, nil
}

func (db *Postgre) GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error) {
	q := `SELECT id, fullname FROM dt_user`

	// Implement Filter
	if filter != nil {
		if filter.Name != "" && filter.Name != "All" {
			q = dbutils.AddStringFilter(q, "AND", "fullname", filter.Name)
		}
		if filter.Role != 0 {
			q = dbutils.AddBigintFilter(q, "AND", "role", int64(filter.Role))
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
			FunctionName: "GetDataPetugas",
			Description:  "failed running queryx",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		var _t presentation.GetDataPetugasResponse

		err = rows.StructScan(&_t)
		if err != nil {
			return nil, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "GetDataPetugas",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}

		res = append(res, _t)
	}

	return res, nil
}

func (db *Postgre) AssignToSoToUser(userID int64, toSoIDs []int64) (id int64, err error) {
	q := `INSERT INTO dt_user_to_so (to_so_id, user_id) VALUES`

	queryParamLen := 2

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range toSoIDs {
		q = fmt.Sprintf("%s ($%d::BIGINT, $%d::BIGINT),", q, paramCount, paramCount+1)
		paramArgs = append(paramArgs, v, userID)
		paramCount += queryParamLen
	}

	rows, err := db.chiefDatabase.Master.Queryx(fmt.Sprintf("%s RETURNING id", q[:len(q)-1]), paramArgs...)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "AssignToSoToUser",
			Description:  "failed exec",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "AssignToSoToUser",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	return id, nil
}

func (db *Postgre) AssignTemuanToUser(userID int64, temuanIDs []int64) (id int64, err error) {
	q := `INSERT INTO dt_user_temuan (temuan_id, user_id) VALUES`

	queryParamLen := 2

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range temuanIDs {
		q = fmt.Sprintf("%s ($%d::BIGINT, $%d::BIGINT),", q, paramCount, paramCount+1)
		paramArgs = append(paramArgs, v, userID)
		paramCount += queryParamLen
	}

	q = fmt.Sprintf("%s RETURNING id", q[:len(q)-1])

	rows, err := db.chiefDatabase.Master.Queryx(q, paramArgs...)
	if err != nil {
		return 0, response.InternalError{
			Type:         "Repo",
			Name:         "Postgre",
			FunctionName: "AssignToSoToUser",
			Description:  "failed exec",
			Trace:        err,
		}.Error()
	}

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, response.InternalError{
				Type:         "Repo",
				Name:         "Postgre",
				FunctionName: "AssignToSoToUser",
				Description:  "failed scan",
				Trace:        err,
			}.Error()
		}
	}

	return id, nil
}
