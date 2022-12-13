#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.zigc
rm ./build/$1.zigc.o
echo "Building program..."
zig build-exe -femit-bin="build/$1.zigc" "quest/$1/index.zig"
echo "Running!"
cd "quest/$1"
../../build/$1.zigc
exit
