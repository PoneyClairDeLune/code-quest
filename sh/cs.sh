#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.dncs
echo "Building program..."
#RUST_BACKTRACE=1 rustc -o "build/$1.rustc" "quest/$1/index.rs"
echo "Running!"
cd "quest/$1"
#../../build/$1.dncs
exit
