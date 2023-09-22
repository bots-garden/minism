module minism

go 1.20

require (
	github.com/extism/extism v0.4.0
	github.com/tetratelabs/wazero v1.3.0
)

require golang.org/x/net v0.15.0 // indirect

require (
	github.com/go-resty/resty/v2 v2.8.0
	github.com/gobwas/glob v0.2.3 // indirect
)

replace github.com/extism/extism => ./go-sdk
