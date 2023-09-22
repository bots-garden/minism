package main

import (
	"os"

	"github.com/extism/go-pdk"
)

//export __wasm_call_ctors
func __wasm_call_ctors()

//export _initialize
func _initialize() {
	__wasm_call_ctors()
}

//export createFile
func createFile() {
	// Write to the file
	err := os.WriteFile("/mnt/test.txt", []byte("🎉 tada!"), 0644)
	if err != nil {
		pdk.Log(pdk.LogInfo, err.Error())
	}

	// Read the file and get its contents as a byte slice
	content, err := os.ReadFile("/mnt/test.txt")
	if err != nil {
		pdk.Log(pdk.LogInfo, err.Error())
	}
	pdk.Log(pdk.LogInfo, string(content))

}

func main() {}
