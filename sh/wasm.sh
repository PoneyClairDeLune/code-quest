#!/bin/bash
echo "Running!"
cd "quest/$1"
wasmtime index.wat
exit
