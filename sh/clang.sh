#!/bin/bash
mkdir -p ./build/
rm ./build/$1.clang
clang -lm -o "build/$1.clang" "quest/$1/index.c"
./build/$1.gcc
exit