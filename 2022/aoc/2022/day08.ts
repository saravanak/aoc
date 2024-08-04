require("source-map-support").install();

import { readFileAsArray } from "./utils/file";

(async function solveDay01() {
  const array = await readFileAsArray("./2022/data/08/input.txt");

  console.log("AAAAA-AAAAA-AAAAA-AAAA//////////////");

  const grid = array.map((item) => {
    return item.split("").map((v) => {
      return {
        height: parseInt(v),
        tallestLeft: 0,
        tallestRight: 0,
        tallestTop: 0,
        tallestBottom: 0,
      };
    });
  });

  const GRID_SIZE = grid.length;
  const GRID_INDEX_UPPER = GRID_SIZE - 1;
  for (var row = 0; row < grid.length; row++) {
    for (var col = 0; col < grid.length; col++) {
      grid[row][col].tallestLeft =
        col == 0
          ? -1
          : Math.max(grid[row][col - 1].tallestLeft, grid[row][col - 1].height);

      console.log(row, col);

      grid[row][GRID_INDEX_UPPER - col].tallestRight =
        col == 0
          ? -1
          : Math.max(
              grid[row][GRID_INDEX_UPPER - col + 1].tallestRight,
              grid[row][GRID_INDEX_UPPER - col + 1].height
            );

      grid[row][col].tallestTop =
        row == 0
          ? -1
          : Math.max(grid[row - 1][col].tallestTop, grid[row - 1][col].height);

      grid[GRID_INDEX_UPPER - row][col].tallestBottom =
        row == 0
          ? -1
          : Math.max(
              grid[GRID_INDEX_UPPER - row + 1][col].tallestBottom,
              grid[GRID_INDEX_UPPER - row + 1][col].height
            );
    }
  }

  const result = grid.reduce(
    ({ visibleNodes }, row, index) => {
      console.log(`Row = ${index + 1} ////////////`);

      visibleNodes.push(
        ...row.filter((item) => {
          const isVisible =
            item.height > item.tallestBottom ||
            item.height > item.tallestLeft ||
            item.height > item.tallestRight ||
            item.height > item.tallestTop;

          // console.log({ ...item, isVisible });

          return isVisible;
        })
      );
      return {
        visibleNodes,
      };
    },
    { visibleNodes: [] }
  );

  console.log(result.visibleNodes.length);
})();
