#!/bin/bash
ls -1 build | while IFS= read -r file; do
	rm "build/$file"
done
exit