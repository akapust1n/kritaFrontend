package server

import (
	"fmt"
	"net/http"
)

func AgregatedDataHandler(w http.ResponseWriter, r *http.Request) {
	const selectQuery = "SELECT data from agregatedinfo order by generatedTime desc limit 1"
	rows, err := Db.Query(selectQuery)
	checkErr(err)
	var data []byte

	for rows.Next() {
		err := rows.Scan(&data)
		checkErr(err)
		fmt.Printf("frontend request from db")
		fmt.Printf(string(data))
		// fmt.Printf(string(cameToServer))

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
