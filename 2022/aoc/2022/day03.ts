import { textChangeRangeIsUnchanged } from "typescript";
import { readFileAsArray } from "./utils/file";

type Obj = Record<string, string>;

(async function solveDay01() {
  const array = await readFileAsArray("./2022/data/03/input.txt");
  const priority = (bag: string) => {
    const ascii = bag.charCodeAt(0);
    if (ascii < 97) {
      //Uppercase
      return ascii - 65 + 27;
    }
    if (ascii >= 97) {
      //Uppercase
      return ascii - 97 + 1;
    }
    return 0;
  };

  const part1 = array.reduce((a, item) => {
    const [lhs, rhs] = [
      item.slice(0, item.length / 2),
      item.slice(item.length / 2),
    ];

    const lhsRegex = new RegExp(`[${lhs}]`);

    const commonBag = rhs.match(lhsRegex)[0];
    return a + priority(commonBag);
  }, 0);

  const part2 = array.reduce(
    ({ currentGroup, sum }, item, index) => {
      switch (index % 3) {
        case 0:
          //process previous group

          currentGroup = [item];
          break;
        case 1:
          currentGroup.push(item);
          break;
        case 2:
          currentGroup.push(item);
          if (currentGroup.length == 3) {
            const [first, second, third]: string[] = currentGroup;
            const firstRegex = new RegExp(`[${first}]`, "g");
            const allMatches = second.matchAll(firstRegex);
            const matchedChars = Array.from(allMatches).reduce(
              (matchedChars, match) => {
                return matchedChars + match[0];
              },
              ""
            );
            const matchedRegex = new RegExp(`[${matchedChars}]`);
            const commonBag = third.match(matchedRegex)[0];
            sum = sum + priority(commonBag);
          }
          break;
      }
      return { currentGroup, sum };
    },
    { currentGroup: [], sum: 0 }
  );
  console.log(part2.sum);
})();
