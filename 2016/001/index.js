const { readFile } = require('../../utils');
const { first, second } = require('./solve');

const input = readFile('./input.txt');

input
  .then(first)
  .then(solution => console.log('1', solution))
  .catch(e => console.error('1', e.stack));

input
  .then(second)
  .then(solution => console.log('2', solution))
  .catch(e => console.error('2', e.stack));
