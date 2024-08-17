#!/bin/bash
mkdir -p ./build/
echo "Building and running!"
cd "quest/$1"
scala -explain index.sc
exit
