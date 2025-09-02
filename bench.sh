#!/bin/bash

# host="go-http"
# host="elixir-phoenix"
host="js-adonisjs"
# host="python-fastapi"
# host="rust-axum"

# endpoint="/"
# endpoint="/db/"
endpoint="/chaos/"
# endpoint="/chaos-sync/"
# endpoint="/health/"

duration=10s
concurrency=100

declare -A host_map
host_map["go-http"]="http://go-http:8000$endpoint"
host_map["elixir-phoenix"]="http://elixir-phoenix:4000$endpoint"
host_map["js-adonisjs"]="http://js-adonisjs:3333$endpoint"
host_map["python-fastapi"]="http://python-fastapi:8080$endpoint"
host_map["rust-axum"]="http://rust-axum:3000$endpoint"


docker compose run benchmark \
    oha -z $duration -c $concurrency \
    --latency-correction --disable-keepalive \
    "${host_map[$host]}"
