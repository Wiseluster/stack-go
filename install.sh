#!/bin/sh

mkdir -p ${GOPATH:="${HOME}/go"}/src
cp -fR ${PWD}/src ${GOPATH}/
${GO:="go"} install wiselusterlab/container/stack
