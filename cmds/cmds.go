package cmds

import (
	_ "embed"
	"flag"
	"fmt"
	"minism/minism"
)

//go:embed version.txt
var version []byte

/*
go run main.go \
call ../01-simple-go-plugin/simple.wasm \
say_hello \
  --input "Bob Morane" \
  --log-level info \
  --allow-hosts *,*.google.com,yo.com \

*/

func Parse(command string, args []string) error {
	//fmt.Println("Command:", command)
	//fmt.Println("Args:", args)
	switch command {

	case "call":

		wasmFilePath := flag.Args()[1]     // path of the wasm file
		wasmFunctionName := flag.Args()[2] // function name

		flagSet := flag.NewFlagSet("call", flag.ExitOnError)

		input := flagSet.String("input", "", "input")
		logLevel := flagSet.String("log-level", "", "Log level")
		allowHosts := flagSet.String("allow-hosts", `"[*]"`, "use a json array")
		allowPaths := flagSet.String("allow-paths", "{}", "use a json string to define the allowed paths")
		config := flagSet.String("config", "{}", "use a json string to define the config data")
		wasi := flagSet.Bool("wasi", true, "")

		wasmUrl := flagSet.String("wasm-url", "", "url to download the wasm file")
		authHeaderName := flagSet.String("auth-header-name", "", "ex: PRIVATE-TOKEN")
		authHeaderValue := flagSet.String("auth-header-value", "", "ex: IlovePandas")

		flagSet.Parse(args[2:])

		minism.Execute(minism.WasmArguments{
			FilePath:        wasmFilePath,
			FunctionName:    wasmFunctionName,
			Input:           *input,
			LogLevel:        *logLevel,
			AllowHosts:      *allowHosts,
			AllowPaths:      *allowPaths,
			Config:          *config,
			Wasi:            *wasi,
			Url:             *wasmUrl,
			AuthHeaderName:  *authHeaderName,
			AuthHeaderValue: *authHeaderValue,
		})
		return nil

	case "version":
		fmt.Println(string(version))
		//os.Exit(0)
		return nil

	default:
		return fmt.Errorf("🔴 invalid command")
	}
}
