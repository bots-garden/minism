package main

import (
	"os"
	"strconv"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export __wasm_call_ctors
func __wasm_call_ctors()

//export _initialize
func _initialize() {
	__wasm_call_ctors()
}

//export hostPrintln
func hostPrintln(offset uint64) uint64

func Println(message string) {
	messageMemory := pdk.AllocateString(message)
	hostPrintln(messageMemory.Offset())
}

func truncate(s []byte, to int) []byte {
	return s[:to]
}

//export report
func report() {
	fileName := pdk.InputString()

	// Read the file and get its contents as a byte slice
	content, err := os.ReadFile("/mnt/" + fileName)
	if err != nil {
		pdk.Log(pdk.LogDebug, err.Error())
	}

	parser := fastjson.Parser{}
	jsonValue, err := parser.Parse(string(content))
	if err != nil {
		pdk.Log(pdk.LogDebug, err.Error())
	}
	version := string(jsonValue.GetStringBytes("version"))

	Println("version: "+version)

	vulnerabilities := jsonValue.GetArray("vulnerabilities")

	for i := range vulnerabilities {
		Println("----------------------------------------------------------------------------")

		Println("id: "+string(vulnerabilities[i].GetStringBytes("id")))
		Println("category: "+string(vulnerabilities[i].GetStringBytes("category")))
		Println("name: "+string(vulnerabilities[i].GetStringBytes("name")))

		description := vulnerabilities[i].GetStringBytes("description")
		//shortDescription := truncate(description, 30)
		//pdk.Log(pdk.LogInfo, "📝: " + strconv.Itoa(len((description))))

		Println("description:\n\r"+string(description))
		Println("\n\r")

		Println("cve: "+string(vulnerabilities[i].GetStringBytes("cve")))
		Println("severity: "+string(vulnerabilities[i].GetStringBytes("severity")))

		scanner := vulnerabilities[i].GetObject("scanner")
		location := vulnerabilities[i].GetObject("location")

		Println("  scanner.id: "+scanner.Get("id").String())
		Println("  scanner.name: "+scanner.Get("name").String())

		Println("  location.file: "+location.Get("file").String())
		startLine, err := location.Get("start_line").Int()
		if err != nil {
			pdk.Log(pdk.LogDebug, err.Error())
		}
		Println("  location.start_line: "+strconv.Itoa(startLine))

		Println( "  identifiers:")

		identifiers := vulnerabilities[i].GetArray("identifiers")

		for j := range identifiers {
			Println("    ----------------------------------------------------------------")
			Println("    identifier.type: "+string(identifiers[j].GetStringBytes("type")))
			Println("    identifier.name: "+string(identifiers[j].GetStringBytes("name")))
			Println("    identifier.value: "+string(identifiers[j].GetStringBytes("value")))
			Println("    identifier.url: "+string(identifiers[j].GetStringBytes("url")))
		}

	}

}

func main() {}
