#!/bin/sh
trap 'exit' INT
while true
do
  fswatch -o ./**/*.go | go test -cover ./...
done
