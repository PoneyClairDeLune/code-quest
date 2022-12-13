package main;

import (
	"fmt"
	"bufio"
	"os"
);

type Blob struct {
	isDir bool
	id uint64
	size uint64
	name string
	parent uint64
};

var fakeFs []Blob;

func (x Blob) path() string {
	workBlob := x;
	path := "/" + workBlob.name;
	for (workBlob.parent > 0) {
		workBlob = fakeFs[workBlob.parent];
		path = "/" + workBlob.name + path;
	};
	return path;
};

func (x Blob) du() uint64 {
	var fSize uint64 = 0;
	if (x.isDir) {
		for _, e0 := range fakeFs {
			if (e0.parent == x.id) {
				fSize += e0.size;
				if (e0.isDir && e0.id != x.id) {
					fSize += e0.du();
				};
			};
		};
	};
	return fSize;
};

func main() {
	fmt.Printf("Very Awesome Shell Extended v1.14.514, built on 10 Aug 1919\nCopyright Macrohard Shimokitazawa Co. (C) 1919, no rights reserved.\n");
	Stdin := bufio.NewScanner(os.Stdin);
	var command string = " ";
	var cmdSep [4]string;
	var fsFull uint64 = 70000000;
	var fsUpdate uint64 = 30000000;
	fakeFs = append(fakeFs, Blob{true, 0, 0, "", 0});
	curFolder := fakeFs[0];
	for (len(command) > 0) {
		command = "";
		fmt.Printf("[%4d](%4d) %s # ", len(fakeFs), curFolder.id, curFolder.path());
		Stdin.Scan();
		command = Stdin.Text();
		cmdSep[0] = "";
		cmdSep[1] = "";
		cmdSep[2] = "";
		cmdSep[3] = "";
		fmt.Sscanf(command, "%s %s %s %s", &cmdSep[0], &cmdSep[1], &cmdSep[2], &cmdSep[3]);
		switch (cmdSep[0]) {
			case "aocp1": {
				// Challenge part 1
				var sizeSum uint64 = 0;
				for _, e0 := range fakeFs {
					if (e0.isDir) {
						sizeDir := e0.du();
						//fmt.Printf("Directory (%4d) %s occupies %d bytes\n", e0.id, e0.path(), sizeDir);
						if (sizeDir <= 100000) {
							sizeSum += sizeDir;
						};
					};
				};
				fmt.Printf("AoC result: %d B\n", sizeSum);
				break;
			};
			case "aocp2": {
				// Challenge part 2
				minSize := fakeFs[0].du();
				minDir := uint64(0);
				criteria := fsUpdate - (fsFull - minSize);
				fmt.Printf("Further %d bytes required for the update.\n", criteria);
				for _, e0 := range fakeFs {
					if (e0.isDir) {
						sizeDir := e0.du();
						if (sizeDir >= criteria) {
							//fmt.Printf("Directory (%4d) %s occupies %d bytes\n", e0.id, e0.path(), sizeDir);
							if (minSize > sizeDir) {
								minSize = sizeDir;
								minDir = e0.id;
							};
						};
					};
				};
				fmt.Printf("Directory %s with size %d matches the criteria.\n", fakeFs[minDir].path(), minSize);
				break;
			};
			case "$": {
				// Command record mode
				if (cmdSep[1] == "cd") {
					if (cmdSep[2] == "..") {
						curFolder = fakeFs[curFolder.parent];
					} else if (cmdSep[2] == "/") {
						curFolder = fakeFs[0];
					} else {
						var targetDir uint64 = 0;
						for _, e0 := range fakeFs {
							if (targetDir < 1 && e0.parent == curFolder.id && e0.name == cmdSep[2]) {
								targetDir = e0.id;
							};
						};
						if (targetDir > 0) {
							curFolder = fakeFs[targetDir];
						} else {
							fmt.Printf("Directory %s not found.\n", cmdSep[2]);
						};
					};
				};
				break;
			};
			case "cd": {
				// Change directory
				if (cmdSep[1] == "..") {
					curFolder = fakeFs[curFolder.parent];
				} else if (cmdSep[1] == "/") {
					curFolder = fakeFs[0];
				} else {
					var targetDir uint64 = 0;
					for _, e0 := range fakeFs {
						if (targetDir < 1 && e0.parent == curFolder.id && e0.name == cmdSep[1]) {
							targetDir = e0.id;
						};
					};
					if (targetDir > 0) {
						curFolder = fakeFs[targetDir];
					} else {
						fmt.Printf("Directory %s not found.\n", cmdSep[1]);
					};
				};
				break;
			};
			case "dir": {
				// Create new directory
				if (len(cmdSep[1]) > 0) {
					newId := uint64(len(fakeFs));
					fakeFs = append(fakeFs, Blob{true, newId, 0, cmdSep[1], curFolder.id});
				} else {
					fmt.Printf("Name cannot be blank.\n");
				};
				break;
			};
			case "du": {
				// Directory usage
				fmt.Printf("Size      ID    File name\n");
				for i0, e0 := range fakeFs {
					if (i0 > 0 && e0.parent == curFolder.id) {
						var fSize uint64 = e0.size;
						if (e0.isDir) {
							fSize += e0.du();
						};
						fmt.Printf("%8d (%4d) %s\n", fSize, i0, e0.name);
					};
				};
				break;
			};
			case "ls": {
				// List files
				var fileCount uint64 = 0;
				fmt.Printf("Size      ID    File name\n");
				for i0, e0 := range fakeFs {
					if (i0 > 0 && e0.parent == curFolder.id) {
						if (fakeFs[i0].isDir) {
							fmt.Printf("DIR      (%4d) %s\n", i0, e0.name);
						} else {
							fmt.Printf("%8d (%4d) %s\n", e0.size, i0, e0.name);
						};
						fileCount ++;
					};
				};
				fmt.Printf("Contains %d file(s).\n", fileCount);
				break;
			};
			case "tree": {
				// List tree view
				fmt.Printf("Not implemented.\n");
				break;
			};
			case "help": {
				fmt.Printf("$        Record commands\naocp1    Finish the AoC 2022 Day 7 Part 1 challenge\naocp2    Finish the Part 2 challenge\ncd       Change directory\ndir      Create new directory\ndu       Calculate file sizes\nls       List files\ntree     View the current directory as a tree\n<number> Allocate file space\n");
				break;
			};
			default: {
				if (cmdSep[0][0] > 47 && cmdSep[0][0] < 58) {
					var fSize uint64 = 0;
					newId := uint64(len(fakeFs));
					fmt.Sscanf(cmdSep[0], "%d", &fSize);
					fakeFs = append(fakeFs, Blob{false, newId, fSize, cmdSep[1], curFolder.id});
				} else {
					fmt.Printf("Unknown command: %s %s %s %s\n", cmdSep[0], cmdSep[1], cmdSep[2], cmdSep[3]);
				};
			};
		};
	};
};
