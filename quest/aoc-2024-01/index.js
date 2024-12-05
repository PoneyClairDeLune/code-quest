"use strict";

// Required preparation steps

import TextReader from "https://jsr.io/@ltgc/rochelle/0.2.2/dist/textRead.mjs";

let parseLine = async (targetArray, readable) => {
	let lineReader = TextReader.line(readable);
	for await (let line of lineReader) {
		let elements = line.split("   ");
		targetArray[0].push(elements[0]);
		targetArray[1].push(elements[1]);
	};
};
let targetArray = [[], []];
let sorter = (a, b) => {
	return a - b;
};

// Format the input data

await parseLine(targetArray, (await Deno.open("./dataInput.txt")).readable);
targetArray[0].sort(sorter);
targetArray[1].sort(sorter);

// Part 1: Sum the distance

let totalDistance = 0;
for (let i = 0; i < targetArray[0].length; i ++) {
	let countedDistance = Math.abs(targetArray[0][i] - targetArray[1][i]);
	//console.debug(countedDistance);
	totalDistance += countedDistance;
};
//console.debug(targetArray);
console.debug(`Part #1: ${totalDistance}`);

// Part 2: Sum the difference

// Count how many times each element has appeared and store them
let countMap = new Map();
for (let e of targetArray[1]) {
	let appearCount = countMap.get(e) ?? 0;
	countMap.set(e, ++ appearCount);
	//console.debug(`Element ${e} has appeared ${appearCount} time(s).`);
};
// Go through the cached appearance map
let totalDifference = 0;
for (let e of targetArray[0]) {
	if (countMap.has(e)) {
		totalDifference += e * countMap.get(e);
	};
};
console.debug(`Part #2: ${totalDifference}`);
