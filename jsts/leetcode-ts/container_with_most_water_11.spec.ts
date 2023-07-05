import { describe, expect, it } from "bun:test";

/*
https://leetcode.com/problems/container-with-most-water/

SOLVED: https://leetcode.com/problems/container-with-most-water/submissions/987222304/
*/

function maxArea(height: number[]): number {
  let max = 0;
  let i = 0;
  let j = height.length - 1;

  while (j > i) {
    const a = height[i];
    const b = height[j];

    const area = Math.min(a, b) * (j - i);

    if (area > max) max = area;

    if (b > a) i++;
    else j--;
  }

  return max;
}

describe("maxArea", () => {
  it("1", () => {
    expect(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7])).toBe(49);
  });
  it("2", () => {
    expect(maxArea([1, 1])).toBe(1);
  });
});
