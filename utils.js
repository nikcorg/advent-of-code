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

module.exports = {
  chain,
  compose,
  readFile,
  takeUntilNotNull,
  tap,
  trim
};
