#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.rustc
echo "Building program..."
RUST_BACKTRACE=1 rustc --target i686-pc-windows-gnu -C target-feature=+crt-static -C link-self-contained=yes -o "build/$1.rsxp.exe" "quest/$1/index.rs"
echo "Running!"
cd "quest/$1"
wine ../../build/$1.rsxp.exe
exit
