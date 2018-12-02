const { chain, compose, split, readFile } = require("../../utils");
const { first, second } = require("./solve");

const getInput = compose(chain(split("\n")), readFile);
const solve = i => [first(i), second(i)];
const echo = ([a, b]) => {
  console.log("1: %s\n2: \n", a, b);
};

const prog = compose(chain(echo), chain(solve), getInput);

prog("./input.txt");
