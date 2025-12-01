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
  if (direction === "R") {
    currentPos += parseInt(count);
    if (currentPos > max) {
      currentPos = currentPos % totalPositions;
    }
  } else {
    currentPos -= parseInt(count);
    if (currentPos < 0) {
      currentPos = currentPos % totalPositions;
    }
  }
  return currentPos;
}

const lines = input.trim().split(/\r?\n/);

lines.forEach(line => {
  pos = rotate(line, pos);
  console.log(pos);
  if (pos === 0) {
    zeroCount++;
  }
});

console.log("zeroCount:", zeroCount);
