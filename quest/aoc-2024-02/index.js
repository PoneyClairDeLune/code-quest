"use strict";

// Required preparation steps

import TextReader from "https://jsr.io/@ltgc/rochelle/0.2.2/dist/textRead.mjs";

const debugMode = 4;
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
// Part 2: Is it safe when up to 1 node gets removed?
let getUnsafeCount = (route) => {
	let targetDirection = route[1] - route[0];
	let unsafeIndexes = [];
	for (let i = 1; i < route.length; i ++) {
		let difference = route[i] - route[i - 1];
		if (difference == 0) {
			unsafeIndexes.push(i - 1);
			debugMode > 3 && console.debug(`Route deemed unsafe: Flat en route.`);
		} else if (Math.abs(difference) > 3) {
			unsafeIndexes.push(i - 1);
			debugMode > 3 && console.debug(`Route deemed unsafe: Steep en route.`);
		} else if (difference * targetDirection < 0) {
			unsafeIndexes.push(i - 1);
			debugMode > 3 && console.debug(`Route deemed unsafe: Altered direction en route.`);
		};
		targetDirection = difference;
	};
	return unsafeIndexes;
};

let safeCount = 0;
let safeDampenedCount = 0;
for (let route of routes) {
	let unsafeIndexes = getUnsafeCount(route);
	switch (unsafeIndexes.length) {
		case 0: {
			safeCount ++;
			safeDampenedCount ++;
			//debugMode > 1 && console.debug(`Desinated safe route (0): ${route.join(", ")}`);
			break;
		};
		default: {
			let solved = false, triedIndexes = new Set();
			for (let i = 0; i < unsafeIndexes.length; i ++) {
				if (solved) {
					continue;
				};
				let realUnsafeIndex = unsafeIndexes[i];
				for (let i0 = 0; i0 < 2; i0 ++) {
					let operatedIndex = realUnsafeIndex + i0;
					if (solved) {
						continue;
					};
					if (triedIndexes.has(operatedIndex)) {
						continue;
					};
					triedIndexes.add(operatedIndex);
					let copiedRoute = route.slice();
					copiedRoute.splice(operatedIndex, 1);
					debugMode > 2 && console.debug(`Submitted route for testing again on index ${operatedIndex}.`);
					let unsafeDampenedIndexes = getUnsafeCount(copiedRoute);
					if (!unsafeDampenedIndexes.length) {
						solved = true;
						safeDampenedCount ++;
						debugMode > 2 && console.debug(`Route proven to be safe on attempt #${triedIndexes.size} (array index #${operatedIndex})`);
						//debugMode > 0 && console.debug(`Desinated almost-safe route (${unsafeIndexes.length}): ${route.join(", ")}`);
					};
				};
			};
			if (!solved) {
				debugMode > 0 && console.debug(`Desinated irrepairable route (${unsafeIndexes.length}): ${route.join(", ")}`);
			};
		};
	};
};
console.info(`Part #1: ${safeCount}`);
console.info(`Part #2: ${safeDampenedCount}`);
