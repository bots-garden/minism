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

	pdk.Log(pdk.LogDebug, "version: "+version)

	vulnerabilities := jsonValue.GetArray("vulnerabilities")

	for i := range vulnerabilities {
		pdk.Log(pdk.LogDebug, "----------------------------------------------------------------------------")

		pdk.Log(pdk.LogDebug, "id: "+string(vulnerabilities[i].GetStringBytes("id")))
		pdk.Log(pdk.LogDebug, "category: "+string(vulnerabilities[i].GetStringBytes("category")))
		pdk.Log(pdk.LogDebug, "name: "+string(vulnerabilities[i].GetStringBytes("name")))

		// Description could be too long
		description := vulnerabilities[i].GetStringBytes("description")
		//shortDescription := truncate(description, 30)

		//pdk.Log(pdk.LogInfo, "📝: " + strconv.Itoa(len((description))))

		//pdk.Log(pdk.LogInfo, "description: "+string(shortDescription))
		pdk.Log(pdk.LogDebug, "description:\n\r"+string(description))
		pdk.Log(pdk.LogDebug, "\n\r")

		pdk.Log(pdk.LogDebug, "cve: "+string(vulnerabilities[i].GetStringBytes("cve")))
		pdk.Log(pdk.LogDebug, "severity: "+string(vulnerabilities[i].GetStringBytes("severity")))

		scanner := vulnerabilities[i].GetObject("scanner")
		location := vulnerabilities[i].GetObject("location")

		pdk.Log(pdk.LogDebug, "  scanner.id: "+scanner.Get("id").String())
		pdk.Log(pdk.LogDebug, "  scanner.name: "+scanner.Get("name").String())

		pdk.Log(pdk.LogDebug, "  location.file: "+location.Get("file").String())
		startLine, err := location.Get("start_line").Int()
		if err != nil {
			pdk.Log(pdk.LogDebug, err.Error())
		}
		pdk.Log(pdk.LogDebug, "  location.start_line: "+strconv.Itoa(startLine))

		pdk.Log(pdk.LogDebug, "  identifiers:")

		identifiers := vulnerabilities[i].GetArray("identifiers")

		for j := range identifiers {
			pdk.Log(pdk.LogDebug, "    ----------------------------------------------------------------")
			pdk.Log(pdk.LogDebug, "    identifier.type: "+string(identifiers[j].GetStringBytes("type")))
			pdk.Log(pdk.LogDebug, "    identifier.name: "+string(identifiers[j].GetStringBytes("name")))
			pdk.Log(pdk.LogDebug, "    identifier.value: "+string(identifiers[j].GetStringBytes("value")))
			pdk.Log(pdk.LogDebug, "    identifier.url: "+string(identifiers[j].GetStringBytes("url")))
		}

	}

}

func main() {}
