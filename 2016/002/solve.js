const {
  compose,
  join,
  map,
  reduce,
  split,
  tap,
  equals,
  not,
  lookup,
  constrain,
} = require("../../utils");

const motionMap = {
  U: ([x, y]) => [x, y - 1],
  D: ([x, y]) => [x, y + 1],
  L: ([x, y]) => [x - 1, y],
  R: ([x, y]) => [x + 1, y],
};

const posToNumberMap = {
  "0,0": 1,
  "1,0": 2,
  "2,0": 3,
  "0,1": 4,
  "1,1": 5,
  "2,1": 6,
  "0,2": 7,
  "1,2": 8,
  "2,2": 9,
};

const _ = null,
  A = "A",
  B = "B",
  C = "C",
  D = "D";

const posToNumberMap2 = {
  "0,0": _,
  "1,0": _,
  "2,0": 1,
  "3,0": _,
  "4,0": _,
  "0,1": _,
  "1,1": 2,
  "2,1": 3,
  "3,1": 4,
  "4,1": _,
  "0,2": 5,
  "1,2": 6,
  "2,2": 7,
  "3,2": 8,
  "4,2": 9,
  "0,3": _,
  "1,3": A,
  "2,3": B,
  "3,3": C,
  "4,3": _,
  "0,4": _,
  "1,4": _,
  "2,4": D,
  "3,4": _,
  "4,4": _,
};

const translate = from => compose(lookup(from), join(","));
const motion = lookup(motionMap);
const step = (from, move) => motion(move)(from);

const traverse = (step, init) => steps => {
  let from = init;
  return steps.map(motions => (from = motions.reduce(step, from)));
};

const stepIfValid = (isValid, step) => (from, move) => {
  let r = step(from, move);
  return isValid(r) ? r : from;
};

const notEmpty = compose(not, equals(_));

const isValid = compose(notEmpty, translate(posToNumberMap2));

const makeSolver = (translate, traverse) =>
  compose(join(""), map(translate), traverse, map(split("")));

const step1 = compose(map(constrain(0, 2)), step);
const first = makeSolver(translate(posToNumberMap), traverse(step1, [1, 1]));

const step2 = compose(map(constrain(0, 4)), step);
const second = makeSolver(
  translate(posToNumberMap2),
  traverse(stepIfValid(isValid, step2), [0, 2]),
);

module.exports = { first, second };
