#!/bin/bash

curl POST http://localhost:8080/products -H 'Content-Type: application/json' -d "{\"name\":\"${user}\",\"manufacturer\":\"${manufacturer}\"}" -v
