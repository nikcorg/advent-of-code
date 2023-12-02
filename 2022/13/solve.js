const fs = require("fs");

const zip = (as, bs) => {
  cs = [];
  for (let i = 0; i < Math.max(as.length, bs.length); i++) {
    cs.push([as[i], bs[i]]);
  }
  return cs;
};

function compare([left, right]) {
  if (left == null && right == null) {
    return 0;
  }

  if (left == null && right != null) {
    return 1;
  }

  if (left != null && right == null) {
    return -1;
  }

  if (Array.isArray(left) && !Array.isArray(right)) {
    return compare([left, [right]]);
  } else if (!Array.isArray(left) && Array.isArray(right)) {
    return compare([[left], right]);
  }

  if (Array.isArray(left) && Array.isArray(right)) {
    return zip(left, right).reduce(
      (acc, next) => (acc != 0 ? acc : compare(next)),
      0
    );
  }

  if (left > right) {
    return -1;
  } else if (left < right) {
    return 1;
  }

  return 0;
}

// const input = fs.readFileSync("./input_test.txt", "utf-8");
const input = fs.readFileSync("./input.txt", "utf-8");

const pkgs = input
  .split("\n")
  .filter((s) => s.length > 0)
  .map((p) => JSON.parse(p.trim()));

function solveFirst(pkgs) {
  let valid = 0;

  for (let i = 1, n = 1; i < pkgs.length; i += 2, n++) {
    const left = pkgs[i - 1];
    const right = pkgs[i];
    const good = compare([left, right]) > 0;

    if (good) {
      valid += n;
    }
  }

  return valid;
}

// stupidly got my comparison outcome the wrong way and couldn't be bothered to change it
const comparisonToSort = {
  [-1]: 1,
  1: -1,
  0: 0,
};

function solveSecond(pkgs, ds) {
  const sorted = pkgs
    .concat(ds)
    .sort((a, b) => comparisonToSort[compare([a, b])]);

  const ps = ds.map((d) => {
    const sd = JSON.stringify(d);
    return sorted.findIndex((x) => JSON.stringify(x) === sd) + 1;
  });

  return ps[0] * ps[1];
}

const first = solveFirst(pkgs);
const second = solveSecond(pkgs, [[[2]], [[6]]]);

console.log({ first, second });
