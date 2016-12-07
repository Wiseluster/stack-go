#!/bin/sh

GOPATH="${PWD}:${GOPATH}" ${GO:="go"} run main.go
