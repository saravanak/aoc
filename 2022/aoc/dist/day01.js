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
const fs_1 = require("fs");
(function solveDay01() {
    return __awaiter(this, void 0, void 0, function* () {
        const contents = yield fs_1.promises.readFile("./2022/data/01/input.txt");
        const array = contents.toString().replace(/\r\n/g, "\n").split("\n");
        console.log(array.length);
        const result = array.reduce(({ maxCalories, currentCaloriesCount, elfIndex, maxCaloriesList, }, item) => {
            const currentItem = item == "" ? NaN : parseInt(item);
            if (isNaN(currentItem)) {
                maxCaloriesList.push(currentCaloriesCount);
            }
            return {
                maxCalories: isNaN(currentItem)
                    ? Math.max(maxCalories, currentCaloriesCount)
                    : maxCalories,
                currentCaloriesCount: isNaN(currentItem)
                    ? 0
                    : currentCaloriesCount + currentItem,
                elfIndex: isNaN(currentItem) ? elfIndex + 1 : elfIndex,
                maxCaloriesList,
            };
        }, {
            maxCalories: 0,
            currentCaloriesCount: 0,
            elfIndex: 0,
            maxCaloriesList: [0],
        });
        result.maxCaloriesList.sort((a, b) => (a > b ? -1 : 0));
        console.log({
            max3Calories: result,
            sum: result.maxCaloriesList[1] +
                result.maxCaloriesList[2] +
                result.maxCaloriesList[0],
        });
    });
})();
