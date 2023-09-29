#!/bin/bash
go run ../../main.go \
call gitlab-sast.wasm \
report \
--input "gl-sast-report.json" \
--log-level debug \
--allow-paths '{"data":"/mnt"}'

#--allow-paths '{".":"/mnt"}'

echo ""
