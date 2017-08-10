package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"
	"os"
	"strconv"

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

func getProportion(specificCount float64, totalCount float64) string {
	result := strconv.FormatFloat(specificCount/totalCount, 'f', -1, 32)
	return result
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

	ic.WD.L500.Count = getWidth(0, 500, c)
	ic.WD.L1000.Count = getWidth(500, 1000, c)
	ic.WD.L2000.Count = getWidth(1000, 2000, c)
	ic.WD.L4000.Count = getWidth(2000, 4000, c)
	ic.WD.L8000.Count = getWidth(4000, 8000, c)
	ic.WD.M8000.Count = getFloat64(c.Find(bson.M{"images.width": bson.M{"$gt": 8000}}).Count())

	ic.HD.L500.Count = getHeight(500, 1000, c)
	ic.HD.L1000.Count = getHeight(1000, 2000, c)
	ic.HD.L2000.Count = getHeight(1000, 2000, c)
	ic.HD.L4000.Count = getHeight(2000, 4000, c)
	ic.HD.L8000.Count = getHeight(4000, 8000, c)
	ic.HD.M8000.Count = getFloat64(c.Find(bson.M{"images.height": bson.M{"$gt": 8000}}).Count())

	ic.LD.L1.Count = getLayer(0, 1, c)
	ic.LD.L2.Count = getLayer(1, 2, c)
	ic.LD.L4.Count = getLayer(2, 4, c)
	ic.LD.L8.Count = getLayer(4, 8, c)
	ic.LD.L16.Count = getLayer(8, 16, c)
	ic.LD.L32.Count = getLayer(16, 32, c)
	ic.LD.L64.Count = getLayer(32, 64, c)
	ic.LD.M64.Count = getFloat64(c.Find(bson.M{"images.numlayers": bson.M{"$gt": 8000}}).Count())

	ic.ID.Mb1.Count = getFileSize(0, 1, c)
	ic.ID.Mb5.Count = getFileSize(1, 5, c)
	ic.ID.Mb10.Count = getFileSize(5, 10, c)
	ic.ID.Mb25.Count = getFileSize(10, 25, c)
	ic.ID.Mb50.Count = getFileSize(25, 50, c)
	ic.ID.Mb100.Count = getFileSize(50, 100, c)
	ic.ID.Mb200.Count = getFileSize(100, 200, c)
	ic.ID.Mb400.Count = getFileSize(200, 400, c)
	ic.ID.Mb800.Count = getFileSize(400, 800, c)
	ic.ID.More800.Count = getFloat64(c.Find(bson.M{"images.size": bson.M{"$gt": 800}}).Count())

	ic.CPD.RGBA.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "RGB/Alpha"}).Count())
	ic.CPD.CMYK.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "CMYK/Alpha"}).Count())
	ic.CPD.Grayscale.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "Grayscale/Alpha"}).Count())
	ic.CPD.Lab.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "L*a*b*/Alpha"}).Count())
	ic.CPD.XYZ.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "XYZ/Alpha"}).Count())
	ic.CPD.YCbCr.Count = getFloat64(c.Find(bson.M{"images.colorprofile": "YCbCr/Alpha"}).Count())

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
	agreagtedData.Compiler.Type.GCC.Count = getFloat64(c.Find(bson.M{"compiler.type": "GCC"}).Count())
	agreagtedData.Compiler.Type.Clang.Count = getFloat64(c.Find(bson.M{"compiler.type": "Clang"}).Count())
	agreagtedData.Compiler.Type.MSVC.Count = getFloat64(c.Find(bson.M{"compiler.type": "MSVC"}).Count())
	agreagtedData.Compiler.Type.Other.Count = countRecords - agreagtedData.Compiler.Type.GCC.Count - agreagtedData.Compiler.Type.Clang.Count - agreagtedData.Compiler.Type.MSVC.Count
	agreagtedData.Compiler.Type.Other.Count = checkOtherCount(agreagtedData.Compiler.Type.Other.Count)
	//os
	agreagtedData.Platform.Os.Linux.Count = getFloat64(c.Find(bson.M{"platform.os": "linux"}).Count())
	agreagtedData.Platform.Os.Windows.Count = getFloat64(c.Find(bson.M{"platform.os": "windows"}).Count())
	agreagtedData.Platform.Os.Mac.Count = getFloat64(c.Find(bson.M{"platform.os": "mac"}).Count())
	agreagtedData.Platform.Os.Other.Count = countRecords - agreagtedData.Platform.Os.Linux.Count - agreagtedData.Platform.Os.Windows.Count - agreagtedData.Platform.Os.Mac.Count
	agreagtedData.Platform.Os.Other.Count = checkOtherCount(agreagtedData.Platform.Os.Other.Count)

	//version os windows
	agreagtedData.Platform.Version.Windows.V7.Count = getFloat64(c.Find(bson.M{"platform.version": "7"}).Count())
	agreagtedData.Platform.Version.Windows.V8.Count = getFloat64(c.Find(bson.M{"platform.version": "8"}).Count())
	agreagtedData.Platform.Version.Windows.V81.Count = getFloat64(c.Find(bson.M{"platform.version": "8.1"}).Count())
	agreagtedData.Platform.Version.Windows.V10.Count = getFloat64(c.Find(bson.M{"platform.version": "10"}).Count())
	agreagtedData.Platform.Version.Windows.Other.Count = agreagtedData.Platform.Os.Windows.Count - agreagtedData.Platform.Version.Windows.V7.Count - agreagtedData.Platform.Version.Windows.V8.Count - agreagtedData.Platform.Version.Windows.V81.Count - agreagtedData.Platform.Version.Windows.V10.Count
	agreagtedData.Platform.Version.Windows.Other.Count = checkOtherCount(agreagtedData.Platform.Version.Windows.Other.Count)

	//version os linux
	//не уверен, что именно так пишется нужные версии linux
	agreagtedData.Platform.Version.Linux.Ubuntu1404.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.04"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1410.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.10"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1504.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.04"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1510.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.10"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1604.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-16.04"}).Count())
	agreagtedData.Platform.Version.Linux.Other.Count = agreagtedData.Platform.Os.Linux.Count - agreagtedData.Platform.Version.Linux.Ubuntu1404.Count - agreagtedData.Platform.Version.Linux.Ubuntu1410.Count - agreagtedData.Platform.Version.Linux.Ubuntu1504.Count - agreagtedData.Platform.Version.Linux.Ubuntu1510.Count - agreagtedData.Platform.Version.Linux.Ubuntu1604.Count - agreagtedData.Platform.Version.Linux.Ubuntu1610.Count - agreagtedData.Platform.Version.Linux.Ubuntu1704.Count
	agreagtedData.Platform.Version.Linux.Other.Count = checkOtherCount(agreagtedData.Platform.Version.Linux.Other.Count)

	agreagtedData.Platform.Version.Mac.V1012.Count = getFloat64(c.Find(bson.M{"platform.version": "10.12"}).Count())
	agreagtedData.Platform.Version.Mac.Other.Count = agreagtedData.Platform.Os.Mac.Count - agreagtedData.Platform.Version.Mac.V1012.Count
	agreagtedData.Platform.Version.Mac.Other.Count = checkOtherCount(agreagtedData.Platform.Version.Linux.Other.Count)

	agreagtedData.Platform.Version.Linux.Ubuntu1404.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1404.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Ubuntu1410.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1410.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Ubuntu1504.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1504.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Ubuntu1510.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1510.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Ubuntu1604.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1604.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Ubuntu1610.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1610.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Ubuntu1704.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1704.Count, countRecords)
	agreagtedData.Platform.Version.Linux.Other.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Other.Count, countRecords)

	//cpu
	agreagtedData.CPU.Architecture.X86_64.Count = getFloat64(c.Find(bson.M{"cpu.architecture": "x86_64"}).Count())
	agreagtedData.CPU.Architecture.X86.Count = getFloat64(c.Find(bson.M{"cpu.architecture": "i386"}).Count())
	agreagtedData.CPU.Architecture.Other.Count = countRecords - agreagtedData.CPU.Architecture.X86_64.Count - agreagtedData.CPU.Architecture.X86.Count
	agreagtedData.CPU.Architecture.Other.Count = checkOtherCount(agreagtedData.CPU.Architecture.Other.Count)

	agreagtedData.CPU.Cores.C1.Count = getFloat64(c.Find(bson.M{"cpu.count": "1"}).Count())
	agreagtedData.CPU.Cores.C2.Count = getFloat64(c.Find(bson.M{"cpu.count": "2"}).Count())
	agreagtedData.CPU.Cores.C3.Count = getFloat64(c.Find(bson.M{"cpu.count": "3"}).Count())
	agreagtedData.CPU.Cores.C4.Count = getFloat64(c.Find(bson.M{"cpu.count": "4"}).Count())
	agreagtedData.CPU.Cores.C6.Count = getFloat64(c.Find(bson.M{"cpu.count": "6"}).Count())
	agreagtedData.CPU.Cores.C8.Count = getFloat64(c.Find(bson.M{"cpu.count": "8"}).Count())
	agreagtedData.CPU.Cores.Other.Count = countRecords - agreagtedData.CPU.Cores.C1.Count - agreagtedData.CPU.Cores.C2.Count - agreagtedData.CPU.Cores.C3.Count - agreagtedData.CPU.Cores.C4.Count - agreagtedData.CPU.Cores.C6.Count - agreagtedData.CPU.Cores.C8.Count

	//language
	agreagtedData.Locale.Language.English.Count = getFloat64(c.Find(bson.M{"locale.language": "English"}).Count())
	agreagtedData.Locale.Language.Russian.Count = getFloat64(c.Find(bson.M{"locale.language": "Russian"}).Count())
	agreagtedData.Locale.Language.Other.Count = countRecords - agreagtedData.Locale.Language.English.Count - agreagtedData.Locale.Language.Russian.Count
	agreagtedData.Locale.Language.Other.Count = checkOtherCount(agreagtedData.Locale.Language.Other.Count)

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
