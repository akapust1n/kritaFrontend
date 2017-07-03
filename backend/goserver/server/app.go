package server

import (
	"gopkg.in/mgo.v2"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var Session *mgo.Session

func InitDB() {
	var err error
	Session, err = mgo.Dial("mongodb://localhost")

	checkErr(err)

}
