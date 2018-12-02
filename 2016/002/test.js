const test = require("tape");
const { first, second } = require("./solve");
const input = ["ULL", "RRDDD", "LURDL", "UUUUD"];

test("day 2", t => {
  t.test("test case 1", t => {
    const solution = "1985";
    t.equal(first(input), solution);
    t.end();
  });

  t.test("test case 2", t => {
    const solution = "5DB3";
    t.equal(second(input), solution);
    t.end();
  });
});
