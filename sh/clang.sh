#!/bin/bash
mkdir -p ./build/
rm ./build/$1.clang
clang -lm -o "build/$1.clang" "quest/$1/index.c"
cd "quest/$1"
../../build/$1.clang
exit