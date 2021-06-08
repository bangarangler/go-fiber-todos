#!/bin/zsh

curl -X POST http://localhost:3000/todos \
  -H "Content-Type: application/json" \
  -d '{"name": "new todo from curl"}' \
  -v | jq

