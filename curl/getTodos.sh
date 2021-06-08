#!/bin/zsh

curl http://localhost:3000/v1/todos \
  -v | jq
