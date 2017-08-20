package server

import (
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"
)

func InsertGeneralInfo(request []byte) {
	var model md.Request
	err := json.Unmarshal(request, &model)
	CheckErr(err)
	c := Session.DB("telemetry").C("installInfo")
	c.Insert(model)
}
func InsertAssertInfo(request []byte) {
	var model md.Assert
	err := json.Unmarshal(request, &model)
	CheckErr(err)
	c := Session.DB("telemetry").C("asserts")
	c.Insert(model)
	fmt.Println("inserted asserts!")
}

func InsertToolInfo(request []byte) {
	var err error
	var tools md.Tool
	err = json.Unmarshal(request, &tools)
	CheckErr(err)
	c := Session.DB("telemetry").C("tools")
	c.Insert(tools)
	fmt.Println("inserted TOOL info!")
}

func InsertActionInfo(request []byte) {
	var err error
	var actions md.Action
	err = json.Unmarshal(request, &actions)
	CheckErr(err)
	c := Session.DB("telemetry").C("actions")
	c.Insert(actions)
	fmt.Println("inserted ACTION info!")
}

func InsertImageInfo(request []byte) {
	var err error
	var images md.Image
	err = json.Unmarshal(request, &images)
	CheckErr(err)
	c := Session.DB("telemetry").C("images")
	c.Insert(images)
	fmt.Println("inserted IMAGE info!")
}
