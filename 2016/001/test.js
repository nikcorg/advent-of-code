const test = require("tape");
const { first, second, intersection } = require("./solve");

test("intersection", t => {
  t.test("intersects", t => {
    const v1 = [[2, 0], [2, 10]];
    const v2 = [[0, 3], [5, 3]];

    t.deepEqual(intersection(v1, v2), [2, 3]);
    t.end();
  });

  t.test("no intersection", t => {
    const v1 = [[2, 4], [2, 10]];
    const v2 = [[0, 3], [5, 3]];

    t.equal(intersection(v1, v2), null);
    t.end();
  });
});

test("day 1", t => {
  t.test("provided test case 1", t => {
    const input = "R2, L3";
    const expected = 5;

    t.equal(first(input), expected);
    t.end();
  });

  t.test("provided test case 2", t => {
    const input = "R2, R2, R2";
    const expected = 2;

    t.equal(first(input), expected);
    t.end();
  });

  t.test("provided test case 3", t => {
    const input = "R5, L5, R5, R3";
    const expected = 12;

    t.equal(first(input), expected);
    t.end();
  });

  t.test("provided test case 4", t => {
    const input = "R8, R4, R4, R8";
    const expected = 4;

    t.equal(second(input), expected);
    t.end();
  });
});
