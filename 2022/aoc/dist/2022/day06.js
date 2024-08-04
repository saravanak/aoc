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
        const array = yield (0, file_1.readFileAsArray)("./2022/data/06/input.txt");
        const part01DistinctChars = 4;
        const part02DistinctChars = 14;
        const DISTICT_CHARS = part02DistinctChars;
        const result = array.reduce(({ answers }, item, _, list) => {
            const chars = item.split("").reduce(({ freqCount, isFound, markerIndex }, char, index, charList) => {
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
                }
                else {
                    const outGoingCharacter = charList[index - DISTICT_CHARS];
                    freqCount[outGoingCharacter] =
                        (freqCount[outGoingCharacter] || 0) - 1;
                    freqCount[char] = (freqCount[char] || 0) + 1;
                    const isFound = Object.entries(freqCount).every(([_, count]) => count == 1 || count == 0);
                    // console.log({ char, index, outGoingCharacter });
                    return {
                        freqCount,
                        isFound,
                        markerIndex: index + 1,
                    };
                }
            }, { freqCount: {}, isFound: false, markerIndex: NaN });
            answers.push(chars.markerIndex);
            return { answers };
        }, { answers: [] });
        console.log({ result });
    });
})();
//# sourceMappingURL=day06.js.map