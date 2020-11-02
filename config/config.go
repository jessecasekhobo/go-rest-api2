package config

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetMsSQLDB() (condb *sql.DB, err error) {
	condb, errdb := sql.Open("mssql", "server=localhost;user id=sa;password=asupersecurepassword;")
	if errdb != nil {
		fmt.Println(" Error opening db:", errdb.Error())
	}
	return
}
