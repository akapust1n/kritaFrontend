package agregate

import (
	"encoding/json"
	serv "kritaServers/backend/goserver/server"
	md "kritaServers/backend/goserver/server/models"
	"math"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var agreagtedData md.CollectedInstallData
var agregatedTools md.ToolsCollectedData
var agregatedActions []md.ActionCollected
var agregatedActionsJson md.AgregatedActionsJson
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
	divended := specificCount / totalCount
	if math.IsNaN(divended) {
		return "0%%" //+ "%"
	}
	result := strconv.FormatFloat(specificCount/totalCount*100, 'f', -1, 32)
	return result + "%%"
}

func countExist(category string, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{category: bson.M{"$exists": true}}).Count())
	return count
}

// func Agregated(type) string {
// 	out, _ := json.Marshal(agreagtedData)
// 	return string(out)
// }

func AgregatedImages(type1 string) string {
	switch type1 {
	case "height":
		out, _ := json.Marshal(agregatedImageInfo.HD)
		return string(out)
	case "width":
		out, _ := json.Marshal(agregatedImageInfo.WD)
		return string(out)
	case "numlayers":
		out, _ := json.Marshal(agregatedImageInfo.LD)
		return string(out)
	case "filesize":
		out, _ := json.Marshal(agregatedImageInfo.ID)
		return string(out)
	case "colorprofile":
		out, _ := json.Marshal(agregatedImageInfo.CPD)
		return string(out)
	default:
		return string("error")
	}
}
func AgregatedInstall(type1 string) string {
	switch type1 {
	case "os":
		out, _ := json.Marshal(agreagtedData.Platform.Os)
		return string(out)
	case "windows":
		out, _ := json.Marshal(agreagtedData.Platform.Version.Windows)
		return string(out)
	case "linux":
		out, _ := json.Marshal(agreagtedData.Platform.Version.Linux)
		return string(out)
	case "mac":
		out, _ := json.Marshal(agreagtedData.Platform.Version.Mac)
		return string(out)
	case "architecture":
		out, _ := json.Marshal(agreagtedData.CPU.Architecture)
		return string(out)
	case "cores":
		out, _ := json.Marshal(agreagtedData.CPU.Cores)
		return string(out)
	case "compiler":
		out, _ := json.Marshal(agreagtedData.Compiler.Type)
		return string(out)
	case "locale":
		out, _ := json.Marshal(agreagtedData.Locale.Language)
		return string(out)
	default:
		return string("error")
	}
}
func Agregated(dataType string) string {
	switch dataType {
	case "install":
		out, _ := json.Marshal(agreagtedData)
		return string(out)
	case "tools":
		out, _ := json.Marshal(agregatedTools)
		return string(out)
	case "actions":
		agregatedActionsJson.Actions = agregatedActions
		out, _ := json.Marshal(agregatedActionsJson)
		return string(out)
	case "images":
		out, _ := json.Marshal(agregatedImageInfo)
		return string(out)

	default:
		return string("error")
	}
}
