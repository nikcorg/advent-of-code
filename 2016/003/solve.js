const { compose, flatten, map, numSort, reduce, tap } = require('../../utils');

const validTriangle =
  ([a, b, c]) => a + b > c;

const first = compose(
  reduce((n, b) => b ? n + 1 : n, 0),
  map(compose(validTriangle, numSort)));

const partition =
  input =>
    input.length < 9
      ? []
      : input
        .slice(0, 9)
        .reduce((p, x, i) => {
          // this would be so much nicer using a lens, but i'm too lazy
          p[i % 3].push(x);
          return p;
        }, [[], [], []])
        .concat(partition(input.slice(9)));

const second = compose(
  first,
  partition,
  flatten);

module.exports = { first, second };
