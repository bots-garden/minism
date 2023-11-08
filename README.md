# Minism

![Minism Gopher](minism.gopher.png "Minism Gopher")

A minimal(*) **CLI** to run the **[Extism](https://extism.org/)** plugins.

> (*): with zero dependency

Find the **official Extism CLI** here: https://github.com/extism/cli

## Install Minism

```bash
MINISM_VERSION="0.0.2"
MINISM_OS="linux" # or darwin
MINISM_ARCH="arm64" # or amd64
wget https://github.com/bots-garden/minism/releases/download/v${MINISM_VERSION}/minism-v${MINISM_VERSION}-${MINISM_OS}-${MINISM_ARCH}
cp minism-v${MINISM_VERSION}-${MINISM_OS}-${MINISM_ARCH} minism
chmod +x minism
rm minism-v${MINISM_VERSION}-${MINISM_OS}-${MINISM_ARCH}
sudo cp ./minism /usr/bin
rm minism
# check the version
minism version
```

## CLI Syntax

```text
Usage:
  minism [command] [arguments]

Available Commands:
  call        Call a plugin function
              Arguments: [wasm file path] [function name]
  version     Display the Minism version
              Arguments: nothing

Flags:
  --input             string   Argument of the function
  --log-level         string   Log level to print message
                               Possible values: error, warn, info, debug, trace
  --allow-hosts       string   Hosts for HTTP request (json array) 
                               Default: ["*"]
  --allow-paths       string   Allowed paths to write and read files (json string) 
                               Default: {}
  --config            string   Configuration data (json string)
                               Default: {}
  --wasi              bool     Default: true
  --wasm-url          string   Url to download the wasm file
  --auth-header-name  string   Authentication header name, ex: PRIVATE-TOKEN
  --auth-header-value string   Value of the authentication header, ex: IlovePandas  
```


> **Example**:
```bash
minism call simple.wasm say_hello \
  --input "Bob Morane" \
  --log-level info \
  --allow-hosts '["*","*.google.com"]' \
  --config '{"firstName":"John","lastName":"Doe"}' \
  --allow-paths '{"data":"/mnt"}'
```

## Docker image (arm64 + amd64)

> https://hub.docker.com/repository/docker/botsgarden/minism/general

You can use **Docker** to run **Minism**:

```bash
docker run \
    -v $(pwd)/samples/02-ready-to-use-host-functions:/app \
    --rm botsgarden/minism:0.0.2 \
    ./minism call ./app/host-functions.wasm say_hello \
    --input "😀 Hello World 🌍! (from TinyGo)" \
    --log-level info \
    --allow-hosts '["*", "jsonplaceholder.typicode.com"]' \
    --config '{"route":"https://jsonplaceholder.typicode.com/todos/1"}'
```
