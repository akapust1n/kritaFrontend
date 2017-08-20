package agregate

import (
	"bufio"
	serv "kritaServers/backend/goserver/server"
	md "kritaServers/backend/goserver/server/models"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func countActionsUse(name string) float64 {
	results := []bson.M{}
	c := serv.Session.DB("telemetry").C("actions")
	pipe := c.Pipe([]bson.M{{"$unwind": "$actions"}, {"$match": bson.M{"actions.actionname": name}}, {"$group": bson.M{"_id": "$actions.actionname", "total_count": bson.M{"$sum": "$actions.countuse"}}}})
	//fmt.Println(pipe)
	err := pipe.All(&results)
	serv.CheckErr(err)
	if len(results) > 0 {
		num, _ := results[0]["total_count"].(float64)

		return num
	}
	return 0
}
func AgregateActions() {
	file, err := os.Open("list_actions.txt")
	serv.CheckErr(err)
	defer file.Close()

	var action md.ActionCollected
	var actions []md.ActionCollected

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		action.Name = scanner.Text()
		action.CountUse = countActionsUse(action.Name)
		actions = append(actions, action)
	}
	agregatedActions = actions
	err = scanner.Err()
	serv.CheckErr(err)
}
