package main

import (
	"github.com/extism/go-pdk"
)

//export say_hello
func say_hello()  {

	// read function argument from the memory
	input := pdk.Input()

	pdk.Log(pdk.LogInfo, "😀😃😄")

	firstName, _ := pdk.GetConfig("firstName")
	lastName, _ := pdk.GetConfig("lastName")

	pdk.Log(pdk.LogInfo, firstName)
	pdk.Log(pdk.LogInfo, lastName)

	output := "👋 (From Go) Hello " + string(input)

	mem := pdk.AllocateString(output)
	// copy output to host memory
	pdk.OutputMemory(mem)

	//return 0
}

func main() {}
