#!/bin/bash
mkdir -p ./build/
rm ./build/$1.zigc
rm ./build/$1.zigc.o
zig build-exe -femit-bin="build/$1.zigc" "quest/$1/index.zig"
cd "quest/$1"
../../build/$1.zigc
exit