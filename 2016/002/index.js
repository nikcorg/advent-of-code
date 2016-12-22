const { chain, compose, readFile, trim } = require('../../utils');
const { first, second } = require('./solve');

const parseInput =
  raw => raw.split('\n').map(trim);

const getInput = compose(chain(parseInput), readFile);

const input = getInput('./input.txt');

const solveFirst =
  input => first(input);

const solveSecond =
  input => second(input);

const runSolvers =
  input => [solveFirst(input), solveSecond(input)];

const echoSolutions =
  ([first, second]) => console.log('1: %s\n2: %s\n', first, second);

const prog = compose(
  chain(echoSolutions),
  chain(runSolvers),
  getInput
);

prog('./input.txt');

