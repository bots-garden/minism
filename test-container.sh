#!/bin/bash
set -o allexport; source release.env; set +o allexport

docker rmi -f ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG}

docker run \
    -v $(pwd)/samples/01-simple-go-plugin:/app \
    --rm ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG}  \
    ./minism call ./app/simple.wasm say_hello \
    --input "John Doe" \
    --log-level info \
    --config '{"firstName":"Bob","lastName":"Morane"}'

docker run \
    -v $(pwd)/samples/02-ready-to-use-host-functions:/app \
    --rm ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG}  \
    ./minism call ./app/host-functions.wasm say_hello \
    --input "😀 Hello World 🌍! (from TinyGo)" \
    --log-level info \
    --allow-hosts '["*", "jsonplaceholder.typicode.com"]' \
    --config '{"route":"https://jsonplaceholder.typicode.com/todos/1"}'
