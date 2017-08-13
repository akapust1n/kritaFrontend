package agregate

import (
	"bufio"
	"fmt"
	serv "kritaServers/backend/goserver/server"
	md "kritaServers/backend/goserver/server/models"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func countToolsUse(name string) (float64, float64) {
	resultsCount := []bson.M{}
	resultsAvg := []bson.M{}

	var countUse float64
	var averageTimeUse float64
	c := serv.Session.DB("telemetry").C("tools")
	pipe := c.Pipe([]bson.M{{"$unwind": "$tools"}, {"$match": bson.M{"tools.toolname": name}}, {"$group": bson.M{"_id": "$tools.toolname", "total_count": bson.M{"$sum": "$tools.countuse"}}}})
	pipe2 := c.Pipe([]bson.M{{"$unwind": "$tools"}, {"$match": bson.M{"tools.toolname": name}}, {"$group": bson.M{"_id": "$tools.toolname", "total_count": bson.M{"$avg": "$tools.time"}}}})

	resp := []bson.M{}
	resp2 := []bson.M{}
	err := pipe.All(&resp)
	err = pipe2.All(&resp2)
	serv.CheckErr(err)
	err = pipe.All(&resultsCount)
	err = pipe2.All(&resultsAvg)
	serv.CheckErr(err)
	if len(resultsCount) > 0 {
		countUse, _ = resultsCount[0]["total_count"].(float64)
		//	fmt.Println(name, num)
	}
	if len(resultsAvg) > 0 {
		averageTimeUse, _ = resultsAvg[0]["total_count"].(float64)
		fmt.Println(name, averageTimeUse)
	}
	return countUse, averageTimeUse
}
func AgregateTools() {
	file, err := os.Open("list_tools.txt")
	serv.CheckErr(err)
	defer file.Close()

	var ToolUse md.ToolsCollected
	var ToolActivate md.ToolsCollected
	var ToolsUse []md.ToolsCollected
	var ToolsActivate []md.ToolsCollected

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ToolUse.Name = scanner.Text()
		ToolUse.CountUse, _ = countToolsUse("/Use/" + ToolUse.Name)
		ToolsUse = append(ToolsUse, ToolUse)

		ToolActivate.Name = ToolUse.Name
		ToolActivate.CountUse, _ = countToolsUse("/Activate/" + ToolActivate.Name)
		ToolsActivate = append(ToolsActivate, ToolUse)
	}
	agregatedTools.ToolsActivate = ToolsActivate
	agregatedTools.ToolsUse = ToolsUse
	err = scanner.Err()
	serv.CheckErr(err)

}
