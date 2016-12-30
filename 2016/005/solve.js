// A hash indicates the next character in the password if its hexadecimal representation starts with five zeroes. If it does, the sixth character in the hash is the next character of the password.
// The first index which produces a hash that starts with five zeroes is 3231929, which we find by hashing abc3231929; the sixth character of the hash, and thus the first character of the password, is 1.

const crypto = require('crypto');

const first = i => {
  const passcode = [];
  let suffix = 0;

  while (passcode.length < 8) {
    const hash = crypto.createHash('md5');
    hash.update(`${i}${suffix}`);
    const digest = hash.digest('hex');

    if (digest.substr(0, 5) === '00000') {
      passcode.push(digest.substr(5, 1));
    }

    suffix++;
  }

  return passcode.join('');
};

const second = () => null;

module.exports = {
  first,
  second
};
