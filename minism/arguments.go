package minism

// WasmArguments type
type WasmArguments struct {
	FilePath string
	FunctionName string
	Input string
	LogLevel string
	AllowHosts string
	AllowPaths string
	Config string
	Wasi bool
	Url string
	AuthHeaderName string
	AuthHeaderValue string
}
