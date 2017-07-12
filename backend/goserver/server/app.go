package server

import (
	"gopkg.in/mgo.v2"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var Session *mgo.Session

func InitDB() {
	var err error
	Session, err = mgo.Dial("mongodb://localhost")

	CheckErr(err)

}
