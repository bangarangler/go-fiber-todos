#!/bin/zsh

curl -X POST http://localhost:3000/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"name": "walk the fish"}' \
  -v | jq

