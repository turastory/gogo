#!/bin/sh
trap 'exit' INT
while true
do
  fswatch -o ./**/*.go | echo "======" && go test -cover ./...
done
