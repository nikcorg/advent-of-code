const test = require('tape');
const { first, second } = require('./solve');

const input = [[5, 10, 25]];

test('day 3', t => {
  t.test('test case 1', t => {
    t.equal(first(input), 0);
    t.end();
  });
});
