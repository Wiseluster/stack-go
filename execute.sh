#!/bin/sh

GOPATH="${PWD}:${GOPATH:-"${HOME}/go"}" ${GO:-"go"} run main.go
