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
        const array = (yield (0, file_1.readFileAsArray)("./2022/data/04/input.txt")).map((v) => {
            const match = v.match(/(\d+)-(\d+),(\d+)-(\d+)/);
            return {
                lhsStart: parseInt(match[1]),
                lhsEnd: parseInt(match[2]),
                rhsStart: parseInt(match[3]),
                rhsEnd: parseInt(match[4]),
            };
        });
        const part01 = array.reduce((sum, item) => {
            if (item.lhsStart == item.rhsStart && item.lhsEnd == item.rhsEnd) {
                sum += 1;
            }
            else if (item.lhsStart <= item.rhsStart && item.lhsEnd >= item.rhsEnd) {
                //  2 3 4 5
                //    3 4
                sum += 1;
            }
            else if (item.rhsStart <= item.lhsStart && item.rhsEnd >= item.lhsEnd) {
                //    3 4
                //  2 3 4 5
                sum += 1;
            }
            return sum;
        }, 0);
        const part02 = array.reduce((sum, item) => {
            if (item.lhsStart == item.rhsStart && item.lhsEnd == item.rhsEnd) {
                sum += 1;
            }
            else if (item.lhsStart <= item.rhsStart && item.lhsEnd >= item.rhsEnd) {
                //  2 3 4 5
                //    3 4
                sum += 1;
            }
            else if (item.rhsStart <= item.lhsStart && item.rhsEnd >= item.lhsEnd) {
                //    3 4
                //  2 3 4 5
                sum += 1;
            }
            else if (item.lhsEnd >= item.rhsStart && item.lhsStart <= item.rhsEnd) {
                //    3 4
                //        5 6
                //     3 4
                // 1 2
                //    3 4
                //      4 5 6
                //    3 4
                //    3 4 5 6
                sum += 1;
            }
            else if (item.rhsEnd >= item.lhsStart && item.rhsStart <= item.lhsEnd) {
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
    });
})();
//# sourceMappingURL=day04.js.map