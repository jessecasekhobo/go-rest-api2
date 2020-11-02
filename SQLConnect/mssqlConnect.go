package mssqlConnect

/*
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

condb, errdb := sql.Open("mssql", "server=localhost;user id=sa;password=;")
if errdb != nil {
	fmt.Println(" Error open db:", errdb.Error())
}



var (
	rownum   string
	evtclass string
)

rows, err := condb.Query("select TOP (100) [RowNumber], [EventClass] FROM [vendingmachine].[dbo].[temp];")
if err != nil {
	log.Fatal(err)
}
for rows.Next() {
	err := rows.Scan(&rownum, &evtclass)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rownum, evtclass)
}

defer condb.Close()

}

defer condb.Close()
*/
