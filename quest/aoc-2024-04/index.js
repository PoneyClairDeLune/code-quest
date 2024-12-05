"use strict";

// Required preparation steps

import TextReader from "https://jsr.io/@ltgc/rochelle/0.2.2/dist/textRead.mjs";

const debugMode = 4;
let parseLine = async (targetArray, readable) => {
	let lineReader = TextReader.lineRaw(readable);
	for await (let line of lineReader) {
		targetArray.push(line); // Assumes ASCII/UTF-8
	};
};
// The "XMAS" matcher
const xmasBinData = [88, 77, 65, 83];
const preRotMasBin = [
	[77, 83, 65, 77, 83], // L
	[83, 83, 65, 77, 77], // B
	[83, 77, 65, 83, 77], // R
	[77, 77, 65, 83, 83]  // T
]; // Fine, I'll rotate it myself
let matchXmasFull = (bin2d, posX, posY, direction) => {
	/*
	Direction parameter (3-bit):
	L -> R
	T -> B
	TL -> BR
	BL -> TR
	---
	inverted?
	*/
	let sWidth = bin2d[0].length, sHeight = bin2d.length;
	// Error out if the match is attempting to access OOB
	// Start index match
	if (posX < 0 || posY < 0 || posX >= sWidth || posY >= sHeight) {
		return false;
	};
	// End index match
	//console.debug([posX, sWidth - posX - 1][direction & 1] + xmasBinData.length);
	switch (direction >> 1) {
		case 0: {
			if ([posX, sWidth - posX - 1][direction & 1] + xmasBinData.length > sWidth) {
				return false;
			};
			break;
		};
		case 1: {
			if ([posY, sHeight - posY - 1][direction & 1] + xmasBinData.length > sHeight) {
				return false;
			};
			break;
		};
		case 2: {
			if ([posX, sWidth - posX - 1][direction & 1] + xmasBinData.length > sWidth || [posY, sHeight - posY - 1][direction & 1] + xmasBinData.length > sHeight) {
				return false;
			};
			break;
		};
		case 3: {
			//console.debug(direction);
			if ([posX, sWidth - posX - 1][direction & 1] + xmasBinData.length > sWidth || [sHeight - posY - 1, posY][direction & 1] + xmasBinData.length > sHeight) {
				return false;
			};
			break;
		};
	};
	let matchCount = 0;
	for (let i = 0; i < xmasBinData.length; i ++) {
		// Coordinate translation
		let rX = posX, rY = posY;
		switch (direction) {
			case 0:
			case 4:
			case 6: {
				rX += i;
				break;
			};
			case 1:
			case 5:
			case 7: {
				rX -= i;
				break;
			};
		}
		switch (direction) {
			case 2:
			case 4:
			case 7: {
				rY += i;
				break;
			};
			case 3:
			case 5:
			case 6: {
				rY -= i;
				break;
			};
		}
		switch (direction) {
			case 6:
			case 7: {
				//console.debug(`(${rX}, ${rY}) ${direction}: ${bin2d[rY][rX]}`);
				break;
			};
		};
		if (bin2d[rY][rX] == xmasBinData[i]) {
			matchCount ++;
		};
	};
	return matchCount == xmasBinData.length;
};
// The "MMASS" matcher
let matchXmasMas = (bin2d, posX, posY, direction) => {
	/*
	Direction:
	0: L->R
	1: B->T
	2: R->L
	3: T->B
	*/
	let sWidth = bin2d[0].length, sHeight = bin2d.length;
	// Error out if the match is attempting to access OOB
	// Centre index match
	if (posX < 1 || posY < 1 || posX + 2 > sWidth || posY + 2 > sHeight) {
		return false;
	};
	let extractedView = new Uint8Array(5);
	extractedView[0] = bin2d[posY - 1][posX - 1];
	extractedView[1] = bin2d[posY - 1][posX + 1];
	extractedView[2] = bin2d[posY][posX];
	extractedView[3] = bin2d[posY + 1][posX - 1];
	extractedView[4] = bin2d[posY + 1][posX + 1];
	let matchCount = 0;
	for (let i = 0; i < preRotMasBin[direction].length; i ++) {
		if (extractedView[i] == preRotMasBin[direction][i]) {
			matchCount ++;
		};
	};
	if (matchCount == preRotMasBin[direction].length) {
		//console.debug(extractedView);
	};
	return matchCount == preRotMasBin[direction].length;
};

// Test section
/*let testW = 5, testH = 5, testD = 4;
let testData = [];
for (let y = 0; y < testH; y ++) {
	testData.push(new Array(testW).fill(0));
};
for (let y = 0; y < testH; y ++) {
	for (let x = 0; x < testW; x ++) {
		console.debug(`OOB yet at (${x}, ${y}) ${testD}? ${matchXmasMas(testData, x, y, testD)}`);
	};
};*/

// Import and pre-process input
let flatData = [];
await parseLine(flatData, (await Deno.open("./dataInput.txt")).readable);

// Part 1: Count linear XMAS

let firstOccursA = [];
for (let y = 0; y < flatData.length; y ++) {
	for (let x = 0; x < flatData[0].length; x ++) {
		if (flatData[y][x] == 88) { // X
			firstOccursA.push([x, y]);
		};
	};
};
console.debug(`The letter "X" has occurred ${firstOccursA.length} times.`);
let part1Result = 0;
for (let occuredX of firstOccursA) {
	for (let direction = 0; direction < 8; direction ++) {
		let matchResult = matchXmasFull(flatData, ...occuredX, direction);
		//console.debug(`(${occuredX.join(", ")}) ${direction}: ${matchResult}`);
		if (matchResult) {
			part1Result ++;
			//console.debug(`(${occuredX.join(", ")}) ${direction}: ${matchResult}`);
		};
	};
};
console.debug(`Part #1: ${part1Result}`);

// Part 2: Count 2D X-shaped MAS

let firstOccursB = [];
for (let y = 0; y < flatData.length; y ++) {
	for (let x = 0; x < flatData[0].length; x ++) {
		if (flatData[y][x] == 65) { // A
			firstOccursB.push([x, y]);
		};
	};
};
console.debug(`The letter "A" has occurred ${firstOccursB.length} times.`);
let part2Result = 0;
for (let occuredA of firstOccursB) {
	let solved = false;
	for (let direction = 0; direction < 4; direction ++) {
		if (solved) {
			continue;
		};
		let matchResult = matchXmasMas(flatData, ...occuredA, direction);
		//console.debug(`(${occuredA.join(", ")}) ${direction}: ${matchResult}`);
		if (matchResult) {
			part2Result ++;
			//console.debug(`(${occuredA.join(", ")}) ${direction}: ${matchResult}`);
		};
	};
};
console.debug(`Part #2: ${part2Result}`);
