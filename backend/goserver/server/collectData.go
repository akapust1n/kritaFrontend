package server

import (
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"
	"strconv"
)

func count(query string, whatCount string) float64 {
	rows, err := Db.Query(query, whatCount)
	checkErr(err)
	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
		fmt.Println("scanned!")

	}
	return float64(count) //golang хранит числа в json только в флоат64
}

// //ПЕРЕМЕННОЕ КОЛИЧЕСТВО АРГУМЕНТОВ СДЕЛАТЬ
// func countOther(query string, whatCount1 string, whatCount2 string) float64 {
// 	rows, err := Db.Query(query, whatCount1, whatCount2)
// 	checkErr(err)
// 	var count int

// 	for rows.Next() {
// 		err := rows.Scan(&count)
// 		checkErr(err)
// 		fmt.Println("scanned other platforms!")

// 	}
// 	return float64(count) //golang хранит числа в json только в флоат64
// }
func countOther(query string, queryNotEql string, args ...string) float64 {
	execQuery := query
	for i, _ := range args { //Зачем я использую препаред стейтмент, когда можно без него?
		if i != 0 {
			execQuery += " and "
		}
		//execQuery += queryNotEql + "'" + v + "'"
		execQuery += queryNotEql + "$" + strconv.Itoa(i+1)

	}
	//execQuery += ";"
	fmt.Println(execQuery)
	temp, _ := Db.Prepare(execQuery)
	fmt.Println(len(args))

	///rows, err := temp.Query(args)
	old := args
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}
	fmt.Println(new...)
	rows, err := temp.Query(new...)

	checkErr(err)

	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
		fmt.Println("scanned other abstruct platforms!")

	}
	return float64(count) //golang хранит числа в json только в флоат64
}

func CollectData() {
	var result md.CollectedData
	const osQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo where  data -> 'platform'->>'os' = $1"
	result.Platform.Os.Linux = count(osQuery, "linux")
	fmt.Println(result.Platform.Os.Linux)
	result.Platform.Os.Windows = count(osQuery, "windows")
	result.Platform.Os.Mac = count(osQuery, "mac")
	const osOtherQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo WHERE "
	const osOtherNotEql = "data -> 'platform'->>'os' !="
	result.Platform.Os.Other = countOther(osOtherQuery, osOtherNotEql, "linux", "windows", "mac")

	//const osOtherQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo where  data -> 'platform'->>'os' != $1 and data -> 'platform'->>'os'!=$2"
	//result.Platform.Os.Other = countOther(osOtherQuery, "windows", "linux")

	const archCPUQuery = "select count( data -> 'cpu'->>'architecture')   from generalInfo where data -> 'cpu'->>'architecture' = $1"
	result.CPU.Architecture.X86_64 = count(archCPUQuery, "x86_64")
	result.CPU.Architecture.X86 = count(archCPUQuery, "i386") // может быть что-то другое нужно
	const archCPUOtherQuery = "select count( data -> 'cpu'->>'architecture')   from generalInfo where "
	const archCPUOtherNotEql = "data -> 'cpu'->>'architecture' !="
	result.CPU.Architecture.Other = countOther(archCPUOtherQuery, archCPUOtherNotEql, "x86_64", "i386")

	fmt.Println("BEFORE JSON")

	const insertQuery = "INSERT into agregatedInfo(data)  VALUES($1)"
	b, err := json.Marshal(result)
	fmt.Println("SUCCESS JSON")
	fmt.Println(b)
	checkErr(err)
	_, err = Db.Exec(insertQuery, b)
	checkErr(err)
}
