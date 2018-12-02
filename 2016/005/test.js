const test = require('tape');
const { first, second, makeHashIterator } = require('./solve');

test('day 5', t => {
  t.test('makeHashIterator', t => {
    const iter = makeHashIterator('abc')

    t.equal(iter.next().value, '00000155f8105dff7f56ee10fa9b9abd');
    t.equal(iter.next().value, '000008f82c5b3924a1ecbebf60344e00');
    t.end();
  });

  t.test('first', t => {
    t.equal(first('abc'), '18f47a30');
    t.end();
  });

  t.test('second', t => {
    t.equal(second('abc'), '05ace8e3');
    t.end();
  });
});
