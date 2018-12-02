import { Either, right } from "fp-ts/lib/Either";
import { fromEither, taskEither } from "fp-ts/lib/TaskEither";
import * as path from "path";
import { ensureError, map, readFile, reduce, sequenceAll, toStringArray } from "../shared";

const add = (a: number, b: number) => a + b;

const first = (input: number[]) => taskEither.of<Error, number[]>(input).map(reduce(add)(0));

const second = (input: number[]) =>
  taskEither
    .of<Error, number[]>(input)
    .map<Either<Error, number>>(nums => {
      const sums = new Set();
      const x = nums.reduce((sum, n) => {
        const nextSum = sum + n;
        sums.add(nextSum);
        return nextSum;
      }, 0);

      let repeat = x;
      while (true) {
        for (const n of nums) {
          repeat += n;
          if (sums.has(repeat)) {
            return right(repeat);
          }
        }
      }
    })
    .chain(fromEither);

readFile(path.join(path.dirname(__filename), "input.txt"))
  .map(toStringArray)
  .map(map(Number))
  .chain(lines => sequenceAll([first(lines), second(lines)]))
  .run()
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
