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
        const array = yield (0, file_1.readFileAsArray)("./2022/data/05/input.txt");
        const use9001 = true;
        const result = array.reduce(({ stacks, parsedStacks }, item) => {
            if (item.match(/^\s\d\s/)) {
                parsedStacks = true;
            }
            else if (parsedStacks) {
                const instruction = item.match(/move (\d+) from (\d+) to (\d+)/);
                if (instruction) {
                    const count = parseInt(instruction[1]);
                    const fromStack = parseInt(instruction[2]);
                    const toStack = parseInt(instruction[3]);
                    console.log({ count, fromStack, toStack });
                    const removedItems = stacks[fromStack - 1].splice(-1 * count, count);
                    stacks[toStack - 1].push(...(use9001 ? removedItems : removedItems.reverse()));
                }
            }
            else if (item.match(/\[([A-Z])\]/)) {
                const crateMatches = Array.from(item.matchAll(/\[([A-Z])\]/g));
                console.log(crateMatches);
                crateMatches.forEach((m, i) => {
                    console.log({ m });
                    const matched = m[1];
                    const { index } = m;
                    const stackIndex = Math.floor(index / 4) + 1;
                    console.log(stackIndex);
                    while (stackIndex > stacks.length) {
                        stacks.push([]);
                    }
                    stacks[stackIndex - 1].splice(Number.NEGATIVE_INFINITY, 0, matched);
                });
            }
            console.log(stacks);
            return { stacks, parsedStacks };
        }, { stacks: [], parsedStacks: false });
        console.log(result.stacks.reduce((a, item) => {
            return a + item[item.length - 1];
        }, ""));
    });
})();
//# sourceMappingURL=day05.js.map