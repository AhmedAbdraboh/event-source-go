#!/usr/bin/env bash

go get -v github.com/codegangsta/gin

go get -v ./...
gin -i -p 3333 run main.go