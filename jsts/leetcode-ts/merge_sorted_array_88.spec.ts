import { describe, expect, it } from "bun:test";

/**
 * https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
 */

/**
 * SOLVED
 * https://leetcode.com/problems/merge-sorted-array/submissions/989563657/
 */

/**
 * Perform a binary-search leftwise on an array of numbers.
 * @param nums
 * @param sliceLength
 * @param n
 */
function binarySearch(nums: number[], sliceLength: number, n: number): number {
  // function binarySearch(nums, sliceLength, n) {
  let low = 0;
  let high = sliceLength - 1;
  while (low <= high) {
    let mid = Math.floor((low + high) / 2);
    if (nums[mid] > n) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }
  return low;
}

/**
 Do not return anything, modify nums1 in-place instead.
 */
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  if (n === 0) {
    return;
  }

  nums2.forEach((n2, i) => {
    const index = binarySearch(nums1, m + i, n2);
    nums1.pop();
    nums1.splice(index, 0, n2);
  });

  return;
}

describe("merge", () => {
  it("works for empty 2nd array", () => {
    const nums1 = [1];
    const nums2: Array<number> = [];
    merge(nums1, nums1.length, nums2, nums2.length);
    expect(nums1).toEqual([1]);
  });

  it("works", () => {
    const nums1 = [1, 2, 3, 0, 0, 0];
    const nums2 = [2, 5, 6];
    merge(nums1, 3, nums2, 3);
    expect(nums1).toEqual([1, 2, 2, 3, 5, 6]);
  });
});

describe("binarySearch", () => {
  it("works", () => {
    const nums = [1, 2, 3, 4, 5, 6, 0, 0, 0];
    const n = 2;
    expect(binarySearch(nums, 6, n)).toEqual(2);
  });
});
