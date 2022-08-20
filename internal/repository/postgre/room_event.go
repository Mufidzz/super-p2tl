package postgre

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
)

func (db *Postgre) CreateLogRoomEvent(input []presentation.CreateLogRoomEventRequest) (lastInsertedID int, err error) {
	q := `INSERT INTO ws_room_event (room_id, namespace, args, status, event) VALUES`

	queryParamLen := 5

	paramCount := 1
	paramArgs := []interface{}{}

	for _, v := range input {
		q = fmt.Sprintf("%s ($%d::TEXT, $%d::TEXT, $%d::TEXT, $%d::INTEGER, $%d::TEXT),", q, paramCount, paramCount+1, paramCount+2, paramCount+3, paramCount+4)
		paramArgs = append(paramArgs, v.Room, v.Namespace, v.Args, v.Status, v.Event)
		paramCount += queryParamLen
	}

	rows, err := db.chiefDatabase.Master.Queryx(fmt.Sprintf("%s RETURNING id", q[:len(q)-1]), paramArgs...)
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
