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
require("source-map-support").install();
const file_1 = require("./utils/file");
function findSize(dir) {
    let dirSize = 0;
    // console.log(Object.keys(dir), dir);
    if (dir.type == "file") {
        return dir.size;
    }
    if (dir._dir_size != undefined) {
        Object.keys(dir).forEach((childFile) => {
            const item = dir[childFile];
            if (childFile == "_dir_size") {
                return;
            }
            const childSize = findSize(item);
            // console.log({ childSize, childFile });
            dir._dir_size += childSize;
            dirSize += childSize;
        });
    }
    return dirSize;
}
function findTotals(dir, acc, limit) {
    if (dir.type == "file") {
        return;
    }
    if (dir._dir_size <= limit) {
        acc.push(dir._dir_size);
    }
    Object.keys(dir).forEach((childFile) => {
        const item = dir[childFile];
        if (childFile == "_dir_size") {
            return;
        }
        findTotals(item, acc, limit);
    });
}
(function solveDay01() {
    return __awaiter(this, void 0, void 0, function* () {
        const array = yield (0, file_1.readFileAsArray)("./2022/data/07/input.txt");
        const root = { _dir_size: 0 };
        const result = array.reduce(({ fs, currentPath, currentNode }, currentCommand) => {
            const isChDirCommand = currentCommand.match(/\$ cd (.*)/);
            const isListCommand = currentCommand.match(/\$ ls/);
            const isDir = currentCommand.match(/dir (.)*/);
            const isFile = currentCommand.match(/(\d+) (.*)/);
            if (isChDirCommand) {
                const dirCommandArg = isChDirCommand[1];
                switch (dirCommandArg) {
                    case "/":
                        currentPath = ["/"];
                        currentNode = fs;
                        break;
                    case "..":
                        currentPath.pop();
                        currentNode = currentPath.reduce((currentNode, dir, index) => {
                            if (index == 0) {
                                return fs;
                            }
                            return currentNode[dir];
                        }, fs);
                        break;
                    default:
                        currentPath.push(dirCommandArg);
                        if (!currentNode[dirCommandArg]) {
                            currentNode[dirCommandArg] = { _dir_size: 0 };
                        }
                        currentNode = currentNode[dirCommandArg];
                        break;
                }
                console.log(currentPath.join());
            }
            if (isListCommand) {
            }
            if (isDir) {
                if (!currentNode[isDir[1]]) {
                    currentNode[isDir[1]] = { _dir_size: 0 };
                }
            }
            if (isFile) {
                currentNode[isFile[2]] = { size: parseInt(isFile[1]), type: "file" };
            }
            return {
                fs,
                currentPath,
                currentNode,
            };
        }, {
            fs: root,
            currentPath: [],
        });
        findSize(result.fs);
        const totalsNeeded = [];
        findTotals(result.fs, totalsNeeded, 100000);
        const totalDirSize = result.fs._dir_size;
        const requiredFreeSpace = 30000000;
        const availableFreeSpace = 70000000 - totalDirSize;
        const additionalFreeSpaceRequired = requiredFreeSpace - availableFreeSpace;
        console.log({ additionalFreeSpaceRequired });
        const deletableNodes = [];
        findTotals(result.fs, deletableNodes, requiredFreeSpace);
        // 13_136_144
        // console.log(totalsNeeded.reduce((sum, item) => sum + item, 0));
        console.log(Math.max(...deletableNodes));
        console.log(Math.max(...deletableNodes.filter(s => s < additionalFreeSpaceRequired)));
        console.log(Math.min(...deletableNodes.filter(s => s >= additionalFreeSpaceRequired)));
    });
})();
//# sourceMappingURL=day07.js.map