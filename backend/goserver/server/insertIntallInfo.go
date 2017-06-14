package server

import (
	"fmt"
)

func InsertGeneralInfo(request []byte) {
	const insertQuery = "INSERT into generalinfo(data)  VALUES($1)"
	s := string(request[:])
	_, err := Db.Exec(insertQuery, s)
	if err != nil {
		fmt.Println("insert error!")
	} else {
		fmt.Println("insert ended!")
	}

}
