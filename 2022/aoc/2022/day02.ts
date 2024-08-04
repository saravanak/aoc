import { readFileAsArray } from "./utils/file";

type Obj = Record<string, string>;

(async function solveDay01() {
  const array = await readFileAsArray("./2022/data/02/input.txt");
  const rpsLHS: Obj = {
    A: "r",
    B: "p",
    C: "s",
  };
  const rpsRHS: Obj = {
    X: "r",
    Y: "p",
    Z: "s",
  };

  const scores: Record<string, number> = {
    r: 1,
    p: 2,
    s: 3,
  };

  const totalScore = array.reduce(
    ({ partOneScore }, gameData, index) => {
      const [lhs, rhs] = gameData.split(" ");

      if (rpsLHS[lhs] == rpsRHS[rhs]) {
        partOneScore += 3 + scores[rpsRHS[rhs]];
        return { partOneScore };
      }

      let isWon = false;
      switch (rpsRHS[rhs]) {
        case "r":
          isWon = rpsLHS[lhs] == "s";
          break;
        case "p":
          isWon = rpsLHS[lhs] == "r";
          break;
        case "s":
          isWon = rpsLHS[lhs] == "p";
          break;
      }
      return {
        partOneScore: partOneScore + (isWon ? 6 : 0) + scores[rpsRHS[rhs]],
      };
    },
    { partOneScore: 0 }
  );
  const part2 = array.reduce(
    ({ score }, gameData, index) => {
      const [lhs, rhs] = gameData.split(" ");

      let moveRHS = "",
        moveScore = NaN;
      switch (rhs) {
        case "Y": //Draw
          moveRHS = rpsLHS[lhs];
          moveScore = 3;
          break;
        case "X": //Lose
          moveScore = 0;
          if (rpsLHS[lhs] == "r") {
            moveRHS = "s";
          } else if (rpsLHS[lhs] == "p") {
            moveRHS = "r";
          } else if (rpsLHS[lhs] == "s") {
            moveRHS = "p";
          }

          break;
        case "Z":
          moveScore = 6;
          if (rpsLHS[lhs] == "r") {
            moveRHS = "p";
          } else if (rpsLHS[lhs] == "p") {
            moveRHS = "s";
          } else if (rpsLHS[lhs] == "s") {
            moveRHS = "r";
          }
          break;
      }
      return {
        score: score + moveScore + scores[moveRHS],
      };
    },
    { score: 0 }
  );
  console.log(part2);
})();
