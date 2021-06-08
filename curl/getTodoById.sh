#!/bin/zsh

curl http://localhost:3000/todos/1 \
  -v | jq
