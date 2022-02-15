#!/bin/bash

## THIS will check if everything hidden is there


## MYSQL
DIR="/etc/mysql"
if [ -d "$DIR" ]; then
    mkdir "$DIR"
fi

FILE="$DIR/my.cnf"
if [ -f "$FILE" ]; then
    touch "$FILE"
fi
