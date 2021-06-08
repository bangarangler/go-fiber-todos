#!/bin/zsh

curl -X DELETE http://localhost:3000/todos/2 \
  -v | jq
