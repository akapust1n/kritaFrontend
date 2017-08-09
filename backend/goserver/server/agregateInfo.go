package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var agreagtedData md.CollectedInstallData
var agregatedTools md.ToolsCollectedData
var agregatedActions []md.ActionCollected
var agregatedImageInfo md.ImageCollected

func getFloat64(n int, err error) float64 {
	CheckErr(err)
	return float64(n)
}
func checkOtherCount(count float64) float64 {
	if count < 0 {
		return 0
	}
	return count

}
func countActionsUse(name string) float64 {
	results := []bson.M{}
	c := Session.DB("telemetry").C("actions")
	pipe := c.Pipe([]bson.M{{"$unwind": "$actions"}, {"$match": bson.M{"actions.actionname": name}}, {"$group": bson.M{"_id": "$actions.actionname", "total_count": bson.M{"$sum": "$actions.countuse"}}}})
	//fmt.Println(pipe)
	resp := []bson.M{}
	err := pipe.All(&resp)
	CheckErr(err)
	//fmt.Println(resp) // simple print proving it's working
	err = pipe.All(&results)
	CheckErr(err)
	if len(results) > 0 {
		num, _ := results[0]["total_count"].(float64)

		return num
	}
	return 0
}
func countToolsUse(name string) (float64, float64) {
	resultsCount := []bson.M{}
	resultsAvg := []bson.M{}

	var countUse float64
	var averageTimeUse float64
	c := Session.DB("telemetry").C("tools")
	pipe := c.Pipe([]bson.M{{"$unwind": "$tools"}, {"$match": bson.M{"tools.toolname": name}}, {"$group": bson.M{"_id": "$tools.toolname", "total_count": bson.M{"$sum": "$tools.countuse"}}}})
	pipe2 := c.Pipe([]bson.M{{"$unwind": "$tools"}, {"$match": bson.M{"tools.toolname": name}}, {"$group": bson.M{"_id": "$tools.toolname", "total_count": bson.M{"$avg": "$tools.time"}}}})

	resp := []bson.M{}
	resp2 := []bson.M{}
	err := pipe.All(&resp)
	err = pipe2.All(&resp2)
	CheckErr(err)
	err = pipe.All(&resultsCount)
	err = pipe2.All(&resultsAvg)
	CheckErr(err)
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
func getWidth(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.width": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func getHeight(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.height": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func getLayer(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.numlayers": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func getFileSize(high int, low int, session *mgo.Collection) float64 {
	count := getFloat64(session.Find(bson.M{"images.size": bson.M{"$gt": high, "$lte": low}}).Count())
	return count
}
func AgregateImageProps() {
	c := Session.DB("telemetry").C("images")
	var ic md.ImageCollected

	ic.WD.L500 = getWidth(0, 500, c)
	ic.WD.L1000 = getWidth(500, 1000, c)
	ic.WD.L2000 = getWidth(1000, 2000, c)
	ic.WD.L4000 = getWidth(2000, 4000, c)
	ic.WD.L8000 = getWidth(4000, 8000, c)
	ic.WD.M8000 = getFloat64(c.Find(bson.M{"images.width": bson.M{"$gt": 8000}}).Count())

	ic.HD.L500 = getHeight(500, 1000, c)
	ic.HD.L1000 = getHeight(1000, 2000, c)
	ic.HD.L2000 = getHeight(1000, 2000, c)
	ic.HD.L4000 = getHeight(2000, 4000, c)
	ic.HD.L8000 = getHeight(4000, 8000, c)
	ic.HD.M8000 = getFloat64(c.Find(bson.M{"images.height": bson.M{"$gt": 8000}}).Count())

	ic.LD.L1 = getLayer(0, 1, c)
	ic.LD.L2 = getLayer(1, 2, c)
	ic.LD.L4 = getLayer(2, 4, c)
	ic.LD.L8 = getLayer(4, 8, c)
	ic.LD.L16 = getLayer(8, 16, c)
	ic.LD.L32 = getLayer(16, 32, c)
	ic.LD.L64 = getLayer(32, 64, c)
	ic.LD.M64 = getFloat64(c.Find(bson.M{"images.numlayers": bson.M{"$gt": 8000}}).Count())

	ic.ID.Mb1 = getFileSize(0, 1, c)
	ic.ID.Mb5 = getFileSize(1, 5, c)
	ic.ID.Mb10 = getFileSize(5, 10, c)
	ic.ID.Mb25 = getFileSize(10, 25, c)
	ic.ID.Mb50 = getFileSize(25, 50, c)
	ic.ID.Mb100 = getFileSize(50, 100, c)
	ic.ID.Mb200 = getFileSize(100, 200, c)
	ic.ID.Mb400 = getFileSize(200, 400, c)
	ic.ID.Mb800 = getFileSize(400, 800, c)
	ic.ID.More800 = getFloat64(c.Find(bson.M{"images.size": bson.M{"$gt": 800}}).Count())
	ic.CPD.RGBA = getFloat64(c.Find(bson.M{"images.colorprofile": "RGB/Alpha"}).Count())
	ic.CPD.CMYK = getFloat64(c.Find(bson.M{"images.colorprofile": "CMYK/Alpha"}).Count())
	ic.CPD.Grayscale = getFloat64(c.Find(bson.M{"images.colorprofile": "Grayscale/Alpha"}).Count())
	ic.CPD.Lab = getFloat64(c.Find(bson.M{"images.colorprofile": "L*a*b*/Alpha"}).Count())
	ic.CPD.XYZ = getFloat64(c.Find(bson.M{"images.colorprofile": "XYZ/Alpha"}).Count())
	ic.CPD.YCbCr = getFloat64(c.Find(bson.M{"images.colorprofile": "YCbCr/Alpha"}).Count())

	agregatedImageInfo = ic
}
func AgregateActions() {
	file, err := os.Open("list_actions.txt")
	CheckErr(err)
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
	CheckErr(err)
}

func AgregateTools() {
	file, err := os.Open("list_tools.txt")
	CheckErr(err)
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
	CheckErr(err)

}
func AgregateInstalInfo() {
	c := Session.DB("telemetry").C("installInfo")
	//Some error takes place. We lose a little bit of data
	countRecords := getFloat64(c.Find(bson.M{}).Count())

	//compiler
	agreagtedData.Compiler.Type.GCC = getFloat64(c.Find(bson.M{"compiler.type": "GCC"}).Count())
	agreagtedData.Compiler.Type.Clang = getFloat64(c.Find(bson.M{"compiler.type": "Clang"}).Count())
	agreagtedData.Compiler.Type.MSVC = getFloat64(c.Find(bson.M{"compiler.type": "MSVC"}).Count())
	agreagtedData.Compiler.Type.Other = countRecords - agreagtedData.Compiler.Type.GCC - agreagtedData.Compiler.Type.Clang - agreagtedData.Compiler.Type.MSVC
	agreagtedData.Compiler.Type.Other = checkOtherCount(agreagtedData.Compiler.Type.Other)
	//os
	agreagtedData.Platform.Os.Linux = getFloat64(c.Find(bson.M{"platform.os": "linux"}).Count())
	agreagtedData.Platform.Os.Windows = getFloat64(c.Find(bson.M{"platform.os": "windows"}).Count())
	agreagtedData.Platform.Os.Mac = getFloat64(c.Find(bson.M{"platform.os": "mac"}).Count())
	agreagtedData.Platform.Os.Other = countRecords - agreagtedData.Platform.Os.Linux - agreagtedData.Platform.Os.Windows - agreagtedData.Platform.Os.Mac
	agreagtedData.Platform.Os.Other = checkOtherCount(agreagtedData.Platform.Os.Other)
	fmt.Println("LINUX")
	fmt.Println(agreagtedData.Platform.Os.Linux)
	fmt.Println(agreagtedData.Platform.Os.Other)

	//version os windows
	agreagtedData.Platform.Version.Windows.V7 = getFloat64(c.Find(bson.M{"platform.version": "7"}).Count())
	agreagtedData.Platform.Version.Windows.V8 = getFloat64(c.Find(bson.M{"platform.version": "8"}).Count())
	agreagtedData.Platform.Version.Windows.V81 = getFloat64(c.Find(bson.M{"platform.version": "8.1"}).Count())
	agreagtedData.Platform.Version.Windows.V10 = getFloat64(c.Find(bson.M{"platform.version": "10"}).Count())
	agreagtedData.Platform.Version.Windows.Other = agreagtedData.Platform.Os.Windows - agreagtedData.Platform.Version.Windows.V7 - agreagtedData.Platform.Version.Windows.V8 - agreagtedData.Platform.Version.Windows.V81 - agreagtedData.Platform.Version.Windows.V10
	agreagtedData.Platform.Version.Windows.Other = checkOtherCount(agreagtedData.Platform.Version.Windows.Other)

	//version os linux
	//не уверен, что именно так пишется нужные версии linux
	agreagtedData.Platform.Version.Linux.Ubuntu1404 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.04"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1410 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.10"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1504 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.04"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1510 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.10"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1604 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-16.04"}).Count())
	agreagtedData.Platform.Version.Linux.Other = agreagtedData.Platform.Os.Linux - agreagtedData.Platform.Version.Linux.Ubuntu1404 - agreagtedData.Platform.Version.Linux.Ubuntu1410 - agreagtedData.Platform.Version.Linux.Ubuntu1504 - agreagtedData.Platform.Version.Linux.Ubuntu1510 - agreagtedData.Platform.Version.Linux.Ubuntu1604 - agreagtedData.Platform.Version.Linux.Ubuntu1610 - agreagtedData.Platform.Version.Linux.Ubuntu1704
	agreagtedData.Platform.Version.Linux.Other = checkOtherCount(agreagtedData.Platform.Version.Linux.Other)

	agreagtedData.Platform.Version.Mac.V1012 = getFloat64(c.Find(bson.M{"platform.version": "10.12"}).Count())
	agreagtedData.Platform.Version.Mac.Other = agreagtedData.Platform.Os.Mac - agreagtedData.Platform.Version.Mac.V1012
	agreagtedData.Platform.Version.Mac.Other = checkOtherCount(agreagtedData.Platform.Version.Linux.Other)

	//cpu
	agreagtedData.CPU.Architecture.X86_64 = getFloat64(c.Find(bson.M{"cpu.architecture": "x86_64"}).Count())
	agreagtedData.CPU.Architecture.X86 = getFloat64(c.Find(bson.M{"cpu.architecture": "i386"}).Count())
	agreagtedData.CPU.Architecture.Other = countRecords - agreagtedData.CPU.Architecture.X86_64 - agreagtedData.CPU.Architecture.X86
	agreagtedData.CPU.Architecture.Other = checkOtherCount(agreagtedData.CPU.Architecture.Other)

	agreagtedData.CPU.Cores.C1 = getFloat64(c.Find(bson.M{"cpu.count": "1"}).Count())
	agreagtedData.CPU.Cores.C2 = getFloat64(c.Find(bson.M{"cpu.count": "2"}).Count())
	agreagtedData.CPU.Cores.C3 = getFloat64(c.Find(bson.M{"cpu.count": "3"}).Count())
	agreagtedData.CPU.Cores.C4 = getFloat64(c.Find(bson.M{"cpu.count": "4"}).Count())
	agreagtedData.CPU.Cores.C6 = getFloat64(c.Find(bson.M{"cpu.count": "6"}).Count())
	agreagtedData.CPU.Cores.C8 = getFloat64(c.Find(bson.M{"cpu.count": "8"}).Count())
	agreagtedData.CPU.Cores.Other = countRecords - agreagtedData.CPU.Cores.C1 - agreagtedData.CPU.Cores.C2 - agreagtedData.CPU.Cores.C3 - agreagtedData.CPU.Cores.C4 - agreagtedData.CPU.Cores.C6 - agreagtedData.CPU.Cores.C8

	agreagtedData.Locale.Language.English = getFloat64(c.Find(bson.M{"locale.language": "English"}).Count())
	agreagtedData.Locale.Language.Russian = getFloat64(c.Find(bson.M{"locale.language": "Russian"}).Count())
	agreagtedData.Locale.Language.Other = countRecords - agreagtedData.Locale.Language.English - agreagtedData.Locale.Language.Russian
	agreagtedData.Locale.Language.Other = checkOtherCount(agreagtedData.Locale.Language.Other)

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
