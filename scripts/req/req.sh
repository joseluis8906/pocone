#!/usr/bin/env bash

curl -X POST -H "Content-Type: application/json" -d "@$1" http://localhost:8080/json-rpc
