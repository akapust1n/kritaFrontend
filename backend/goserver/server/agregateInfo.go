package server

import (
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

func AgregateInstalInfo() {
	var result md.CollectedData
	c := Session.DB("telemetry").C("installInfo")
	//Some error takes place. We lose a little bit of data
	countRecords := getFloat64(c.Find(bson.M{}).Count())

	//compiler
	result.Compiler.Type.GCC = getFloat64(c.Find(bson.M{"compiler.type": "GCC"}).Count())
	result.Compiler.Type.Clang = getFloat64(c.Find(bson.M{"compiler.type": "Clang"}).Count())
	result.Compiler.Type.MSVC = getFloat64(c.Find(bson.M{"compiler.type": "MSVC"}).Count())
	result.Compiler.Type.Other = countRecords - result.Compiler.Type.GCC - result.Compiler.Type.Clang - result.Compiler.Type.MSVC
	result.Compiler.Type.Other = checkOtherCount(result.Compiler.Type.Other)
	//os
	result.Platform.Os.Linux = getFloat64(c.Find(bson.M{"platform.os": "linux"}).Count())
	result.Platform.Os.Windows = getFloat64(c.Find(bson.M{"platform.os": "windows"}).Count())
	result.Platform.Os.Mac = getFloat64(c.Find(bson.M{"platform.os": "mac"}).Count())
	result.Platform.Os.Other = countRecords - result.Platform.Os.Linux - result.Platform.Os.Windows - result.Platform.Os.Mac
	result.Platform.Os.Other = checkOtherCount(result.Platform.Os.Other)
	fmt.Println("LINUX")
	fmt.Println(result.Platform.Os.Linux)
	fmt.Println(result.Platform.Os.Other)

	//version os windows
	result.Platform.Version.Windows.V7 = getFloat64(c.Find(bson.M{"platform.version": "7"}).Count())
	result.Platform.Version.Windows.V8 = getFloat64(c.Find(bson.M{"platform.version": "8"}).Count())
	result.Platform.Version.Windows.V81 = getFloat64(c.Find(bson.M{"platform.version": "8.1"}).Count())
	result.Platform.Version.Windows.V10 = getFloat64(c.Find(bson.M{"platform.version": "10"}).Count())
	result.Platform.Version.Windows.Other = result.Platform.Os.Windows - result.Platform.Version.Windows.V7 - result.Platform.Version.Windows.V8 - result.Platform.Version.Windows.V81 - result.Platform.Version.Windows.V10
	result.Platform.Version.Windows.Other = checkOtherCount(result.Platform.Version.Windows.Other)

	//version os linux
	//не уверен, что именно так пишется нужные версии linux
	result.Platform.Version.Linux.Ubuntu1404 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.04"}).Count())
	result.Platform.Version.Linux.Ubuntu1410 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.10"}).Count())
	result.Platform.Version.Linux.Ubuntu1504 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.04"}).Count())
	result.Platform.Version.Linux.Ubuntu1510 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.10"}).Count())
	result.Platform.Version.Linux.Ubuntu1604 = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-16.04"}).Count())
	result.Platform.Version.Linux.Other = result.Platform.Os.Linux - result.Platform.Version.Linux.Ubuntu1404 - result.Platform.Version.Linux.Ubuntu1410 - result.Platform.Version.Linux.Ubuntu1504 - result.Platform.Version.Linux.Ubuntu1510 - result.Platform.Version.Linux.Ubuntu1604 - result.Platform.Version.Linux.Ubuntu1610 - result.Platform.Version.Linux.Ubuntu1704
	result.Platform.Version.Linux.Other = checkOtherCount(result.Platform.Version.Linux.Other)

	result.Platform.Version.Mac.V1012 = getFloat64(c.Find(bson.M{"platform.version": "10.12"}).Count())
	result.Platform.Version.Mac.Other = result.Platform.Os.Mac - result.Platform.Version.Mac.V1012
	result.Platform.Version.Mac.Other = checkOtherCount(result.Platform.Version.Linux.Other)

	//cpu
	result.CPU.Architecture.X86_64 = getFloat64(c.Find(bson.M{"cpu.architecture": "x86_64"}).Count())
	result.CPU.Architecture.X86 = getFloat64(c.Find(bson.M{"cpu.architecture": "i386"}).Count())
	result.CPU.Architecture.Other = countRecords - result.CPU.Architecture.X86_64 - result.CPU.Architecture.X86
	result.CPU.Architecture.Other = checkOtherCount(result.CPU.Architecture.Other)

	result.CPU.Cores.C1 = getFloat64(c.Find(bson.M{"cpu.count": "1"}).Count())
	result.CPU.Cores.C2 = getFloat64(c.Find(bson.M{"cpu.count": "2"}).Count())
	result.CPU.Cores.C3 = getFloat64(c.Find(bson.M{"cpu.count": "3"}).Count())
	result.CPU.Cores.C4 = getFloat64(c.Find(bson.M{"cpu.count": "4"}).Count())
	result.CPU.Cores.C6 = getFloat64(c.Find(bson.M{"cpu.count": "6"}).Count())
	result.CPU.Cores.C8 = getFloat64(c.Find(bson.M{"cpu.count": "8"}).Count())
	result.CPU.Cores.Other = countRecords - result.CPU.Cores.C1 - result.CPU.Cores.C2 - result.CPU.Cores.C3 - result.CPU.Cores.C4 - result.CPU.Cores.C6 - result.CPU.Cores.C8

	result.Locale.Language.English = getFloat64(c.Find(bson.M{"locale.language": "English"}).Count())
	result.Locale.Language.Russian = getFloat64(c.Find(bson.M{"locale.language": "Russian"}).Count())
	result.Locale.Language.Other = countRecords - result.Locale.Language.English - result.Locale.Language.Russian
	result.Locale.Language.Other = checkOtherCount(result.Locale.Language.Other)
	//finish
	agreagtedData = result
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
