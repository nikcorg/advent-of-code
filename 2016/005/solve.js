// A hash indicates the next character in the password if its hexadecimal representation starts with five zeroes. If it does, the sixth character in the hash is the next character of the password.
// The first index which produces a hash that starts with five zeroes is 3231929, which we find by hashing abc3231929; the sixth character of the hash, and thus the first character of the password, is 1.

const crypto = require('crypto');

function* makeHashIterator(stem) {
  let suffix = 0;

  while (true) {
    const hash = crypto.createHash('md5');
    hash.update(`${stem}${suffix}`);
    const digest = hash.digest('hex');

    if (digest.substr(0, 5) === '00000') {
      yield digest;
    }

    suffix++;
  }
}

const first = i => {
  const passcode = [];
  const iter = makeHashIterator(i);

  while (passcode.length < 8) {
    passcode.push(iter.next().value.substr(5, 1));
  }

  return passcode.join('');
};

const second = () => null;
// 1. take zip(index, hash) until 8
// 2. sort by index
// 3. map over snd

module.exports = {
  first,
  second,
  makeHashIterator
};
