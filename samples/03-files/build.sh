#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o files.wasm \
  -target wasi main.go

ls -lh *.wasm
