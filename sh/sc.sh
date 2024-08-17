#!/bin/bash
mkdir -p ./build/
echo "Building and running!"
cd "quest/$1"
scala -howtorun:script -explain index.sc
exit
