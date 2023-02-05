#!/bin/bash

read -p "id: " id


curl PUT http://localhost:8080/products/$id -H 'Content-Type: application/json' -d "{\"name\":\"${user}\",\"manufacturer\":\"${manufacturer}\"}" -v
