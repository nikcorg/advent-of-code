const test = require('tape');
const { first } = require('./solve');

const testInput = `
eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`.trim().split('\n');

test('day 6', t => {
  t.test('test case', t => {
    t.equal(first(testInput), 'easter');
    t.end();
  });
});
