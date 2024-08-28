require("source-map-support").install();

import { readFileAsArray } from "./utils/file";

let nodes: any;

function walkNodes(
  node: any,
  cost: number,
  directionStack: string,
  seenNodes: any[]
) {
  const sameThreeDirs = directionStack.match(/.*(.)\1{2}$/);

  const endingDirection = directionStack.slice(-1);

  // if (seenNodes.find(({ x, y }) => node.x == x && node.y == y)) {
  //   console.log("Ignoring vistied node", cost);

  //   return;
  // }
  if(node.paths["a"] < cost) {
    return;
  }
  seenNodes = [...seenNodes, node];


  node.allowedDirections.forEach((d: any) => {
    let nextNode;
    switch (d) {
      case "r":
        nextNode = nodes[node.x][node.y + 1];
        break;
      case "l":
        nextNode = nodes[node.x][node.y - 1];
        break;

      case "t":
        nextNode = nodes[node.x - 1][node.y];
        break;

      case "b":
        nextNode = nodes[node.x + 1][node.y];
        break;
    }

    if (endingDirection == d && sameThreeDirs) {
      return;
    } else {
      const nextCost = cost + node.cost;
      nextNode.paths["a"] =Math.min(
          nextNode.paths["a"] || Number.POSITIVE_INFINITY,
          nextCost
        );
        
      const nextDirectionStack = directionStack + d;

      walkNodes(nextNode, nextCost, nextDirectionStack, [...seenNodes]);
    }
  });
}

(async function solveDay01() {
  const array = await readFileAsArray("./2023/data/17/example.txt");

  nodes = Array(array.length).fill(Array(array.length).fill(0));

  for (var line = 0; line < array.length; line++) {
    nodes[line] = array[line].split("").map((v, i) => {
      let removedDirections = [];
      const x = line;
      const y = i;
      if (x == 0) {
        removedDirections.push("t");
      }
      if (x == array.length - 1) {
        removedDirections.push("b");
      }
      if (y == array.length - 1) {
        removedDirections.push("r");
      }
      if (y == 0) {
        removedDirections.push("l");
      }
      return {
        cost: parseInt(v, 10),
        x,
        y,
        allowedDirections: ["t", "r", "b", "l"].filter(
          (v: any) => !removedDirections.includes(v)
        ),
        paths: {},
      };
    });
  }

  console.log(nodes[0][0]);
  walkNodes(nodes[0][0], 0, "", []);

  for (var line = 0; line < array.length; line++) {
    for (var j = 0; j < array.length; j++) {
      console.log(`${line},${j}`, nodes[line][j].paths);
    }
  }

  console.log(nodes[array.length - 1][array.length - 1].paths);

  //
})();

