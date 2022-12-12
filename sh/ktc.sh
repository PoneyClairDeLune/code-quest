#!/bin/bash
mkdir -p ./build/
rm ./build/$1.kexe
kotlinc-native -o "build/$1" "quest/$1/index.kt"
cd "quest/$1"
../../build/$1.kexe
exit