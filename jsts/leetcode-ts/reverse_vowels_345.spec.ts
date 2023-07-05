import { describe, expect, it } from "bun:test";

/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/

const VOWELS = ["a", "A", "e", "E", "i", "I", "o", "O", "u", "U"];

function reverseVowels(s: string): string {
  const stack: string[] = [];

  const switched = s.split("");

  switched.forEach((c, index) => {
    if (VOWELS.includes(c)) {
      stack.push(c);
      switched[index] = "REPLACE_ME";
    }
  });

  switched.forEach((c, index) => {
    if (c == "REPLACE_ME") {
      switched[index] = stack.pop() as string;
    }
  });

  return switched.join("");
}

describe("reverseVowels", () => {
  it('reverses "hello" to "holle"', () => {
    expect(reverseVowels("hello")).toBe("holle");
  });

  it('reverses "leetcode" to "leotcede"', () => {
    expect(reverseVowels("leetcode")).toBe("leotcede");
  });
});
