package server

import (
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"
)

func InsertGeneralInfo(request []byte) {
	var model md.Request

	err := json.Unmarshal(request, &model)
	checkErr(err)
	c := Session.DB("telemetry").C("installInfo")
	c.Insert(model)
	fmt.Println("inserted info!")

}

func InsertToolInfo(request []byte) {
	var err error
	var tools []md.Tool
	err = json.Unmarshal(request, &tools)
	checkErr(err)
	c := Session.DB("telemetry").C("tools")
	c.Insert(tools)
	fmt.Println("inserted TOOL info!")
}
