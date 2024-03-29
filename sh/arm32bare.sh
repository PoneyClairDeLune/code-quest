#!/bin/bash
exit
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.as
rm ./build/$1.o
echo "Building program..."
as "quest/$1/index.arm7.asm" -o "build/$1.o"
gcc -nostdlib -o "build/$1.as" "build/$1.o"
echo "Running!"
cd "quest/$1"
../../build/$1.as
exit