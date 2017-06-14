package server

import (
	"fmt"
)

func InsertGeneralInfo(request []byte) {
	const insertQuery = "INSERT into generalinfo(data)  VALUES($1)"
	fmt.Println(string(request))
	s := string(request[:])
	_, err := Db.Exec(insertQuery, s)
	checkErr(err)

}
