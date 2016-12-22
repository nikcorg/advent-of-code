const { compose, takeUntilNotNull } = require('../../utils');

const compassMap = {
  N: { L: 'W', R: 'E' },
  E: { L: 'N', R: 'S' },
  S: { L: 'E', R: 'W' },
  W: { L: 'S', R: 'N' }
};

const parseInput =
  input =>
    input.split(', ');

const distance =
  ([x, y]) =>
    Math.abs(x) + Math.abs(y);

const step =
  ([x, y, h], instr) => {
    const to = instr[0];
    const steps = parseInt(instr.slice(1), 10);
    const newH = compassMap[h][to];

    switch (newH) {
      case 'N':
        return [x, y - steps, newH];
      case 'E':
        return [x + steps, y, newH];
      case 'S':
        return [x, y + steps, newH];
      case 'W':
        return [x - steps, y, newH];
    }
  };

const traverse =
  steps =>
    steps.reduce(step, [0, 0, 'N']);

module.exports.first = compose(distance, traverse, parseInput);

const mapVectors =
  steps =>
    steps.reduce(({ vectors, loc }, instr) => {
      const nextLoc = step(loc, instr);
      vectors.push([loc, nextLoc]);
      return { vectors, loc: nextLoc };
    }, { vectors: [], loc: [0, 0, 'N'] });

const omit =
  (list, exclude) => list.filter(i => i !== exclude);

const intersection =
  ([[afx, afy], [atx, aty]], [[bfx, bfy], [btx, bty]]) => {
    if (Math.min(bfx, btx) < afx && afx < Math.max(bfx, btx)) {
      if (Math.min(afy, aty) < bfy && bfy < Math.max(afy, aty)) {
        return [afx, bfy];
      }
    } else if (Math.min(afx, atx) < bfx && bfx < Math.max(afx, atx)) {
      if (Math.min(bfy, bty) < afy && afy < Math.max(bfy, bty)) {
        return [bfx, afy];
      }
    }

    return null;
  };

module.exports.intersection = intersection;

const findIntersection =
  ({ vectors }) => vectors.length < 2
    ? null
    : takeUntilNotNull(v =>
        takeUntilNotNull(vv => intersection(v, vv), omit(vectors, v)),
        vectors);

const tap = msg => v => console.error(msg) || console.dir(v, { depth: null }) || v;

module.exports.second = compose(distance, findIntersection, mapVectors, parseInput);

