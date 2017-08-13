package agregate

import (
	"encoding/json"
	serv "kritaServers/backend/goserver/server"
	md "kritaServers/backend/goserver/server/models"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var agreagtedData md.CollectedInstallData
var agregatedTools md.ToolsCollectedData
var agregatedActions []md.ActionCollected
var agregatedImageInfo md.ImageCollected

func getFloat64(n int, err error) float64 {
	serv.CheckErr(err)
	return float64(n)
}
func checkOtherCount(count float64) float64 {
	if count < 0 {
		return 0
	}
	return count
}

func getProportion(specificCount float64, totalCount float64) string {
	result := strconv.FormatFloat(specificCount/totalCount, 'f', -1, 32)
	return result
}

func countExist(category string, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{category: bson.M{"$exists": true}}).Count())
	return count
}

// func Agregated(type) string {
// 	out, _ := json.Marshal(agreagtedData)
// 	return string(out)
// }

func Agregated(dataType string) string {
	switch dataType {
	case "install":
		out, _ := json.Marshal(agreagtedData)
		return string(out)
	case "tools":
		out, _ := json.Marshal(agregatedTools)
		return string(out)
	case "actions":
		out, _ := json.Marshal(agregatedActions)
		return string(out)
	case "images":
		out, _ := json.Marshal(agregatedImageInfo)
		return string(out)

	default:
		return string("error")
	}
}
