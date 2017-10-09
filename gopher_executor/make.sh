#!/usr/bin/env bash

rm -rf dist

go build -o dist/executor -race src/executor.go
./dist/executor
