package server

import (
	"fmt"
	"net/http"
	"time"
)

func AgregatedDataHandler(w http.ResponseWriter, r *http.Request) {
	const selectQuery = "SELECT cameToServer,data from generalinfo order by cameToServer desc limit 1"
	rows, err := Db.Query(selectQuery)
	checkErr(err)
	for rows.Next() {
		var data []byte
		var cameToServer time.Time
		err := rows.Scan(&cameToServer, &data)
		checkErr(err)
		fmt.Printf("Data from db")
		fmt.Printf(string(data))
		// fmt.Printf(string(cameToServer))

	}

}
