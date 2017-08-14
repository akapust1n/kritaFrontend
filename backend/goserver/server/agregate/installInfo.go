package agregate

import (
	serv "kritaServers/backend/goserver/server"

	"gopkg.in/mgo.v2/bson"
)

func AgregateInstalInfo() {
	c := serv.Session.DB("telemetry").C("installInfo")
	//Some error takes place. We lose a little bit of data
	countRecords := getFloat64(c.Find(bson.M{}).Count())

	//compiler
	agreagtedData.Compiler.Type.GCC.Count = getFloat64(c.Find(bson.M{"compiler.type": "GCC"}).Count())
	agreagtedData.Compiler.Type.Clang.Count = getFloat64(c.Find(bson.M{"compiler.type": "Clang"}).Count())
	agreagtedData.Compiler.Type.MSVC.Count = getFloat64(c.Find(bson.M{"compiler.type": "MSVC"}).Count())
	existsRecords := countExist("compiler.type", c)
	agreagtedData.Compiler.Type.Other.Count = existsRecords - agreagtedData.Compiler.Type.GCC.Count - agreagtedData.Compiler.Type.Clang.Count - agreagtedData.Compiler.Type.MSVC.Count
	agreagtedData.Compiler.Type.Other.Count = checkOtherCount(agreagtedData.Compiler.Type.Other.Count)
	agreagtedData.Compiler.Type.Unknown.Count = countRecords - existsRecords

	agreagtedData.Compiler.Type.GCC.Proportion = getProportion(agreagtedData.Compiler.Type.GCC.Count, existsRecords)
	agreagtedData.Compiler.Type.Clang.Proportion = getProportion(agreagtedData.Compiler.Type.Clang.Count, existsRecords)
	agreagtedData.Compiler.Type.MSVC.Proportion = getProportion(agreagtedData.Compiler.Type.MSVC.Count, existsRecords)
	agreagtedData.Compiler.Type.Other.Proportion = getProportion(agreagtedData.Compiler.Type.Other.Count, existsRecords)

	//os
	agreagtedData.Platform.Os.Linux.Count = getFloat64(c.Find(bson.M{"platform.os": "linux"}).Count())
	agreagtedData.Platform.Os.Windows.Count = getFloat64(c.Find(bson.M{"platform.os": "windows"}).Count())
	agreagtedData.Platform.Os.Mac.Count = getFloat64(c.Find(bson.M{"platform.os": "mac"}).Count())
	existsRecords = countExist("platform.os", c)
	agreagtedData.Platform.Os.Other.Count = existsRecords - agreagtedData.Platform.Os.Linux.Count - agreagtedData.Platform.Os.Windows.Count - agreagtedData.Platform.Os.Mac.Count
	agreagtedData.Platform.Os.Other.Count = checkOtherCount(agreagtedData.Platform.Os.Other.Count)
	agreagtedData.Platform.Os.Unknown.Count = countRecords - existsRecords

	agreagtedData.Platform.Os.Linux.Proportion = getProportion(agreagtedData.Platform.Os.Linux.Count, existsRecords)
	agreagtedData.Platform.Os.Windows.Proportion = getProportion(agreagtedData.Platform.Os.Windows.Count, existsRecords)
	agreagtedData.Platform.Os.Mac.Proportion = getProportion(agreagtedData.Platform.Os.Mac.Count, existsRecords)
	agreagtedData.Platform.Os.Other.Proportion = getProportion(agreagtedData.Platform.Os.Other.Count, existsRecords)
	agreagtedData.Platform.Os.Unknown.Proportion = getProportion(agreagtedData.Platform.Os.Unknown.Count, existsRecords)
	//version os windows
	agreagtedData.Platform.Version.Windows.V7.Count = getFloat64(c.Find(bson.M{"platform.version": "7"}).Count())
	agreagtedData.Platform.Version.Windows.V8.Count = getFloat64(c.Find(bson.M{"platform.version": "8"}).Count())
	agreagtedData.Platform.Version.Windows.V81.Count = getFloat64(c.Find(bson.M{"platform.version": "8.1"}).Count())
	agreagtedData.Platform.Version.Windows.V10.Count = getFloat64(c.Find(bson.M{"platform.version": "10"}).Count())
	agreagtedData.Platform.Version.Windows.Other.Count = agreagtedData.Platform.Os.Windows.Count - agreagtedData.Platform.Version.Windows.V7.Count - agreagtedData.Platform.Version.Windows.V8.Count - agreagtedData.Platform.Version.Windows.V81.Count - agreagtedData.Platform.Version.Windows.V10.Count
	agreagtedData.Platform.Version.Windows.Other.Count = checkOtherCount(agreagtedData.Platform.Version.Windows.Other.Count)

	agreagtedData.Platform.Version.Windows.V7.Proportion = getProportion(agreagtedData.Platform.Version.Windows.V7.Count, agreagtedData.Platform.Os.Windows.Count)
	agreagtedData.Platform.Version.Windows.V8.Proportion = getProportion(agreagtedData.Platform.Version.Windows.V8.Count, agreagtedData.Platform.Os.Windows.Count)
	agreagtedData.Platform.Version.Windows.V81.Proportion = getProportion(agreagtedData.Platform.Version.Windows.V81.Count, agreagtedData.Platform.Os.Windows.Count)
	agreagtedData.Platform.Version.Windows.V10.Proportion = getProportion(agreagtedData.Platform.Version.Windows.V10.Count, agreagtedData.Platform.Os.Windows.Count)
	agreagtedData.Platform.Version.Windows.Other.Proportion = getProportion(agreagtedData.Platform.Version.Windows.Other.Count, agreagtedData.Platform.Os.Windows.Count)

	//version os linux
	//не уверен, что именно так пишется нужные версии linux
	agreagtedData.Platform.Version.Linux.Ubuntu1404.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.04"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1410.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-14.10"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1504.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.04"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1510.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-15.10"}).Count())
	agreagtedData.Platform.Version.Linux.Ubuntu1604.Count = getFloat64(c.Find(bson.M{"platform.version": "ubuntu-16.04"}).Count())
	agreagtedData.Platform.Version.Linux.Other.Count = agreagtedData.Platform.Os.Linux.Count - agreagtedData.Platform.Version.Linux.Ubuntu1404.Count - agreagtedData.Platform.Version.Linux.Ubuntu1410.Count - agreagtedData.Platform.Version.Linux.Ubuntu1504.Count - agreagtedData.Platform.Version.Linux.Ubuntu1510.Count - agreagtedData.Platform.Version.Linux.Ubuntu1604.Count - agreagtedData.Platform.Version.Linux.Ubuntu1610.Count - agreagtedData.Platform.Version.Linux.Ubuntu1704.Count
	agreagtedData.Platform.Version.Linux.Other.Count = checkOtherCount(agreagtedData.Platform.Version.Linux.Other.Count)

	agreagtedData.Platform.Version.Linux.Ubuntu1404.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1404.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Ubuntu1410.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1410.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Ubuntu1504.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1504.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Ubuntu1510.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1510.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Ubuntu1604.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1604.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Ubuntu1610.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1610.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Ubuntu1704.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Ubuntu1704.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Linux.Other.Proportion = getProportion(agreagtedData.Platform.Version.Linux.Other.Count, agreagtedData.Platform.Os.Linux.Count)

	//mac version
	agreagtedData.Platform.Version.Mac.V1012.Count = getFloat64(c.Find(bson.M{"platform.version": "10.12"}).Count())
	agreagtedData.Platform.Version.Mac.Other.Count = agreagtedData.Platform.Os.Mac.Count - agreagtedData.Platform.Version.Mac.V1012.Count
	agreagtedData.Platform.Version.Mac.Other.Count = checkOtherCount(agreagtedData.Platform.Version.Linux.Other.Count)

	agreagtedData.Platform.Version.Mac.V1012.Proportion = getProportion(agreagtedData.Platform.Version.Mac.V1012.Count, agreagtedData.Platform.Os.Linux.Count)
	agreagtedData.Platform.Version.Mac.Other.Proportion = getProportion(agreagtedData.Platform.Version.Mac.Other.Count, agreagtedData.Platform.Os.Linux.Count)

	//cpu
	agreagtedData.CPU.Architecture.X86_64.Count = getFloat64(c.Find(bson.M{"cpu.architecture": "x86_64"}).Count())
	agreagtedData.CPU.Architecture.X86.Count = getFloat64(c.Find(bson.M{"cpu.architecture": "i386"}).Count())
	existsRecords = countExist("cpu.architecture", c)

	agreagtedData.CPU.Architecture.Other.Count = existsRecords - agreagtedData.CPU.Architecture.X86_64.Count - agreagtedData.CPU.Architecture.X86.Count
	agreagtedData.CPU.Architecture.Other.Count = checkOtherCount(agreagtedData.CPU.Architecture.Other.Count)
	agreagtedData.CPU.Architecture.Unknown.Count = countRecords - existsRecords

	agreagtedData.CPU.Architecture.X86_64.Proportion = getProportion(agreagtedData.CPU.Architecture.X86_64.Count, existsRecords)
	agreagtedData.CPU.Architecture.X86.Proportion = getProportion(agreagtedData.CPU.Architecture.X86.Count, existsRecords)
	agreagtedData.CPU.Architecture.Other.Proportion = getProportion(agreagtedData.CPU.Architecture.Other.Count, existsRecords)
	agreagtedData.CPU.Architecture.Unknown.Proportion = getProportion(agreagtedData.CPU.Architecture.Unknown.Count, countRecords)

	//cpu cores
	agreagtedData.CPU.Cores.C1.Count = getFloat64(c.Find(bson.M{"cpu.count": "1"}).Count())
	agreagtedData.CPU.Cores.C2.Count = getFloat64(c.Find(bson.M{"cpu.count": "2"}).Count())
	agreagtedData.CPU.Cores.C3.Count = getFloat64(c.Find(bson.M{"cpu.count": "3"}).Count())
	agreagtedData.CPU.Cores.C4.Count = getFloat64(c.Find(bson.M{"cpu.count": "4"}).Count())
	agreagtedData.CPU.Cores.C6.Count = getFloat64(c.Find(bson.M{"cpu.count": "6"}).Count())
	agreagtedData.CPU.Cores.C8.Count = getFloat64(c.Find(bson.M{"cpu.count": "8"}).Count())
	existsRecords = countExist("cpu.count", c)
	agreagtedData.CPU.Cores.Other.Count = existsRecords - agreagtedData.CPU.Cores.C1.Count - agreagtedData.CPU.Cores.C2.Count - agreagtedData.CPU.Cores.C3.Count - agreagtedData.CPU.Cores.C4.Count - agreagtedData.CPU.Cores.C6.Count - agreagtedData.CPU.Cores.C8.Count
	agreagtedData.CPU.Cores.Unknown.Count = countRecords - existsRecords

	agreagtedData.CPU.Cores.C1.Proportion = getProportion(agreagtedData.CPU.Cores.C1.Count, existsRecords)
	agreagtedData.CPU.Cores.C2.Proportion = getProportion(agreagtedData.CPU.Cores.C2.Count, existsRecords)
	agreagtedData.CPU.Cores.C3.Proportion = getProportion(agreagtedData.CPU.Cores.C3.Count, existsRecords)
	agreagtedData.CPU.Cores.C4.Proportion = getProportion(agreagtedData.CPU.Cores.C4.Count, existsRecords)
	agreagtedData.CPU.Cores.C6.Proportion = getProportion(agreagtedData.CPU.Cores.C6.Count, existsRecords)
	agreagtedData.CPU.Cores.C8.Proportion = getProportion(agreagtedData.CPU.Cores.C8.Count, existsRecords)
	agreagtedData.CPU.Cores.Other.Proportion = getProportion(agreagtedData.CPU.Cores.Other.Count, existsRecords)
	agreagtedData.CPU.Cores.Unknown.Proportion = getProportion(agreagtedData.CPU.Cores.Unknown.Count, countRecords)

	//language
	agreagtedData.Locale.Language.English.Count = getFloat64(c.Find(bson.M{"locale.language": "English"}).Count())
	agreagtedData.Locale.Language.Russian.Count = getFloat64(c.Find(bson.M{"locale.language": "Russian"}).Count())
	existsRecords = countExist("locale.language", c)
	agreagtedData.Locale.Language.Other.Count = existsRecords - agreagtedData.Locale.Language.English.Count - agreagtedData.Locale.Language.Russian.Count
	agreagtedData.Locale.Language.Other.Count = checkOtherCount(agreagtedData.Locale.Language.Other.Count)
	agreagtedData.Locale.Language.Unknown.Count = countRecords - existsRecords

	agreagtedData.Locale.Language.English.Proportion = getProportion(agreagtedData.Locale.Language.English.Count, existsRecords)
	agreagtedData.Locale.Language.Russian.Proportion = getProportion(agreagtedData.Locale.Language.Russian.Count, existsRecords)
	agreagtedData.Locale.Language.Other.Proportion = getProportion(agreagtedData.Locale.Language.Other.Count, existsRecords)
	agreagtedData.Locale.Language.Unknown.Proportion = getProportion(agreagtedData.Locale.Language.Unknown.Count, countRecords)

}
