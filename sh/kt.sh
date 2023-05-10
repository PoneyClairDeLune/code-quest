#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.kexe
echo "Building program..."
kotlinc -o "build/$1" "quest/$1/index.kt"
echo "Running!"
cd "quest/$1"
../../build/$1.kexe
exit
