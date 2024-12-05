"use strict";

// Required preparation steps

let parseCommand = (cmd) => {
	/*
	0: waiting for (
	1: waiting for ,
	2: waiting for )
	*/
	let command = cmd.match(/[A-Za-z0-9_']+\(/)[0];
	//console.debug(command);
	let args = (cmd.match(/\([0-9,]+\)/) ?? [])[0] || [];
	if (args?.length) {
		args = args.substring(1, args.length - 1).split(",");
	};
	return [
		command.substring(0, command.length - 1),
		...args
	];
};

// Format the input

let inputData = await Deno.readTextFile("./dataInput.txt");
let commandChain = [];
for (let e of inputData.match(/(mul|do(|n't))\((\d+,\d+|)\)/g)) {
	let command = parseCommand(e);
	commandChain.push(command);
	console.debug(command);
};

// Part 1: Sum of all simply-matched multiplication

let sumSimple = 0;
for (let command of commandChain) {
	switch (command[0]) {
		case "mul": {
			sumSimple += parseInt(command[1]) * parseInt(command[2]);
			break;
		};
		case "do":
		case "don't": {
			// Ignored commands
			break;
		};
		default: {
			console.debug(`Unknown command parsed: "${command[0]}"`);
		};
	};
};
console.info(`Part #1: ${sumSimple}`);

// Part 2: Sum of all valid multiplications

let sumToggled = 0, isMultiplying = true;
for (let command of commandChain) {
	switch (command[0]) {
		case "don't": {
			isMultiplying = false;
			break;
		};
		case "do": {
			isMultiplying = true;
			break;
		};
		case "mul": {
			if (isMultiplying) {
				sumToggled += parseInt(command[1]) * parseInt(command[2]);
			};
			break;
		};
		default: {
			console.debug(`Unknown command parsed: "${command[0]}"`);
		};
	};
};
console.info(`Part #2: ${sumToggled}`);
