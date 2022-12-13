#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.goc && echo "done."
echo "Building program..."
go build -o "build/$1.goc" "quest/$1/index.go"
echo "Running!"
cd "quest/$1"
../../build/$1.goc
exit