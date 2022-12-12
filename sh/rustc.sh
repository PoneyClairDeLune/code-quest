#!/bin/bash
mkdir -p ./build/
rm ./build/$1.rustc
rustc -o "build/$1.rustc" "quest/$1/index.rs"
cd "quest/$1"
../../build/$1.rustc
exit