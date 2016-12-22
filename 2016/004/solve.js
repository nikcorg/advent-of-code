// dirty, dirty code. too many puzzles for one day.
const {
  compose,
  equals,
  filter,
  ifNotNull,
  join,
  map,
  match,
  not,
  reduce,
  replace,
  slice,
  tap
} = require('../../utils');

const RE_ROOM_ID = /([a-z\-]+)-([0-9]+)\[([a-z]+)\]/;

const splitId = compose(ifNotNull(slice(1)), match(RE_ROOM_ID))

const parseIds = compose(
    map(([key, id, csum]) => [key, Number(id), csum]),
    filter(compose(not, equals(null))),
    map(splitId));

const letterFreq =
  s => s
    .split('')
    .reduce((f, c) => Object.assign(f, { [c]: (f[c] || 0) + 1 }), {});

const checksum = compose(
  join(''),
  slice(0, 5),
  f => {
    return Object
      .keys(f)
      .sort((a, b) => {
        const n = f[b] - f[a];
        if (n === 0) {
          return b > a ? -1 : 1;
        }
        return n;
      });
  },
  letterFreq,
  replace(/-/g, ''));

const verifyChecksum =
  (name, csum) => csum === checksum(name);

const isValid =
  ([name, _, csum]) => verifyChecksum(name, csum);

const get =
  k =>
    x => x[k];

const add =
  (a, b) => a + b;

const rotate = n =>
  compose(
    cc => Number.isInteger(cc) ? String.fromCharCode(cc) : cc,
    c => {
      if (c === ' ') {
        return c;
      }
      const cc = c.charCodeAt(0) - 97 + (n % 26);
      if (cc > 25) {
        return cc - 26 + 97;
      }
      return cc + 97;
    });

const decrypt =
  ([name, id]) =>
    ({ id,
      name: name
        .replace(/-/g, ' ')
        .split('')
        .map(rotate(id))
        .join('') });

const first = compose(
  reduce(add, 0),
  map(get(1)),
  filter(isValid),
  parseIds);

const second = compose(
  filter(({ name }) => /northpole/.test(name)),
  map(decrypt),
  filter(isValid),
  parseIds);

module.exports = { decrypt, first, second, isValid, parseIds, letterFreq, checksum };
