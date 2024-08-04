import { readFileAsArray } from "./utils/file";

type Obj = Record<string, string>;

(async function solveDay01() {
  const array = await readFileAsArray("./2022/data/05/input.txt");
  const use9001 = true;
  const result = array.reduce(
    ({ stacks, parsedStacks }, item) => {
      if (item.match(/^\s\d\s/)) {
        parsedStacks = true;
      } else if (parsedStacks) {
        const instruction = item.match(/move (\d+) from (\d+) to (\d+)/);
        if (instruction) {
          const count = parseInt(instruction[1]);
          const fromStack = parseInt(instruction[2]);
          const toStack = parseInt(instruction[3]);

          console.log({ count, fromStack, toStack });

          const removedItems = (stacks[fromStack - 1] as string[][]).splice(
            -1 * count,
            count
          );
          (stacks[toStack - 1] as string[][]).push(
            ...(use9001 ? removedItems : removedItems.reverse())
          );
        }
      } else if (item.match(/\[([A-Z])\]/)) {
        const crateMatches = Array.from(item.matchAll(/\[([A-Z])\]/g));

        console.log(crateMatches);

        crateMatches.forEach((m, i) => {
          console.log({ m });

          const matched = m[1];
          const { index } = m;
          const stackIndex = Math.floor((index as number) / 4) + 1;
          console.log(stackIndex);

          while (stackIndex > stacks.length) {
            (stacks as string[][]).push([]);
          }
          (stacks[stackIndex - 1] as string[]).splice(
            Number.NEGATIVE_INFINITY,
            0,
            matched
          );
        });
      }
      console.log(stacks);

      return { stacks, parsedStacks };
    },
    { stacks: [], parsedStacks: false }
  );
  console.log(
    result.stacks.reduce((a, item) => {
      return a + item[item.length - 1];
    }, "")
  );
})();
