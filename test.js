const test = require('tape');
const sinon = require('sinon');

test('compose', t => {
  const { compose } = require('./utils');
  const id = a => a;
  const append = a => s => `${s}${a}`;

  t.test('calls first', t => {
    const spy = sinon.spy(id);
    const prog = compose(spy);

    prog('hello');

    t.equal(spy.callCount, 1);
    t.equal(spy.args[0][0], 'hello');
    t.end();
  });

  t.test('calls first with multiple args', t => {
    const spy = sinon.spy(id);
    const prog = compose(spy);

    prog('hello', 'world');

    t.equal(spy.callCount, 1);
    t.deepEqual(spy.args[0], ['hello', 'world']);
    t.end();
  });

  t.test('calls all funcs', t => {
    const spy1 = sinon.spy(id);
    const spy2 = sinon.spy(id);
    const prog = compose(spy2, spy1);

    prog('beep');

    t.equal(spy1.callCount, 1);
    t.equal(spy2.callCount, 1);
    t.end();
  });

  t.test('return value is passed to next function', t => {
    const spy1 = sinon.spy(append('b'));
    const spy2 = sinon.spy(append('c'));
    const prog = compose(spy2, spy1);

    const res = prog('a');

    t.equal(spy1.args[0][0], 'a');
    t.equal(spy2.args[0][0], 'ab');
    t.equal(res, 'abc');
    t.end();
  });
});
