package postgre

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgre struct {
	chiefDatabase *PgDB
}

func New(
	chiefDBDriverMaster string,
) (*Postgre, error) {
	var err error
	chiefDatabase, err := InitPostgreDB(chiefDBDriverMaster, "")
	if err != nil {
		return nil, fmt.Errorf("[Postgre][Init] Failed init user database, trace %v", err)
	}

	return &Postgre{
		chiefDatabase: &chiefDatabase,
	}, nil
}

type PgDB struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

func InitPostgreDB(
	masterDriver string,
	slaveDriver string,
) (db PgDB, err error) {
	const DB_DRIVER_NAME_POSTGRE = "postgres"

	db.Master, err = sqlx.Connect(DB_DRIVER_NAME_POSTGRE, masterDriver)
	if err != nil {
		return db, err
	}

	if slaveDriver != "" {
		db.Slave, err = sqlx.Connect(DB_DRIVER_NAME_POSTGRE, slaveDriver)
		if err != nil {
			return db, err
		}
	}

	return db, err
}
