import { taskEither } from "fp-ts/lib/TaskEither";
import * as path from "path";
import { readFile, toStringArray } from "../shared";

const first = (input: string[]) => taskEither.of(input).map(lines => {});

readFile(path.join(__dirname, "input.txt"))
  .map(toStringArray)
  .map(first);
