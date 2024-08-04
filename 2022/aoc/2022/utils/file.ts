import { promises as fs } from "fs";
export async function readFileAsArray(file: string) {
    const contents = await fs.readFile(file);
  return contents.toString().replace(/\r\n/g, "\n").split("\n");

}