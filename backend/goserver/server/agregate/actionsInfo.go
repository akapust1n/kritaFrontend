package agregate

import (
	"bufio"
	"fmt"
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
	file, err := os.Open("list_actions_generated.txt")
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
func AgregateListActions() {
	c := serv.Session.DB("telemetry").C("actions")
	var actions []string
	err := c.Find(nil).Distinct("actions.actionname", &actions)
	serv.CheckErr(err)
	file, err := os.Create("list_actions_generated.txt")
	serv.CheckErr(err)
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, action := range actions {
		if action != "" {
			fmt.Fprintln(w, action)
		}
	}
	err = w.Flush()
	serv.CheckErr(err)
}
