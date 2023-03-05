#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.as
rm ./build/$1.o
echo "Building program..."
arm-none-eabi-as "quest/$1/index.arm7.asm" -o "build/$1.o"
arm-none-eabi-gcc -nostdlib -o "build/$1.as" "build/$1.o"
echo "Running!"
cd "quest/$1"
qemu-arm ../../build/$1.as
exit