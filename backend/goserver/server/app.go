package server

import "database/sql"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("postgres", "user=root password=1111 dbname=root") //небезопасно, но пока сойдет

	err = Db.Ping()

	checkErr(err)
}
