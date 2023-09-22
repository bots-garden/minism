#!/bin/bash
go run ../../main.go \
call files.wasm \
createFile \
--log-level info \
--allow-paths '{"data":"/mnt"}'
