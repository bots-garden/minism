#!/bin/bash
go run ../../main.go \
call host-functions.wasm \
say_hello \
--input "😀 Hello World 🌍! (from TinyGo)" \
--log-level info \
--allow-hosts '["*", "jsonplaceholder.typicode.com"]' \
--config '{"route":"https://jsonplaceholder.typicode.com/todos/1"}'

echo ""
