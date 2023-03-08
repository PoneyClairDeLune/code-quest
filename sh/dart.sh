#!/bin/bash
mkdir -p ./build/
printf "Removing old build result... "
rm ./build/$1.dartc
echo "Building program..."
dart compile exe -o "build/$1.dartc" "quest/$1/index.dart"
echo "Running!"
cd "quest/$1"
../../build/$1.dartc
exit
