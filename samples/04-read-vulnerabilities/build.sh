#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o gitlab-sast.wasm \
  -target wasi main.go

ls -lh *.wasm
