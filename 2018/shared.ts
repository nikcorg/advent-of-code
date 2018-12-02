import { array } from "fp-ts/lib/Array";
import { Either, left, right } from "fp-ts/lib/Either";
import { Task } from "fp-ts/lib/Task";
import { TaskEither, taskEither } from "fp-ts/lib/TaskEither";
import * as fs from "fs";

export const assertNotNull = <A>(x: A | null | undefined): A => {
  if (x == null) {
    throw new Error("Expected non-null value");
  }
  return x;
};

export const ensureError = (err: Error | string): Error =>
  err instanceof Error ? err : new Error(err);

export const flatten = <A>(xs: A[]): A[] => xs.reduce((ys: A[], x) => ys.concat(x), []);

export const map = <A, B>(fn: (a: A) => B) => (data: A[]): B[] => data.map(fn);

export const readFile = (fn: string): TaskEither<Error, string> =>
  new TaskEither(
    new Task<Either<Error, string>>(
      () =>
        new Promise((res, rej) =>
          fs.readFile(fn, "utf-8", (err: Error, data: string | Buffer) =>
            err
              ? rej(left<Error, string>(ensureError(err)))
              : res(right<Error, string>(data.toString())),
          ),
        ),
    ),
  );

export const reduce = <A, B>(fn: (a: A, b: B) => A) => (init: A) => (data: B[]) =>
  data.reduce(fn, init);

export const sequenceAll = <L, A>(xs: Array<TaskEither<L, A>>) => array.sequence(taskEither)(xs);

export const toStringArray = (str: string): string[] => str.split("\n");

export const zip = <A, B>(as: A[], bs: B[]): Array<Array<A | B>> =>
  as.length < bs.length ? as.map((a, i) => [a, bs[i]]) : bs.map((b, i) => [b, as[i]]);

export const zip3 = <A, B, C>(as: A[], bs: B[], cs: C[]) => zip(zip(as, bs), cs).map(flatten);
