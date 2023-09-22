package minism

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"minism/levels"
	"os"

	"github.com/extism/extism"
	"github.com/tetratelabs/wazero"

	"github.com/go-resty/resty/v2"
)

func getHostsFromString(allowHosts string) []string {
	var hosts []string
	unmarshallError := json.Unmarshal([]byte(allowHosts), &hosts)
	if unmarshallError != nil {
		fmt.Println(unmarshallError)
		os.Exit(1)
	}
	return hosts

}

func getPathsFromJsonString(allowPaths string) map[string]string {
	var paths map[string]string
	unmarshallError := json.Unmarshal([]byte(allowPaths), &paths)
	if unmarshallError != nil {
		fmt.Println(unmarshallError)
		os.Exit(1)
	}
	return paths
}

func getConfigFromJsonString(config string) map[string]string {
	var manifestConfig map[string]string
	unmarshallError := json.Unmarshal([]byte(config), &manifestConfig)
	if unmarshallError != nil {
		fmt.Println(unmarshallError)
		os.Exit(1)
	}
	return manifestConfig
}

func downloadWasmFile(wasmArgs WasmArguments) error {
	// authenticationHeader:
	// Example: "PRIVATE-TOKEN: ${GITLAB_WASM_TOKEN}"
	client := resty.New()

	if wasmArgs.AuthHeaderName != "" {
		client.SetHeader(wasmArgs.AuthHeaderName, wasmArgs.AuthHeaderValue)
	}

	resp, err := client.R().
		SetOutput(wasmArgs.FilePath).
		Get(wasmArgs.Url)

	if resp.IsError() {
		return errors.New("error while downloading the wasm file")
	}

	if err != nil {
		return err
	}
	return nil
}

func Execute(wasmArgs WasmArguments) {

	if wasmArgs.Url != "" { // we need to download the wasm file
		fmt.Println("🌍 downloading...", wasmArgs.Url)
		err := downloadWasmFile(wasmArgs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	hosts := getHostsFromString(wasmArgs.AllowHosts)
	paths := getPathsFromJsonString(wasmArgs.AllowPaths)
	manifestConfig := getConfigFromJsonString(wasmArgs.Config)

	ctx := context.Background()

	level := levels.GetLevel(wasmArgs.LogLevel)

	extismConfig := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   wasmArgs.Wasi,
		LogLevel:     &level,
	}

	/*
		type Manifest struct {
			Wasm   []Wasm `json:"wasm"`
			Memory struct {
				MaxPages uint32 `json:"max_pages,omitempty"`
			} `json:"memory,omitempty"`
			Config       map[string]string `json:"config,omitempty"`
			AllowedHosts []string          `json:"allowed_hosts,omitempty"`
			AllowedPaths map[string]string `json:"allowed_paths,omitempty"`
			Timeout      time.Duration     `json:"timeout_ms,omitempty"`
		}
	*/

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmArgs.FilePath},
		},
		AllowedHosts: hosts,
		AllowedPaths: paths,
		Config:       manifestConfig,
	}

	// get an instance of the wasm Extism plugin
	wasmPlugin, err := extism.NewPlugin(ctx, manifest, extismConfig, nil)

	if err != nil {
		panic(err)
	}

	// rc: result code : if 1 error with the wasm module

	_, res, err := wasmPlugin.Call(
		wasmArgs.FunctionName,
		[]byte(wasmArgs.Input),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(res))
		os.Exit(0)
	}

}
