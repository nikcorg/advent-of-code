const test = require('tape');
const { first } = require('./solve');

test('day 5', t => {
  t.test('first', t => {
    t.equal(first('abc'), '18f47a30');
    t.end();
  });
});
