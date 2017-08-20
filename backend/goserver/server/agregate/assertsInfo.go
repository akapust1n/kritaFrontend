package agregate

import (
	"fmt"
	serv "kritaServers/backend/goserver/server"
	md "kritaServers/backend/goserver/server/models"

	"gopkg.in/mgo.v2/bson"
)

func AgregateAsserts() {

	c := serv.Session.DB("telemetry").C("asserts")

	var results []md.AssertsCollected

	err := c.Find(bson.M{"assert": bson.M{"$exists": true}}).Limit(100).All(&results)
	serv.CheckErr(err)
	fmt.Println("hey")
	agregatedAsserts.Asserts = results
	//fmt.Println(results)

}
