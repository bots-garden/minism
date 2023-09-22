package main

import (
	"strconv"

	"github.com/extism/go-pdk"
	"github.com/valyala/fastjson"
)

//export say_hello
func say_hello() int32 {

	// read function argument from the memory
	input := pdk.Input()

	pdk.Log(pdk.LogInfo, string(input))

	// https://jsonplaceholder.typicode.com/todos/3
	url, _ := pdk.GetConfig("route")

	// use request (host function)
	req := pdk.NewHTTPRequest("GET", url)
	res := req.Send()

	pdk.Log(pdk.LogInfo, "res status:" + strconv.FormatUint(uint64(res.Status()), 10))
	pdk.Log(pdk.LogInfo, "res body:" + string(res.Body()))

	parser := fastjson.Parser{}
	jsonValue, err := parser.Parse(string(res.Body()))
	if err != nil {
		pdk.Log(pdk.LogInfo, err.Error())
	}
	title := string(jsonValue.GetStringBytes("title"))

	output := "🌍 url: " + url + " 📝 title: " + title

	mem := pdk.AllocateString(output)
	// copy output to host memory
	pdk.OutputMemory(mem)

	return 0
}

func main() {}
