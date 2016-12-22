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

const match =
  x =>
    s => s.match(x);

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

const empty =
  x => x.length < 1;

const sort =
  f =>
    xs => xs.slice(0).sort(f); // slice, to avoid mutation

const numSort = sort((a, b) => a - b);

const flatten =
  xs => xs.reduce((f, x) => f.concat(Array.isArray(x) ? x : [x]), []);

const ifNotNull =
  f =>
    x => x == null ? x : f(x);

const filter =
  f =>
    xs => xs.filter(f);

const slice =
  (a, b) =>
    xs => xs.slice(a, b);

const replace =
  (s, r) =>
    x => x.replace(s, r);

module.exports = {
  chain,
  compose,
  constrain,
  empty,
  equals,
  filter,
  flatten,
  ifNotNull,
  join,
  lookup,
  match,
  map,
  numSort,
  not,
  readFile,
  reduce,
  replace,
  slice,
  sort,
  split,
  takeUntilNotNull,
  tap,
  trim
};
