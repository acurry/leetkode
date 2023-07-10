/**
 * https://leetcode.com/problems/path-sum/?envType=study-plan-v2&envId=top-interview-150
 */

/**
 * Latest progress:
 * Wrong Answer

97 / 117 testcases passed
Input
root =
[1,2,null,3,null,4,null,5]
targetSum =
15
Use Testcase
Output
false
Expected
true
 */

import { describe, expect, it } from "bun:test";

class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
  if (root == null) {
    return false;
  } else if (root.left !== null && root.right === null) {
    return targetSum === root.val + root.left.val;
  } else if (root.left === null && root.right !== null) {
    return targetSum === root.val + root.right.val;
  }

  const left_hasPathSum = recur(root.left, targetSum - root.val);
  const right_hasPathSum = recur(root.right, targetSum - root.val);

  return left_hasPathSum || right_hasPathSum;
}

function recur(root: TreeNode | null, targetSum: number): boolean {
  if (root == null) {
    return targetSum === 0;
  }

  const left = recur(root.left, targetSum - root.val);
  const right = recur(root.right, targetSum - root.val);

  return left || right;
}

const FAILING_TEST_CASE_1 = new TreeNode(1, new TreeNode(2, null, null), null);

// [5,4,8,11,null,13,4,7,2,null,null,null,1]
const FAILING_TEST_CASE_2 = new TreeNode(
  5,
  new TreeNode(4, new TreeNode(11, new TreeNode(7), new TreeNode(2))),
  new TreeNode(8, new TreeNode(13), new TreeNode(4, null, new TreeNode(1)))
);

// [1,2,null,3,null,4,null,5]
/*
            1
          /
        2
      /
    3     4
*/
const FAILING_TEST_CASE_3 = new TreeNode(
  1,
  new TreeNode(2, new TreeNode(3)),
  null
);

describe("hasPathSum", () => {
  it("works for []", () => {
    expect(hasPathSum(null, 0)).toBeFalse();
  });
  it("works for [1,2], 0", () => {
    expect(hasPathSum(FAILING_TEST_CASE_1, 0)).toBeFalse();
  });
  it("works for [1,2], 1", () => {
    expect(hasPathSum(FAILING_TEST_CASE_1, 1)).toBeFalse();
  });
  it("works for [5,4,8,11,null,13,4,7,2,null,null,null,1]", () => {
    expect(hasPathSum(FAILING_TEST_CASE_2, 22)).toBeTrue();
  });
});
