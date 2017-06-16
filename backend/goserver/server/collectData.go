package server

import (
	"encoding/json"
	"fmt"
	md "kritaServers/backend/goserver/server/models"
	"strconv"
)

func count(query string, whatCount string) float64 {
	rows, err := Db.Query(query, whatCount)
	checkErr(err)
	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
		fmt.Println("scanned!")

	}
	return float64(count) //golang хранит числа в json только в флоат64
}

// //ПЕРЕМЕННОЕ КОЛИЧЕСТВО АРГУМЕНТОВ СДЕЛАТЬ
// func countOther(query string, whatCount1 string, whatCount2 string) float64 {
// 	rows, err := Db.Query(query, whatCount1, whatCount2)
// 	checkErr(err)
// 	var count int

// 	for rows.Next() {
// 		err := rows.Scan(&count)
// 		checkErr(err)
// 		fmt.Println("scanned other platforms!")

// 	}
// 	return float64(count) //golang хранит числа в json только в флоат64
// }
func countOther(query string, queryNotEql string, args ...string) float64 {
	execQuery := query
	for i, _ := range args { //Зачем я использую препаред стейтмент, когда можно без него?
		if i != 0 {
			execQuery += " and "
		}
		//execQuery += queryNotEql + "'" + v + "'"
		execQuery += queryNotEql + "$" + strconv.Itoa(i+1)

	}
	//execQuery += ";"
	fmt.Println(execQuery)
	temp, _ := Db.Prepare(execQuery)
	fmt.Println(len(args))

	///rows, err := temp.Query(args)
	old := args
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}
	fmt.Println(new...)
	rows, err := temp.Query(new...)

	checkErr(err)

	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
		fmt.Println("scanned other abstruct platforms!")

	}
	return float64(count) //golang хранит числа в json только в флоат64
}

