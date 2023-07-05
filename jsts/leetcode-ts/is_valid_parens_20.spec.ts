import { describe, expect, it } from "bun:test";

// SOLVED
// https://leetcode.com/problems/valid-parentheses/submissions/986360472/

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

class Stack<T> {
  private storage: T[] = [];

  constructor(private capacity: number = Infinity) {}

  push(item: T): void {
    if (this.size() === this.capacity) {
      throw Error("Stack has reached max capacity, you cannot add more items");
    }
    this.storage.push(item);
  }

  pop(): T | undefined {
    return this.storage.pop();
  }

  peek(): T | undefined {
    return this.storage[this.size() - 1];
  }

  size(): number {
    return this.storage.length;
  }
}

const OPENERS = ["(", "{", "["];

function isValid(s: string): boolean {
  const stack = new Stack();

  const list = s.split("");

  if (list.length % 2 !== 0) {
    return false;
  }

  let valid = true;

  for (let i = 0; i < list.length; i++) {
    const c = list[i];
    if (OPENERS.includes(c)) {
      stack.push(c);
    } else {
      const opener_to_expect = stack.peek();
      if (c === ")") {
        if (opener_to_expect === "(") {
          stack.pop();
        } else {
          valid = false;
          break;
        }
      } else if (c === "}") {
        if (opener_to_expect === "{") {
          stack.pop();
        } else {
          valid = false;
          break;
        }
      } else if (c === "]") {
        if (opener_to_expect === "[") {
          stack.pop();
        } else {
          valid = false;
          break;
        }
      }
    }
  }

  return valid && stack.size() == 0;
}

describe("isValid", () => {
  it("works for test case 1: '()' = true", () => {
    expect(isValid("()")).toBeTrue();
  });

  it("works for test case 2: '()[]{}' = true", () => {
    expect(isValid("()[]{}")).toBeTrue();
  });

  it("works for test case 3: '(]' = false", () => {
    expect(isValid("(]")).toBeFalse();
  });

  it("test case 4: '[' = false", () => {
    expect(isValid("[")).toBeFalse();
  });

  it("test case 5: '((' = false", () => {
    expect(isValid("((")).toBeFalse();
  });
});
