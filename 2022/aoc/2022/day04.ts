import { textChangeRangeIsUnchanged } from "typescript";
import { readFileAsArray } from "./utils/file";

type Obj = Record<string, string>;

(async function solveDay01() {
  const array = (await readFileAsArray("./2022/data/04/input.txt")).map(
    (v) => {
      const match = v.match(/(\d+)-(\d+),(\d+)-(\d+)/);
      return {
        lhsStart: parseInt(match[1]),
        lhsEnd: parseInt(match[2]),
        rhsStart: parseInt(match[3]),
        rhsEnd: parseInt(match[4]),
      };
    }
  );
  const part01 = array.reduce((sum, item) => {
    if (item.lhsStart == item.rhsStart && item.lhsEnd == item.rhsEnd) {
      sum += 1;
    } else if (item.lhsStart <= item.rhsStart && item.lhsEnd >= item.rhsEnd) {
      //  2 3 4 5
      //    3 4
      sum += 1;
    } else if (item.rhsStart <= item.lhsStart && item.rhsEnd >= item.lhsEnd) {
      //    3 4
      //  2 3 4 5
      sum += 1;
    }
    return sum;
  }, 0);
  const part02 = array.reduce((sum, item) => {
    if (item.lhsStart == item.rhsStart && item.lhsEnd == item.rhsEnd) {
      sum += 1;
    } else if (item.lhsStart <= item.rhsStart && item.lhsEnd >= item.rhsEnd) {
      //  2 3 4 5
      //    3 4
      sum += 1;
    } else if (item.rhsStart <= item.lhsStart && item.rhsEnd >= item.lhsEnd) {
      //    3 4
      //  2 3 4 5
      sum += 1;
    } else if (item.lhsEnd >= item.rhsStart && item.lhsStart <= item.rhsEnd) {
      //    3 4
      //        5 6

      //     3 4
      // 1 2
      

      //    3 4
      //      4 5 6

      //    3 4
      //    3 4 5 6
      sum += 1;
    } else if (item.rhsEnd >= item.lhsStart && item.rhsStart <= item.lhsEnd) {
      //        5 6
      //    3 4

      //      4 5 6
      //    3 4

      //    3 4 5 6
      //    3 4

      sum += 1;
    }

    return sum;
  }, 0);
  console.log(part02);
})();
