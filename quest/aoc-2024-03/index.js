"use strict";

// Required preparation steps
let parseCommandSimple = (cmd) => {
	/*
	0: waiting for (
	1: waiting for ,
	2: waiting for )
	*/
	let command = cmd.match(/[A-Za-z0-9_]+\(/)[0];
	let args = cmd.match(/\([0-9,]+\)/)[0];
	args = args.substring(1, args.length - 1).split(",");
	return [
		command.substring(0, command.length - 1),
		...args
	];
};

// Format the input

let inputData = await Deno.readTextFile("./dataInput.txt");
let commandChain = [];
for (let e of inputData.match(/mul\(\d+,\d+\)/g)) {
	commandChain.push(parseCommandSimple(e));
};
console.debug(commandChain);

// Part 1: Sum of all simply-matched multiplication
let sumSimple = 0;
for (let command of commandChain) {
	switch (command[0]) {
		case "mul": {
			sumSimple += parseInt(command[1]) * parseInt(command[2]);
			break;
		};
		default: {
			console.debug(`Unknown command parsed: "${command[0]}"`);
		};
	};
};
console.info(`Part #1: ${sumSimple}`);
