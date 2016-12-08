#!/bin/sh

GOPATH=${GOPATH:-"${HOME}/go"}

mkdir -p ${GOPATH}/src
cp -fR ${PWD}/src ${GOPATH}
${GO:="go"} install wiselusterlab/container/stack
