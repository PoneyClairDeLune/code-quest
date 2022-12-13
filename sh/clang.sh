#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.clang
echo "Building program..."
clang -lm -o "build/$1.clang" "quest/$1/index.c"
echo "Running!"
cd "quest/$1"
../../build/$1.clang
exit
