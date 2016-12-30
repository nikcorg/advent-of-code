const { compose, chain, head, readFile, trim } = require('../../utils');
const { first, second } = require('./solve');

const prog = compose(chain(compose(first, trim, head)), readFile);

prog('./input.txt').then(r => console.log('1: %s\n', r));
