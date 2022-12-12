#!/bin/bash
mkdir -p ./build/
rm ./build/$1.goc
go build -o "build/$1.goc" "quest/$1/index.go"
cd "quest/$1"
../../build/$1.goc
exit