#!/bin/sh

cp -fR ${PWD}/src ${GOPATH:="${HOME}/go"}
${GO:="go"} install wiselusterlab/container/stack
