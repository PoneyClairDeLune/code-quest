"use strict";

// Required preparation steps

import TextReader from "https://jsr.io/@ltgc/rochelle/0.2.2/dist/textRead.mjs";

const debugMode = true;
let parseLine = async (targetArray, readable) => {
	let lineReader = TextReader.line(readable);
	for await (let line of lineReader) {
		let elements = line.split(" ");
		for (let i = 0; i < elements.length; i ++) {
			// I know there are ES6+ alternatives, yet this is still the fastest
			elements[i] = parseInt(elements[i]);
		};
		targetArray.push(elements);
	};
};
let routes = [];

// Format the input data
await parseLine(routes, (await Deno.open("./dataInput.txt")).readable);

// Part 1: Is it safe or unsafe?
let safeCount = 0;
for (let route of routes) {
	if (route[1] == route[0]) {
		debugMode && console.debug(`Route deemed unsafe: Flat at start.`);
		continue;
	};
	let targetDirection = route[1] - route[0];
	let unsafe = false;
	for (let i = 1; i < route.length; i ++) {
		let difference = route[i] - route[i - 1];
		if (difference == 0) {
			unsafe = true;
			debugMode && console.debug(`Route deemed unsafe: Flat en route.`);
			continue;
		};
		if (Math.abs(difference) > 3) {
			unsafe = true;
			debugMode && console.debug(`Route deemed unsafe: Steep en route.`);
			continue;
		};
		if (difference * targetDirection < 0) {
			unsafe = true;
			debugMode && console.debug(`Route deemed unsafe: Altered direction en route.`);
			continue;
		};
	};
	if (unsafe) {
		//debugMode && console.debug(`Desinated route: ${route.join(", ")}`);
		continue;
	};
	safeCount ++;
	debugMode && console.debug(`Route deemed safe.`);
	debugMode && console.debug(`Desinated route: ${route.join(", ")}`);
};
console.info(`Part #1: ${safeCount}`);
