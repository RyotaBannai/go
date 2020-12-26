#!/bin/sh

GIT_VER=`git describe --tags`
go build -ldflags "-X main.version=${GIT_VER}"

# git tag -a v1.4 -m "my version 1.4"
# go build -idflags "-X github.com/XXX/packageName.{variableName}=${GIT_VER}" // set the package variable.