#!/bin/bash
go run ../../main.go \
call gitlab-sast.wasm \
report \
--input "gl-sast-report.json" \
--log-level info \
--allow-paths '{"data":"/mnt"}'

echo ""
