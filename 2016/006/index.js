const { chain, compose, filter, map, readFile, split, trim } = require('../../utils');
const { first, second } = require('./solve');

const parseInput = compose(filter(r => r.length > 0), map(trim), split('\n'));

const prog = compose(
  chain(first),
  chain(parseInput),
  readFile);

prog('./input.txt')
  .then(r => console.log('1: %s\n', r))
  .catch(e => console.error(e.stack));
