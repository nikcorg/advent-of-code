const {
  compose,
  zip,
  split
} = require('../../utils');

// An IP supports TLS if it has an Autonomous Bridge Bypass Annotation, or ABBA. An ABBA is any four-character sequence which consists of a pair of two different characters followed by the reverse of that pair, such as xyyx or abba. However, the IP also must not have an ABBA within any hypernet sequences, which are contained by square brackets.

const hasAbba = compose(
  cs => zip(cs, cs.slice(1)),
  split(''));

const first = i => hasAbba(i);

const second = i => null;

module.exports = { first, second };
