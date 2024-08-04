
import { readFileAsArray} from "./utils/file";

(async function solveDay01() {
  const array = await readFileAsArray("./2022/data/01/input.txt");

  const result = array.reduce(
    (
      { maxCalories, currentCaloriesCount, elfIndex, maxCaloriesList },
      item
    ) => {
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
    },
    {
      maxCalories: 0,
      currentCaloriesCount: 0,
      elfIndex: 0,
      maxCaloriesList: [0],
    }
  );

  result.maxCaloriesList.sort((a, b) => (a > b ? -1 : 0));
  console.log({
    max3Calories: result,
    sum:
      result.maxCaloriesList[1] +
      result.maxCaloriesList[2] +
      result.maxCaloriesList[0],
  });
})();
