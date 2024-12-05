#!/bin/bash
echo "Running!"
cd "quest/$1"
deno run -A index.js
exit
