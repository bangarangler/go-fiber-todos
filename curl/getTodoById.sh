#!/bin/zsh

curl http://localhost:3000/v1/todos/8 \
  -v | jq
