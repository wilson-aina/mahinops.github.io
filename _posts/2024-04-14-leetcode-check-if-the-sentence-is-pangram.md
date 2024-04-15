---
title: "Leetcode - 1832. Check if the Sentence Is Pangram"
date: 2024-04-14
categories: [Leetcode]
tags: [Leetcode, Problem Solving, Golang, DSA, Data Structure, Programming, Algorithm, Hash Table, String]
---

# [Leetcode - 1832. Check if the Sentence Is Pangram](https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/){:target="_blank"}
[![Hits](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-check-if-the-sentence-is-pangram.svg)](https://hits.sh/mokhlesurr031.github.io/posts/leetcode-check-if-the-sentence-is-pangram/)

- A pangram is a sentence where every letter of the English alphabet appears at least once.

  Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.



Solution
```go
func checkIfPangram(sentence string) bool {
    charMap := make(map[rune]bool)

	for _, char := range strings.ToLower(sentence){
		if unicode.IsLetter(char){
			charMap[char]=true
		}
	}
	return len(charMap)==26
}

```
