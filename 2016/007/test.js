const test = require('tape');
const { first, second } = require('./solve');

test('day 7', t => {
  t.test('should match', t => {
    [
      'ioxxoj[asdfgh]zxcvbn',
      'abba[mnop]qrst']
        .forEach(x =>
          t.equal(first(x), true));
    t.end();
  });

  t.test('should not match', t => {
    [
      'abcd[bddb]xyyx',
      'aaaa[qwer]tyui']
        .forEach(x =>
          t.equal(first(x), false));
    t.end();
  });
});