func CollectData() {
	var result md.CollectedData

	const osQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo where  data -> 'platform'->>'os' = $1"
	result.Platform.Os.Linux = count(osQuery, "linux")
	fmt.Println(result.Platform.Os.Linux)
	result.Platform.Os.Windows = count(osQuery, "windows")
	result.Platform.Os.Mac = count(osQuery, "mac")

	const osOtherQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo WHERE "
	const osOtherNotEql = "data -> 'platform'->>'os' !="
	result.Platform.Os.Other = countOther(osOtherQuery, osOtherNotEql, "linux", "windows", "mac")

	const osWindowsVersionQuery = "select count( data -> 'platform'->>'version') from generalInfo WHERE data -> 'platform'->>'os' ='windows' and data -> 'platform'->>'version'=$1"
	result.Platform.Version.Windows.V7 = count(osWindowsVersionQuery, "7")
	result.Platform.Version.Windows.V8 = count(osWindowsVersionQuery, "8")
	result.Platform.Version.Windows.V81 = count(osWindowsVersionQuery, "8.1")
	result.Platform.Version.Windows.V10 = count(osWindowsVersionQuery, "10")

	const osWindowsVersionOtherQuery = "select count( data -> 'platform'->>'version') from generalInfo WHERE data -> 'platform'->>'os' ='windows' "
	const osWindowsVersionOtherNotEql = " data -> 'platform'->>'version' != "
	result.Platform.Version.Windows.Other = countOther(osWindowsVersionOtherQuery, osWindowsVersionOtherNotEql, "7", "8", "8.1", "10")

	//не уверен, что именно так пишется нужные версии linux
	const osLinuxVersionQuery = "select count( data -> 'platform'->>'version') from generalInfo WHERE data -> 'platform'->>'os' ='linux' and data -> 'platform'->>'version'=$1"
	result.Platform.Version.Linux.Ubuntu1404 = count(osLinuxVersionQuery, "ubuntu-14.04")
	result.Platform.Version.Linux.Ubuntu1410 = count(osLinuxVersionQuery, "ubuntu-14.10")
	result.Platform.Version.Linux.Ubuntu1504 = count(osLinuxVersionQuery, "ubuntu-15.04")
	result.Platform.Version.Linux.Ubuntu1510 = count(osLinuxVersionQuery, "ubuntu-15.10")
	result.Platform.Version.Linux.Ubuntu1604 = count(osLinuxVersionQuery, "ubuntu-16.04")
	result.Platform.Version.Linux.Ubuntu1610 = count(osLinuxVersionQuery, "ubuntu-16.10")
	result.Platform.Version.Linux.Ubuntu1704 = count(osLinuxVersionQuery, "ubuntu-17.04")

	const osLinuxVersionOtherQuery = "select count( data -> 'platform'->>'version') from generalInfo WHERE data -> 'platform'->>'os' ='linux' "
	const osLinuxVersionOtherNotEql = " data -> 'platform'->>'version' != "
	result.Platform.Version.Linux.Other = countOther(osLinuxVersionOtherQuery, osLinuxVersionOtherNotEql, "ubuntu-14.04", "ubuntu-14.10", "ubuntu-15.04", "ubuntu-15.10", "ubuntu-16.04", "ubuntu-16.10", "ubuntu-17.04")

	const osMacVersionQuery = "select count( data -> 'platform'->>'version') from generalInfo WHERE data -> 'platform'->>'os' ='mac' and data -> 'platform'->>'version'=$1"
	result.Platform.Version.Mac.V1012 = count(osMacVersionQuery, "10.12")

	const osMacVersionOtherQuery = "select count( data -> 'platform'->>'version') from generalInfo WHERE data -> 'platform'->>'os' ='mac' and data -> 'platform'->>'version'=$1"
	const osMacVersionOtherNotEql = " data -> 'platform'->>'version' != "
	result.Platform.Version.Mac.Other = countOther(osMacVersionOtherQuery, osMacVersionOtherNotEql, "10.12")

	//const osOtherQuery = "select count( data -> 'platform'->>'os') as os  from generalInfo where  data -> 'platform'->>'os' != $1 and data -> 'platform'->>'os'!=$2"
	//result.Platform.Os.Other = countOther(osOtherQuery, "windows", "linux")

	const archCPUQuery = "select count( data -> 'cpu'->>'architecture')   from generalInfo where data -> 'cpu'->>'architecture' = $1"
	result.CPU.Architecture.X86_64 = count(archCPUQuery, "x86_64")
	result.CPU.Architecture.X86 = count(archCPUQuery, "i386") // может быть что-то другое нужно
	const archCPUOtherQuery = "select count( data -> 'cpu'->>'architecture')   from generalInfo where  "
	const archCPUOtherNotEql = "data -> 'cpu'->>'architecture' !="
	result.CPU.Architecture.Other = countOther(archCPUOtherQuery, archCPUOtherNotEql, "x86_64", "i386")

	const coreCountQuery = "select data -> 'cpu'->>'count'   from generalInfo  where data -> 'cpu'->>'count' = $1"
	result.CPU.Cores.C1 = count(coreCountQuery, "1")
	result.CPU.Cores.C2 = count(coreCountQuery, "2")
	result.CPU.Cores.C3 = count(coreCountQuery, "3")
	result.CPU.Cores.C4 = count(coreCountQuery, "4")
	result.CPU.Cores.C6 = count(coreCountQuery, "6")
	result.CPU.Cores.C8 = count(coreCountQuery, "8")

	const coreCountOtherQuery = "select count( data -> 'cpu'->>'count')   from generalInfo where "
	const coreCountOtherNotEql = "data -> 'cpu'->>'count' !="
	result.CPU.Cores.Other = countOther(coreCountOtherQuery, coreCountOtherNotEql, "1", "2", "3", "4", "6", "8")

	const compilerTypeQuery = "select count(data->'compiler'->>'type') from generalInfo where data->'compiler'->>'type' = $1"
	result.Compiler.Type.GCC = count(compilerTypeQuery, "GCC")
	result.Compiler.Type.Clang = count(compilerTypeQuery, "Clang")
	result.Compiler.Type.MSVC = count(compilerTypeQuery, "MSVC")

	const compilerTypeOtherQuery = "select count(data->'compiler'->>'type') from generalInfo where "
	const compilerTypeOtherNotEql = "data->'compiler'->>'type' !="
	result.Compiler.Type.Other = countOther(compilerTypeOtherQuery, compilerTypeOtherNotEql, "GCC", "Clang", "MSVC")

	const localeLanguageQuery = "select count(data->'locale'->>'language') from generalInfo where data->'locale'->>'language' = $1"
	result.Locale.Language.English = count(localeLanguageQuery, "English")
	result.Locale.Language.Russian = count(localeLanguageQuery, "Russian")

	const localeLanguageOtherQuery = "select count(data->'locale'->>'language') from generalInfo where "
	const localeLanguageOtherNotEql = "data->'locale'->>'language' !="
	result.Locale.Language.Other = countOther(localeLanguageOtherQuery, localeLanguageOtherNotEql, "English", "Russian")

	fmt.Println("BEFORE JSON")

	const insertQuery = "INSERT into agregatedInfo(data)  VALUES($1)"
	b, err := json.Marshal(result)
	fmt.Println("SUCCESS JSON")
	fmt.Println(b)
	checkErr(err)
	_, err = Db.Exec(insertQuery, b)
	checkErr(err)
}
