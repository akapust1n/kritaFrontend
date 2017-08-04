package server

import (
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"

	"gopkg.in/mgo.v2/bson"
)

var agreagtedData md.CollectedData

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
		num, _ := results[0][name].(float64)
		return num
	}
	return 0
}
func AgregateActions() {
	agreagtedData.Actions.Add_new_paint_layer = countActionsUse("Add_new_paint_layer")
	agreagtedData.Actions.Clear = countActionsUse("clear")
	agreagtedData.Actions.Copy_layer_clipboard = countActionsUse("copy_layer_clipboard")
	agreagtedData.Actions.Cut_layer_clipboard = countActionsUse("сut_layer_clipboard")
	agreagtedData.Actions.Edit_cut = countActionsUse("edit_cut")
	agreagtedData.Actions.Edit_redo = countActionsUse("edit_redo")
	agreagtedData.Actions.Edit_undo = countActionsUse("edit_undo")
	agreagtedData.Actions.File_new = countActionsUse("file_new")
	agreagtedData.Actions.Fill_selection_background_color = countActionsUse("fill_selection_background_color")
	agreagtedData.Actions.Fill_selection_foreground_color = countActionsUse("fill_selection_foreground_color")
	agreagtedData.Actions.Fill_selection_pattern = countActionsUse("fill_selection_pattern")
	agreagtedData.Actions.Paste_at = countActionsUse("paste_at")
	agreagtedData.Actions.Stroke_selection = countActionsUse("stroke_selection")
	agreagtedData.Actions.View_show_canvas_only = countActionsUse("view_show_canvas_only")

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
func Agregated() string {
	out, _ := json.Marshal(agreagtedData)
	return string(out)
}
func AgregateToolsInfo() {
}
func GetAgregatedData(dataType string) md.CollectedData {

	switch dataType {
	case "install":
		return agreagtedData
	case "tools":
	default:
		var err md.CustomError
		err.Message = "Wrong request"
	}
	return agreagtedData
}
