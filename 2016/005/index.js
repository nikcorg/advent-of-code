const { compose, chain, head, readFile, split, trim } = require("../../utils");
const { first, second } = require("./solve");

const parseInput = compose(trim, head, split("\n"));
const solve = i => [first(i), second(i)];
const prog = compose(chain(solve), chain(parseInput), readFile);

prog("./input.txt").then(([first, second]) => console.log("1: %s\n2: %s\n", first, second));
