const test = require('tape');
const { decrypt, letterFreq, parseIds, isValid, first, second, checksum } = require('./solve');

test('day 4', t => {
  const input = [
    'aaaaa-bbb-z-y-x-123[abxyz]',
    'totally-real-room-200[decoy]',
    'not-a-real-room-404[oarel]',
    'a-b-c-d-e-f-g-h-987[abcde]'];

  t.test('parse ids', t => {
    t.deepEqual(parseIds(input), [
      ['aaaaa-bbb-z-y-x', 123, 'abxyz'],
      ['totally-real-room', 200, 'decoy'],
      ['not-a-real-room', 404, 'oarel'],
      ['a-b-c-d-e-f-g-h', 987, 'abcde'],
    ]);
    t.end();
  });

  t.test('letter frequency', t => {
    t.deepEqual(letterFreq('aabbcc'), { a: 2, b: 2, c: 2 });
    t.deepEqual(letterFreq('aaaaabbbzyx'), { a: 5, b: 3, z: 1, y: 1, x: 1 });
    t.end();
  });

  t.test('top letters', t => {
    t.deepEqual(checksum('aaaaabbbzyx'), 'abxyz');
    t.deepEqual(checksum('abcdefgh'), 'abcde');
    t.end();
  });

  t.test('test case 1', t => {
    const input = ['aaaaabbbzyx', 123, 'abxyz'];
    const outcome = true;

    t.equal(isValid(input), true);
    t.end();
  });

  t.test('test case 2', t => {
    const input = ['abcdefgh', 987, 'abcde'];
    const outcome = true;

    t.equal(isValid(input), true);
    t.end();
  });

  t.test('test case 3', t => {
    const input = ['notarealroom', 404, 'oarel'];
    const outcome = true;

    t.equal(isValid(input), outcome);
    t.end();
  });

  t.test('test case 4', t => {
    const input = ['totallyrealroom', 200, 'decoy'];
    const outcome = false;

    t.equal(isValid(input), outcome);
    t.end();
  });

  t.test('first', t => {
    const outcome = 1514;

    t.equal(first(input), outcome);
    t.end();
  });

  t.test('test case 5', t => {
    const input = ['qzmt-zixmtkozy-ivhz', 343];
    const outcome = 'very encrypted name';

    t.equal(decrypt(input), outcome);
    t.end();
  });
});
