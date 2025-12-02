import { readFileSync } from "node:fs";

let inputFile = "day_1_input.txt";
// let inputFile = "day_1_sample.txt";
let input = readFileSync(inputFile, "utf8");

let max = 99;
let totalPositions = max + 1; // account for 0 position
let pos = 50;

let zeroCount = 0;

function rotate(instruction: string, currentPos: number): number {
  let direction = instruction[0];
  let count = instruction.slice(1);
  let startingPos = currentPos;
  if (direction === "R") {
    currentPos += parseInt(count);
    zeroCount = zeroCount + (Math.floor(currentPos / totalPositions) - Math.floor(startingPos / totalPositions));
    currentPos = currentPos % totalPositions;
  } else {
    currentPos -= parseInt(count);
    zeroCount = zeroCount + (Math.floor((startingPos - 1) / totalPositions) - Math.floor((currentPos - 1) / totalPositions));
    
    currentPos = currentPos % totalPositions;
    if (currentPos < 0) {
      currentPos += totalPositions;
    }
  }
  console.log("pos: ", currentPos);
  return currentPos;
}

const lines = input.trim().split(/\r?\n/);

console.log("pos: ", pos);
lines.forEach(line => {
  pos = rotate(line, pos);
  // console.log("pos: ", pos);
  // if (pos === 0) {
  //   zeroCount++;
  // }
});

console.log("");
console.log("zeroCount:", zeroCount);