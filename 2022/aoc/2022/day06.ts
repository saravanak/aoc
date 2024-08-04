import { freemem } from "os";
import { readFileAsArray } from "./utils/file";

type Obj = Record<string, string>;

(async function solveDay01() {
  const array = await readFileAsArray("./2022/data/06/input.txt");
  const part01DistinctChars = 4;
  const part02DistinctChars = 14;

  const DISTICT_CHARS = part02DistinctChars;
  const result = array.reduce(
    ({ answers }, item, _, list) => {
      const chars = item.split("").reduce(
        ({ freqCount, isFound, markerIndex }, char, index, charList) => {
          if (index < DISTICT_CHARS - 1) {
            return {
              freqCount,
              isFound,
              markerIndex,
            };
          }
          if (index == DISTICT_CHARS - 1) {
            const window = charList.slice(0, index + 1);

            const freqCount = window.reduce((freq, windowChar) => {
              freq[windowChar] = (freq[windowChar] || 0) + 1;
              return freq;
            }, {});

            return { freqCount, isFound, markerIndex };
          }

          if (isFound) {
            return {
              freqCount,
              isFound,
              markerIndex,
            };
          } else {
            const outGoingCharacter = charList[index - DISTICT_CHARS];
            freqCount[outGoingCharacter] =
              (freqCount[outGoingCharacter] || 0) - 1;
            freqCount[char] = (freqCount[char] || 0) + 1;

            const isFound = Object.entries(freqCount).every(
              ([_, count]) => count == 1 || count == 0
            );
 
            // console.log({ char, index, outGoingCharacter });

            return {
              freqCount,
              isFound,
              markerIndex: index + 1,
            };
          }
        },
        { freqCount: {}, isFound: false, markerIndex: NaN }
      );
      answers.push(chars.markerIndex);
      return { answers };
    },
    { answers: [] }
  );

  console.log({ result });
})();
