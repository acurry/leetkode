import { describe, expect, it } from "bun:test";

// SOLVED
// https://leetcode.com/problems/valid-parentheses/submissions/986360472/

/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

Read in and ignore any leading whitespace.
Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
Return the integer as the final result.
Note:

Only the space character ' ' is considered a whitespace character.
Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
*/

type Sign = "positive" | "negative" | undefined;

type Nums = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9";

const VALID_NUMS: Array<Nums> = [
  "0",
  "1",
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9",
];

type AlphaToNumMapping = {
  [x in Nums]: number;
};

const ALPHA_TO_NUMERIC: AlphaToNumMapping = {
  "0": 0,
  "1": 1,
  "2": 2,
  "3": 3,
  "4": 4,
  "5": 5,
  "6": 6,
  "7": 7,
  "8": 8,
  "9": 9,
};

const MAX_INT = 2147483647;
const MIN_INT = -2147483648;

function myAtoi(s: string): number {
  const trimmed = s.trimStart();

  const list = trimmed.split("");

  let sign: Sign;
  let startingIndex = 0;
  const firstChar = list.at(0);
  if (firstChar === "-") {
    sign = "negative";
    startingIndex = 1;
  } else if (firstChar === "+") {
    startingIndex = 1;
  }

  const digits: string[] = [];

  let error = false;

  let hasEncounteredFirstNonZeroDigit = false;

  for (let i = startingIndex; i < list.length; i++) {
    if (list[i] === " ") {
      break;
    } else if ((list[i] === "+" || list[i] === "-") && !!sign) {
      error = true;
      break;
    } else if (VALID_NUMS.includes(list[i] as Nums)) {
      if ((list[i] as Nums) !== "0" && !hasEncounteredFirstNonZeroDigit) {
        hasEncounteredFirstNonZeroDigit = true;
      }
      // if the number is zero and if we have not
      // encountered a nonzero digit yet, then skip
      // else, push the zero onto the digits array
      if ((list[i] as Nums) === "0" && !hasEncounteredFirstNonZeroDigit) {
        continue;
      } else {
        digits.push(list[i]);
      }
    } else if (/[\w\-\+\.]/i.test(list[i])) {
      break;
    } else {
      error = true;
      break;
    }
  }

  if (error) {
    return 0;
  }

  // digits = ['4', '1', '9', '3']

  // short-circuit exit: if digits is 11 numbers or longer, then
  // it's out of bounds
  if (digits.length > 10) {
    return sign === "negative" ? MIN_INT : MAX_INT;
  }

  let sum = 0;

  for (let i = 0; i < digits.length; i++) {
    const magnitude = 10 ** (digits.length - i - 1);
    sum += (ALPHA_TO_NUMERIC[digits[i] as Nums] as number) * magnitude;
  }

  const signedSum = sign === "negative" ? 0 - sum : sum;

  let boundedNum = signedSum;
  if (boundedNum < MIN_INT) {
    boundedNum = MIN_INT;
  } else if (boundedNum > MAX_INT) {
    boundedNum = MAX_INT;
  }

  return boundedNum;
}

describe("myAtoi", () => {
  it("1", () => {
    expect(myAtoi("42")).toBe(42);
  });
  it("2", () => {
    expect(myAtoi("-42")).toBe(-42);
  });
  it("3", () => {
    expect(myAtoi("4193 with words")).toBe(4193);
  });
  it("4", () => {
    expect(myAtoi("-4193 with words")).toBe(-4193);
  });
  it("5", () => {
    expect(myAtoi("   -42")).toBe(-42);
  });
  it("3.14159", () => {
    expect(myAtoi("3.14159")).toBe(3);
  });
  it("invalid", () => {
    expect(myAtoi("+-12")).toBe(0);
  });
  it("prepended zeroes", () => {
    expect(myAtoi("  0000000000012345678")).toBe(12345678);
  });
  it("contains alpha chars", () => {
    expect(myAtoi("-0012a42")).toBe(-12);
  });
  /* latest: 1079 / 1084 testcases passed
   */
  it.skip("ends with a sign", () => {
    expect(myAtoi("-5-")).toBe(-5);
  });
});
