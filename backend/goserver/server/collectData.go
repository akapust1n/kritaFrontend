package server

import (
	"fmt"
	md "kritaServers/backend/goserver/server/models"
)

func count(query string, whatCount string) float64 {
	rows, err := Db.Query(query, whatCount)
	checkErr(err)
	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
		fmt.Printf("scanned!")

		//fmt.Printf(string(rows))
		// fmt.Printf(string(cameToServer))

	}
	return float64(count) //golang хранит числа в json только в флоат64
}
func CollectData() {
	var result md.CollectedData
	const osQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo where  data -> 'platform'->>'os' = $1"
	result.Platform.Os.Linux = count(osQuery, "linux")
	fmt.Println(result.Platform.Os.Linux)

}
