const fs = require('fs');

const compose =
  (...fns) => {
    const [first, ...rest] = fns.slice(0).reverse();

    return (...args) =>
      rest.reduce(
        (p, f) => f(p),
        first(...args)
      );
  };

const readFile =
  file =>
    new Promise((res, rej) =>
      fs.readFile(file,
        (err, data) =>
          (err
            ? rej(err)
            : res(String(data)))));

const takeUntilNotNull =
  (f, [x, ...xs]) =>
    x == null
      ? null
      : (f(x) || takeUntilNotNull(f, xs));

const trim =
  str => str.trim();

const chain =
  fn => p => p.then(fn);

const tap =
  m =>
    x => console.log(m) || console.dir(x, { depth: null }) || x;

const split =
  x =>
    s => s.split(x);

const join =
  g =>
    s => s.join(g);

const map =
  f =>
    x => x.map(f);

const reduce =
  (f, i) =>
    x => x.reduce(f, i);

const equals =
  a =>
    b => a === b;

const not =
  v => !v;

const lookup =
  from =>
    what => from[what];

const constrain =
  (min, max) =>
    n => Math.min(max, Math.max(n, min));

module.exports = {
  chain,
  compose,
  constrain,
  equals,
  join,
  lookup,
  map,
  not,
  readFile,
  reduce,
  split,
  takeUntilNotNull,
  tap,
  trim
};
