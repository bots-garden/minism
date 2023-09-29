package main

import (
	"github.com/extism/go-pdk"
)

//export hostPrint
func hostPrint(offset uint64) uint64

func Print(message string) {
	messageMemory := pdk.AllocateString(message)
	hostPrint(messageMemory.Offset())
}

//export hostInput
func hostInput(offset uint64) uint64

func Input(prompt string) string {
	promptMemory := pdk.AllocateString(prompt)
	offset := hostInput(promptMemory.Offset())

	memoryResult := pdk.FindMemory(offset)
	buffResult := make([]byte, memoryResult.Length())
	memoryResult.Load(buffResult)

	return string(buffResult)
}

//export run
func run() {

	name := Input("What's your name? ")
	Print("Hello " + name)

}

func main() {}
