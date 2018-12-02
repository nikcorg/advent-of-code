const { chain, compose, empty, map, not, readFile, match, tap, trim } = require("../../utils");

const { first, second } = require("./solve");

const parseInput = raw =>
  raw
    .split("\n")
    .map(trim)
    .filter(compose(not, empty))
    .map(compose(map(Number), match(/[0-9]+/g)));

const echo = ([a, b]) => console.log("1: %s\n2: %s\n", a, b);

const getInput = compose(chain(parseInput), readFile);

const solve = input => [first(input), second(input)];

const prog = compose(chain(echo), chain(solve), getInput);

prog("./input.txt").catch(e => console.error(e));
