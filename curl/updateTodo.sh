#!/bin/zsh

curl -X PATCH http://localhost:3000/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "updated walk the dog!"}' \
  -v | jq


# -d '{"name": "updated todo from curl"}' \
