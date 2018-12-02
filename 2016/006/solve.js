// which character is most frequent for each position
const freq = (init = {}, w) =>
  w.split("").reduce(
    (wc, c, i) =>
      Object.assign(wc, {
        [i]: Object.assign(wc[i] || {}, { [c]: ((wc[i] || {})[c] || 0) + 1 }),
      }),
    init,
  );

const first = i => {
  const f = i.reduce((tc, word) => Object.assign(tc, freq(tc, word)), {});

  return Object.keys(f)
    .map(i =>
      Object.keys(f[i])
        .sort((a, b) => f[i][b] - f[i][a])
        .slice(0, 1)
        .pop(),
    )
    .join("");
};

const second = i => null;

module.exports = { first, second };
