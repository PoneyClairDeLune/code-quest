#!/bin/bash
args=( "$@" )
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.gcc
echo "Building program..."
gcc -lm -o "build/$1.gcc" "quest/$1/index.c" ${args[@]:1}
echo "Running!"
cd "quest/$1"
../../build/$1.gcc
exit
