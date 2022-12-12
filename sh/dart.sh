#!/bin/bash
mkdir -p ./build/
rm ./build/$1.dart
#dartPkg="common/dart.json"
#if [ -e "quest/$1/dart.json" ] ; then
#	dartPkg="quest/$1/dart.json"
#	echo "Using project-specific Dart packages."
#else
#	echo "Using common Dart packages."
#fi
#dart compile exe -p "$dartPkg" -o "build/$1.dartc" "quest/$1/index.dart"
dart compile exe -o "build/$1.dartc" "quest/$1/index.dart"
cd "quest/$1"
../../build/$1.dartc
exit