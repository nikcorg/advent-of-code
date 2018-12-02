import { left, right } from "fp-ts/lib/Either";
import { compose } from "fp-ts/lib/function";
import { fromEither, taskEither, TaskEither } from "fp-ts/lib/TaskEither";
import * as path from "path";
import { ensureError, readFile, sequenceAll, toStringArray, zip, zip3 } from "../shared";

const hasTwins = (s: string[]): boolean => zip(s, s.slice(1)).some(([a, b]) => a === b);

const hasTriplets = (s: string[]): boolean =>
  zip3(s, s.slice(1), s.slice(2)).some(([a, b, c]) => a === b && c === a);

const mul = (a: number, b: number) => a * b;

const sort = <A>(xs: A[]): A[] => [...xs].sort();

const stringToArray = (s: string) => [...s];

const stringToSortedArray = compose<string, string[], string[]>(
  sort,
  stringToArray,
);

const first = (input: string[]) =>
  taskEither.of<Error, string[]>(input).map<number>(lines =>
    lines
      .map(stringToSortedArray)
      .map(line => [hasTwins(line) ? 1 : 0, hasTriplets(line) ? 1 : 0])
      .reduce(([acc2, acc3], [h2, h3]) => [acc2 + h2, acc3 + h3], [0, 0])
      .reduce(mul),
  );

const second = (input: string[]): TaskEither<Error, string> =>
  taskEither
    .of<Error, string[]>(input)
    .map(lines => [...lines].sort())
    .map((sorted: string[]) => {
      for (const [i, candidateA] of sorted.entries()) {
        for (const candidateB of sorted.slice(i)) {
          const diff = zip(candidateA.split(""), candidateB.split("")).reduce(
            (acc, [a, b]) => (a === b ? `${acc}${a}` : acc),
            "",
          );
          if (candidateA.length - diff.length === 1) {
            return right<Error, string>(diff);
          }
        }
      }
      return left<Error, string>(new Error("No match found"));
    })
    .chain(fromEither);

readFile(path.join(path.dirname(__filename), "input.txt"))
  .map(toStringArray)
  .chain(xs => sequenceAll<Error, string | number>([first(xs), second(xs)]))
  .run()
  // tslint:disable-next-line:no-console
  .then(rE => {
    if (rE.isRight()) {
      const [fst, snd] = rE.value;
      // tslint:disable-next-line:no-console
      console.log("first=%s, second=%s", fst, snd);
    } else {
      throw ensureError(rE.value);
    }
  })
  .catch((err: Error) => {
    // tslint:disable-next-line:no-console
    console.error("Failed with", err);
  });
