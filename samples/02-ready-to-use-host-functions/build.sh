#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o host-functions.wasm \
  -target wasi main.go

ls -lh *.wasm
