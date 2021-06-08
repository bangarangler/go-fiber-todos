#!/bin/zsh

curl -X PATCH http://localhost:3000/todos/2 \
  -H "Content-Type: application/json" \
  -d '{"completed": true}' \
  -v | jq


# -d '{"name": "updated todo from curl"}' \
