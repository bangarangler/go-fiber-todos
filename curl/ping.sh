#!/bin/zsh

curl http://localhost:3000 \
  # -d "{\"email\": \"jonathan.palacio@nowigence.com\",\"password\":
  # \"Dan3ce18\"}" -H "Content-Type: application/json" \
  #   | jq -r '.data.accessToken')

# rm /Users/jonathanpalacio/Desktop/PLUARIS-API/token.txt
# echo -H "Authorization: Bearer $TOKEN" >> token.txt
# echo $TOKEN >> token.txt




