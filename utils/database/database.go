package database

import (
	"database/sql"
	"sync"

	"github.com/accuknox/observability/utils/constants"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var once sync.Once
var db *sql.DB

//Connect DB
func ConnectDB() *sql.DB {
	once.Do(func() {
		//Open database
		conn, err := sql.Open("sqlite3", "knox.db")
		db = conn
		if err != nil {
			log.Panic().Msg("Error in DB connection : " + err.Error())
		}

		//Create Cilium table if not exist
		db.Exec(constants.CREATE_CILIUM_TABLE)
		//Create Kubearmor table if not exist
		db.Exec(constants.CREATE_KUBEARMOR_TABLE)

	})
	return db
}
