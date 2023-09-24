#!/bin/bash
go run ../../main.go \
call simple.wasm \
say_hello \
--input "Bob Morane" \
--log-level info \
--allow-hosts '["*","*.google.com","yo.com"]' \
--config '{"firstName":"Philippe","lastName":"Charrière"}'
