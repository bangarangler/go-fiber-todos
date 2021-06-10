#!/bin/zsh

curl -X DELETE http://localhost:3000/v1/todos/7 \
  -v | jq
