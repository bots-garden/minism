#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o read-user-input.wasm \
  -target wasi main.go

ls -lh *.wasm
