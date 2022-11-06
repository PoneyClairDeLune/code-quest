#!/bin/bash
mkdir -p ./build/
rm ./build/$1.gcc
gcc -lm -o "build/$1.gcc" "quest/$1/index.c"
./build/$1.gcc
exit