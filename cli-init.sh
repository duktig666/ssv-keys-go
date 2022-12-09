#!/bin/bash

echo "exec 'go mod tidy -compat=1.18'"
go mod tidy -compat=1.18 || exit

echo "exec 'go build -o ssv-key'"
go build -o ssv-key || exit
