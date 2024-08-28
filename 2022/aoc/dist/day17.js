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
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
exports.__esModule = true;
require("source-map-support").install();
var file_1 = require("./utils/file");
var nodes;
function walkNodes(node, cost, directionStack, seenNodes) {
    var sameThreeDirs = directionStack.match(/.*(.)\1{2}$/);
    var endingDirection = directionStack.slice(-1);
    // if (seenNodes.find(({ x, y }) => node.x == x && node.y == y)) {
    //   console.log("Ignoring vistied node", cost);
    //   return;
    // }
    if (node.paths["a"] < cost) {
        return;
    }
    seenNodes = __spreadArray(__spreadArray([], seenNodes, true), [node], false);
    node.allowedDirections.forEach(function (d) {
        var nextNode;
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
        }
        else {
            var nextCost = cost + node.cost;
            nextNode.paths["a"] = Math.min(nextNode.paths["a"] || Number.POSITIVE_INFINITY, nextCost);
            var nextDirectionStack = directionStack + d;
            walkNodes(nextNode, nextCost, nextDirectionStack, __spreadArray([], seenNodes, true));
        }
    });
}
(function solveDay01() {
    return __awaiter(this, void 0, void 0, function () {
        var array, line, line, j;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, (0, file_1.readFileAsArray)("./2023/data/17/example.txt")];
                case 1:
                    array = _a.sent();
                    nodes = Array(array.length).fill(Array(array.length).fill(0));
                    for (line = 0; line < array.length; line++) {
                        nodes[line] = array[line].split("").map(function (v, i) {
                            var removedDirections = [];
                            var x = line;
                            var y = i;
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
                                x: x,
                                y: y,
                                allowedDirections: ["t", "r", "b", "l"].filter(function (v) { return !removedDirections.includes(v); }),
                                paths: {}
                            };
                        });
                    }
                    console.log(nodes[0][0]);
                    walkNodes(nodes[0][0], 0, "", []);
                    for (line = 0; line < array.length; line++) {
                        for (j = 0; j < array.length; j++) {
                            console.log("".concat(line, ",").concat(j), nodes[line][j].paths);
                        }
                    }
                    console.log(nodes[array.length - 1][array.length - 1].paths);
                    return [2 /*return*/];
            }
        });
    });
})();
