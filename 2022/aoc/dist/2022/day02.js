"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const file_1 = require("./utils/file");
(function solveDay01() {
    return __awaiter(this, void 0, void 0, function* () {
        const array = yield (0, file_1.readFileAsArray)("./2022/data/02/input.txt");
        const rpsLHS = {
            A: "r",
            B: "p",
            C: "s",
        };
        const rpsRHS = {
            X: "r",
            Y: "p",
            Z: "s",
        };
        const scores = {
            r: 1,
            p: 2,
            s: 3,
        };
        const totalScore = array.reduce(({ partOneScore }, gameData, index) => {
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
        }, { partOneScore: 0 });
        const part2 = array.reduce(({ score }, gameData, index) => {
            const [lhs, rhs] = gameData.split(" ");
            let moveRHS = "", moveScore = NaN;
            switch (rhs) {
                case "Y": //Draw
                    moveRHS = rpsLHS[lhs];
                    moveScore = 3;
                    break;
                case "X": //Lose
                    moveScore = 0;
                    if (rpsLHS[lhs] == "r") {
                        moveRHS = "s";
                    }
                    else if (rpsLHS[lhs] == "p") {
                        moveRHS = "r";
                    }
                    else if (rpsLHS[lhs] == "s") {
                        moveRHS = "p";
                    }
                    break;
                case "Z":
                    moveScore = 6;
                    if (rpsLHS[lhs] == "r") {
                        moveRHS = "p";
                    }
                    else if (rpsLHS[lhs] == "p") {
                        moveRHS = "s";
                    }
                    else if (rpsLHS[lhs] == "s") {
                        moveRHS = "r";
                    }
                    break;
            }
            return {
                score: score + moveScore + scores[moveRHS],
            };
        }, { score: 0 });
        console.log(part2);
    });
})();
//# sourceMappingURL=day02.js.map