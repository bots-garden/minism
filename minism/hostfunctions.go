package minism

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero/api"
)

/*
------------------------------------------------------------
How to use the `print` host function (from an Extism plugin)
------------------------------------------------------------

```golang
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

//export say_hello
func say_hello() {
	input := pdk.Input()
	Print("hello " + string(input))
}

func main() {}
```

```rust
#![no_main]

use extism_pdk::*;

extern "C" {
    fn hostPrint(ptr: u64) -> u64;
}

pub fn print_message(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrint(memory_text.offset) };
}

#[plugin_fn]
pub fn hello(input: String) -> FnResult<u64> {
    let msg: String = "🦀 Hello ".to_string() + &input;
    print_message(msg);
    Ok(0)
}
```
*/

func print(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	offset := stack[0]
	bufferInput, err := plugin.ReadBytes(offset)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	message := string(bufferInput)
	fmt.Print(message)

	stack[0] = 0
}

func println(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	offset := stack[0]
	bufferInput, err := plugin.ReadBytes(offset)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	message := string(bufferInput)
	fmt.Println(message)

	stack[0] = 0
}

func input(ctx context.Context, plugin *extism.CurrentPlugin, stack []uint64) {
	offset := stack[0]
	bufferInput, err := plugin.ReadBytes(offset)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	prompt := string(bufferInput)
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")

	// return the value
	offset, errReturn := plugin.WriteString(input)
	if errReturn != nil {
		fmt.Println(errReturn.Error())
		panic(errReturn)
	}
	stack[0] = offset

}

func InitializeHostFunctions() []extism.HostFunction {

	hostPrint := extism.NewHostFunctionWithStack(
		"hostPrint",
		"env",
		print,
		[]api.ValueType{api.ValueTypeI64},
		[]api.ValueType{api.ValueTypeI64},
	)

	hostPrintln := extism.NewHostFunctionWithStack(
		"hostPrintln",
		"env",
		println,
		[]api.ValueType{api.ValueTypeI64},
		[]api.ValueType{api.ValueTypeI64},
	)

	hostInput := extism.NewHostFunctionWithStack(
		"hostInput",
		"env",
		input,
		[]api.ValueType{api.ValueTypeI64},
		[]api.ValueType{api.ValueTypeI64},
	)

	return []extism.HostFunction{
		hostPrint,
		hostPrintln,
		hostInput,
	}

}
